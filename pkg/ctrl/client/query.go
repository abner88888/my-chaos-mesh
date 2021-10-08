// Copyright 2021 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hasura/go-graphql-client"
	"github.com/iancoleman/strcase"
)

const (
	SeperatorSegment  = "/"
	SeperatorArgument = ":"
	SeperatorTag      = "@"
	SeperatorLeaf     = ","
	TagNameAll        = "all"
)

type Query struct {
	Name      string
	Argument  string
	ArgValue  interface{} // inner immutable
	Type      *Type
	Tag       *Tag          // inner immutable
	Decorator TypeDecorator // inner immutable
	Fields    map[string]*Query
}

type Segment struct {
	Name     string
	Argument *string
	Tag      *Tag
}

type Tag struct {
	Name     string
	Argument *string
}

type (
	EnumValueNotFound struct {
		Target   string
		Variants []string
	}

	FieldNotFound struct {
		Target string
		Type   string
		Fields []string
	}

	ScalarValueParseFail struct {
		Value    string
		Argument string
	}

	RequireArgument struct {
		Leaf     string
		Argument string
	}
)

type Component string

type Variables struct {
	inner []*Variable
}

type Variable struct {
	name  string
	Value interface{}
}

func NewVariables() *Variables {
	return &Variables{
		inner: make([]*Variable, 0),
	}
}

func (v *Variable) Name(index int) string {
	return fmt.Sprintf("%s%d", v.name, index)
}

func (v *Variables) Append(name string, val interface{}) string {
	variable := &Variable{
		name:  name,
		Value: val,
	}

	v.inner = append(v.inner, variable)
	return variable.Name(len(v.inner) - 1)
}

func (v *Variables) GenMap() map[string]interface{} {
	variableMap := make(map[string]interface{})
	for i, val := range v.inner {
		variableMap[val.Name(i)] = val.Value
	}
	return variableMap
}

var (
	enumMap = map[string]reflect.Type{
		"Component": reflect.TypeOf(Component("")),
	}
)

func (e *EnumValueNotFound) Error() string {
	return fmt.Sprintf("enum value `%s` not found", e.Target)
}

func (e *FieldNotFound) Error() string {
	return fmt.Sprintf("field `%s` not found in type: %s; fields: %#v", e.Target, e.Type, e.Fields)
}

func (e ScalarValueParseFail) Error() string {
	return fmt.Sprintf("fail to parse scalar value `%s`", e.Value)
}

func (e RequireArgument) Error() string {
	return fmt.Sprintf("field `%s` require argument", e.Leaf)
}

func NewQuery(name string, typ *Type, decorator TypeDecorator) *Query {
	return &Query{
		Name:      name,
		Type:      typ,
		Decorator: decorator,
		Fields:    make(map[string]*Query),
	}
}

func (q *Query) SetArgument(argument, value string, typ *Type) (err error) {
	switch typ.Kind {
	case EnumKind:
		err = q.setEnumValue(value, typ)
	case ScalarKind:
		err = q.setScalarValue(argument, value, typ)
	default:
		err = fmt.Errorf("type %#v is not supported as arguments", typ)
	}

	if err != nil {
		return
	}

	q.Argument = argument
	return nil
}

func (q *Query) setEnumValue(value string, typ *Type) error {
	reflectType, ok := enumMap[string(typ.Name)]
	if !ok {
		return fmt.Errorf("unsupported enum type: %s", typ.Name)
	}

	variant, err := typ.mustGetEnum(value)
	if err != nil {
		return err
	}

	enumValue := reflect.New(reflectType)
	enumValue.Elem().SetString(string(variant.Name))
	q.ArgValue = enumValue.Elem().Interface()
	return nil
}

func (q *Query) setScalarValue(name, value string, typ *Type) error {
	parseErr := &ScalarValueParseFail{
		Argument: name,
		Value:    value,
	}

	switch typ.Name {
	case ScalarString:
		q.ArgValue = graphql.String(value)
	case ScalarBoolean:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return parseErr
		}
		q.ArgValue = graphql.Boolean(v)
	case ScalarInt:
		v, err := strconv.Atoi(value)
		if err != nil {
			return parseErr
		}
		q.ArgValue = graphql.Int(v)
	case ScalarFloat:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return parseErr
		}
		q.ArgValue = graphql.Float(v)
	default:
		return fmt.Errorf("unsupported argument type: %#v", typ)
	}
	return nil
}

