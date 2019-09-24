package api_server

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cwen0/chaos-operator/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/juju/errors"
	"github.com/ngaut/log"
)

type mysqlClient struct {
	db *sqlx.DB
}

func NewMysqlClient(dataSource string) (*mysqlClient, error) {
	log.Infof("connecting to %s", dataSource)
	db, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		return nil, errors.Trace(err)
	}
	log.Info("database connected")

	return &mysqlClient{
		db,
	}, nil
}

func (m *mysqlClient) createJob(job *Job) error {
	t := time.Now().Format(util.TimeFormat)
	job.Ctime = t

	tx, err := m.db.Beginx()
	if err != nil {
		log.Error(err)
		return errors.Trace(err)
	}

	resource, err := json.Marshal(job.Resource)
	if err != nil {
		log.Error(err)
		return errors.Trace(err)
	}
	result, err := tx.NamedExec(jobInsert, map[string]interface{}{
		"event_type":  job.EventType,
		"resource":    string(resource),
		"create_time": job.Ctime,
	})
	if err != nil {
		tx.Rollback()
		log.Error(err)
		return errors.Trace(err)
	}

	job.ID, err = result.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Error(err)
		return errors.Trace(err)
	}

	_, err = tx.NamedExec(jobPodInsert, job.getPodRelation())
	if err != nil {
		tx.Rollback()
		log.Error(err)
		return errors.Trace(err)
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		return errors.Trace(err)
	}

	return nil
}

func (m *mysqlClient) getJobs(fs *Filters) ([]*Job, error) {
	filtersSQL, err := GenSQL(fs)
	if err != nil {
		return nil, errors.Trace(err)
	}

	rows, err := m.db.Queryx(fmt.Sprintf(jobSelect, " WHERE "+filtersSQL))
	if err != nil {
		log.Error(err)
		return nil, errors.Trace(err)
	}

	var jobs []*Job
	for rows.Next() {
		job := new(JobPodJoinSelect)
		rows.StructScan(&job)

		resource, ok := job.Resource.([]byte)
		if !ok {
			return nil, errors.New("resource is not []byte")
		}
		json.Unmarshal(resource, &job.Resource)

		job.Pods = strings.Split(job.PodsStr, ",")
		jobs = append(jobs, &job.Job)
	}

	return jobs, nil
}

func (job *Job) getPodRelation() []JobPodRelation {
	var list []JobPodRelation

	for _, pod := range job.Pods {
		list = append(list, JobPodRelation{
			JobID: job.ID,
			Pod:   pod,
		})
	}

	return list
}

const jobInsert = `
	INSERT INTO job (
		event_type,
		resource,
		create_time
	) VALUES (
		:event_type,
		:resource,
		:create_time
	)
`

const jobPodInsert = `
	INSERT INTO job_pod (
		job_id,
		pod
	) VALUES (
		:job_id,
		:pod
	)
`

const jobSelect = `
  SELECT id,event_type,resource,create_time,GROUP_CONCAT(pod separator ',') AS pods_str FROM job JOIN job_pod ON id=job_id %s GROUP BY id
`