func (q Query) Clone() *Query {
	newFields := make(map[string]*Query)
	for name, field := range q.Fields {
		newFields[name] = field.Clone()
	}
	q.Fields = newFields
	return &q
}

func (q *Query) Merge(other *Query) (*Query, error) {
	if q == nil {
		return other.Merge(q)
	}

	if other == nil {
		return q.Clone(), nil
	}

	if q.Name != other.Name || q.Argument != other.Argument || !reflect.DeepEqual(q.ArgValue, other.ArgValue) {
		return nil, fmt.Errorf("query %#v cannot merge %#v", q, other)
	}

	newQuery := q.Clone()
	for name, field := range other.Fields {
		newField, err := newQuery.Fields[name].Merge(field)
		if err != nil {
			return nil, err
		}
		newQuery.Fields[name] = newField
	}
	return newQuery, nil
}

func (q *Query) Tail() *Query {
	if q == nil {
		return q
	}

	for _, field := range q.Fields {
		return field.Tail()
	}

	return q.Clone()
}

func (t *Tag) String() string {
	tag := t.Name
	if t.Argument != nil {
		tag = fmt.Sprintf("%s%s%s", tag, SeperatorArgument, *t.Argument)
	}
	return tag
}

func (s *Segment) String() string {
	segment := s.Name
	if s.Argument != nil {
		segment = fmt.Sprintf("%s%s%s", segment, SeperatorArgument, *s.Argument)
	}
	if s.Tag != nil {
		segment = fmt.Sprintf("%s%s%s", segment, SeperatorTag, s.Tag.String)
	}
	return segment
}

func (q *Query) String() string {
	segment := q.Name
	if q.ArgValue != nil {
		switch v := q.ArgValue.(type) {
		case graphql.String:
			segment = strings.Join([]string{segment, string(v)}, SeperatorArgument)
		case graphql.Boolean:
			segment = strings.Join([]string{segment, strconv.FormatBool(bool(v))}, SeperatorArgument)
		case graphql.Int:
			segment = strings.Join([]string{segment, fmt.Sprintf("%d", v)}, SeperatorArgument)
		case graphql.Float:
			segment = strings.Join([]string{segment, fmt.Sprintf("%f", v)}, SeperatorArgument)
		default:
			value := reflect.ValueOf(q.ArgValue)
			if value.Type().Kind() == reflect.String {
				arg := value.Convert(reflect.TypeOf("")).Interface().(string)
				segment = strings.Join([]string{segment, strings.ToLower(arg)}, SeperatorArgument)
			}
		}
	}

	if q.Tag != nil {
		segment = strings.Join([]string{segment, q.Tag.String()}, SeperatorTag)
	}

	if len(q.Fields) == 0 {
		return segment
	}

	fields := make([]string, 0, len(q.Fields))
	for _, f := range q.Fields {
		fields = append(fields, f.String())
	}

	fieldStr := strings.Join(fields, SeperatorLeaf)
	return strings.Join([]string{segment, fieldStr}, SeperatorSegment)
}

func parseSegment(segmentStr string) (*Segment, error) {
	segment := new(Segment)
	rawSegment := strings.Split(strings.Trim(segmentStr, SeperatorTag), SeperatorTag)
	if len(rawSegment) != 2 && len(rawSegment) != 1 {
		return nil, fmt.Errorf("invalid segment: %s", segmentStr)
	}

	rawName := strings.Split(strings.Trim(rawSegment[0], SeperatorArgument), SeperatorArgument)
	if len(rawName) != 2 && len(rawName) != 1 {
		return nil, fmt.Errorf("invalid segment: %s", segmentStr)
	}

	segment.Name = rawName[0]
	if len(rawName) == 2 {
		segment.Argument = &rawName[1]
	}

	if len(rawSegment) == 2 {
		rawTag := strings.Split(strings.Trim(rawSegment[1], SeperatorArgument), SeperatorArgument)
		if len(rawTag) != 2 && len(rawTag) != 1 {
			return nil, fmt.Errorf("invalid segment: %s", segmentStr)
		}

		segment.Tag = &Tag{
			Name: rawTag[0],
		}

		if len(rawTag) == 2 {
			segment.Tag.Argument = &rawTag[1]
		}
	}

	return segment, nil
}

func (s *Schema) ParseQuery(rawQuery []string, super *Type, partial bool) (*Query, error) {
	var query []*Segment
	for _, s := range rawQuery {
		segment, err := parseSegment(s)
		if err != nil {
			return nil, err
		}
		query = append(query, segment)
	}

	queries, err := s.parseQuery(query, super, partial)
	if err != nil {
		return nil, err
	}

	if len(queries) == 0 {
		return nil, errors.New("empty query")
	}

	return queries[0], nil
}

func (s *Schema) parseQuery(query []*Segment, super *Type, partial bool) ([]*Query, error) {
	if len(query) == 0 {
		return nil, nil
	}

	if len(query) == 1 {
		fields := strings.Split(query[0].Name, SeperatorLeaf)
		queries := make([]*Query, 0, len(fields))
		for _, f := range fields {
			field, err := super.mustGetField(f)
			if err != nil {
				return nil, err
			}

			if len(field.Args) != 0 && !partial {
				// TODO: support default args
				return nil, &RequireArgument{
					Leaf:     f,
					Argument: string(field.Args[0].Name),
				}
			}

			typ, err := s.resolve(&field.Type)
			if err != nil {
				return nil, err
			}

			if typ.Kind != ScalarKind && !partial {
				// TODO: support object kind
				return nil, fmt.Errorf("type %s is not a scalar kind", typ.Name)
			}

			decorator, err := s.typeDecorator(&field.Type)
			if err != nil {
				return nil, err
			}

			queries = append(queries, NewQuery(f, typ, decorator))
		}
		return queries, nil
	}

	segment := query[0]
	subQuery := query[1:]

	field, err := super.mustGetField(segment.Name)
	if err != nil {
		return nil, err
	}

	typ, err := s.resolve(&field.Type)
	if err != nil {
		return nil, err
	}

	decorator, err := s.typeDecorator(&field.Type)
	if err != nil {
		return nil, err
	}

	newQuery := NewQuery(segment.Name, typ, decorator)
	newQuery.Tag = segment.Tag

	if len(field.Args) != 0 {
		if segment.Argument != nil {
			argument := field.Args[0]
			argType, err := s.resolve(&argument.Type)
			if err != nil {
				return nil, err
			}
			err = newQuery.SetArgument(string(argument.Name), *segment.Argument, argType)
			if err != nil {
				return nil, err
			}
		} else if segment.Tag != nil && segment.Tag.Name == TagNameAll {
			// do nothing
		} else if field.Args[0].DefaultValue == nil && !partial {
			// argument is nil
			return nil, fmt.Errorf("query is imcomplete: path(%s) needs argument(%s)", field.Name, field.Args[0].Name)
		}
	}

	fields, err := s.parseQuery(subQuery, typ, partial)
	if err != nil {
		return nil, err
	}
	for _, field := range fields {
		newQuery.Fields[field.Name] = field
	}

	return []*Query{newQuery}, nil
}

func (s *Schema) Reflect(query *Query, variables *Variables) (typ reflect.Type, err error) {
	nilType := reflect.TypeOf(nil)

	switch query.Type.Kind {
	case ScalarKind:
		typ, err = ScalarType(query.Type.Name).reflect()
		if err != nil {
			return nilType, err
		}
	case EnumKind:
		typ, err = ScalarType(ScalarString).reflect()
		if err != nil {
			return nilType, err
		}
	case ObjectKind:
		fields := make([]reflect.StructField, 0)
		for _, f := range query.Fields {
			fieldType, err := s.Reflect(f, variables)
			if err != nil {
				return nilType, err
			}

			field := reflect.StructField{
				Name: strcase.ToCamel(f.Name),
				Type: fieldType,
			}

			if f.Argument != "" && f.ArgValue != nil {
				variableName := variables.Append(f.Argument, f.ArgValue)
				field.Tag = reflect.StructTag(fmt.Sprintf(`graphql:"%s(%s: $%s)"`, f.Name, f.Argument, variableName))
			}

			fields = append(fields, field)
		}
		typ = reflect.PtrTo(reflect.StructOf(fields))
	default:
		return nilType, fmt.Errorf("unsupported type kind: %s", query.Type.Kind)
	}

	if query.Decorator != nil {
		typ, err = query.Decorator(typ)
	}

	return typ, err
}

func StandardizeQuery(query string) []string {
	return strings.Split(strings.Trim(query, SeperatorSegment), SeperatorSegment)
}
