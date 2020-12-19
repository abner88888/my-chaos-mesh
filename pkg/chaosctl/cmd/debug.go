// Copyright 2020 Chaos Mesh Authors.
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

package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"

	cm "github.com/chaos-mesh/chaos-mesh/pkg/chaosctl/common"
	"github.com/chaos-mesh/chaos-mesh/pkg/chaosctl/debug/iochaos"
	"github.com/chaos-mesh/chaos-mesh/pkg/chaosctl/debug/networkchaos"
	"github.com/chaos-mesh/chaos-mesh/pkg/chaosctl/debug/stresschaos"
)

type debugOptions struct {
	namespace string
}

type chaosT int

const (
	networkChaos = "networkchaos"
	stressChaos  = "stresschaos"
	ioChaos      = "iochaos"
)

func init() {
	o := &debugOptions{}

	c, err := cm.InitClientSet()
	if err != nil {
		log.Fatal(err)
	}

	debugCmd := &cobra.Command{
		Use:   `debug (CHAOSTYPE) [-c CHAOSNAME] [-n NAMESPACE]`,
		Short: `Print the debug information for certain chaos`,
		Long: `Print the debug information for certain chaos.
Currently support networkchaos, stresschaos and iochaos.

Examples:
  # Return debug information from all networkchaos in default namespace
  chaosctl debug networkchaos

  # Return debug information from certain networkchaos
  chaosctl debug networkchaos CHAOSNAME -n NAMESPACE`,
		ValidArgsFunction: noCompletions,
	}

	// Need to separately support chaos-level completion, so split each chaos apart
	networkCmd := &cobra.Command{
		Use:   `networkchaos (CHAOSNAME) [-n NAMESPACE]`,
		Short: `Print the debug information for certain network chaos`,
		Long:  `Print the debug information for certain network chaos`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.Run("networkchaos", args, c); err != nil {
				log.Fatal(err)
			}
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			return listChaos("networkchaos", o.namespace, toComplete, c.CtrlCli)
		},
	}

	stressCmd := &cobra.Command{
		Use:   `stresschaos (CHAOSNAME) [-n NAMESPACE]`,
		Short: `Print the debug information for certain stress chaos`,
		Long:  `Print the debug information for certain stress chaos`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.Run("stresschaos", args, c); err != nil {
				log.Fatal(err)
			}
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			return listChaos("stresschaos", o.namespace, toComplete, c.CtrlCli)
		},
	}

	ioCmd := &cobra.Command{
		Use:   `iochaos (CHAOSNAME) [-n NAMESPACE]`,
		Short: `Print the debug information for certain io chaos`,
		Long:  `Print the debug information for certain io chaos`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.Run("iochaos", args, c); err != nil {
				log.Fatal(err)
			}
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			return listChaos("iochaos", o.namespace, toComplete, c.CtrlCli)
		},
	}

	debugCmd.AddCommand(networkCmd)
	debugCmd.AddCommand(stressCmd)
	debugCmd.AddCommand(ioCmd)

	debugCmd.PersistentFlags().StringVarP(&o.namespace, "namespace", "n", "default", "namespace to find chaos")
	err = debugCmd.RegisterFlagCompletionFunc("namespace", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return listNamespace(toComplete, c.KubeCli)
	})
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(debugCmd)
}

// Run debug
func (o *debugOptions) Run(chaosType string, args []string, c *cm.ClientSet) error {
	if len(args) > 1 {
		return fmt.Errorf("Only one chaos could be specified")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chaosName := ""
	if len(args) == 1 {
		chaosName = args[0]
	}

	chaosList, chaosNameList, err := cm.GetChaosList(ctx, chaosType, chaosName, o.namespace, c.CtrlCli)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	var result []cm.ChaosResult

	for i, chaos := range chaosList {
		var chaosResult cm.ChaosResult
		chaosResult.Name = chaosNameList[i]

		var err error
		switch chaosType {
		case networkChaos:
			err = networkchaos.Debug(ctx, chaos, c, &chaosResult)
		case stressChaos:
			err = stresschaos.Debug(ctx, chaos, c, &chaosResult)
		case ioChaos:
			err = iochaos.Debug(ctx, chaos, c, &chaosResult)
		default:
			return fmt.Errorf("chaos type not supported")
		}
		result = append(result, chaosResult)
		if err != nil {
			cm.PrintResult(result)
			return err
		}
	}
	cm.PrintResult(result)
	return nil
}

func listNamespace(toComplete string, c *kubernetes.Clientset) ([]string, cobra.ShellCompDirective) {
	namespaces, err := c.CoreV1().Namespaces().List(v1.ListOptions{})
	if err != nil {
		return nil, cobra.ShellCompDirectiveDefault
	}
	var ret []string
	for _, ns := range namespaces.Items {
		if strings.HasPrefix(ns.Name, toComplete) {
			ret = append(ret, ns.Name)
		}
	}
	return ret, cobra.ShellCompDirectiveNoFileComp
}

func listChaos(chaosType string, namespace string, toComplete string, c client.Client) ([]string, cobra.ShellCompDirective) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, chaosList, err := cm.GetChaosList(ctx, chaosType, "", namespace, c)
	if err != nil {
		return nil, cobra.ShellCompDirectiveDefault
	}
	var ret []string
	for _, chaos := range chaosList {
		if strings.HasPrefix(chaos, toComplete) {
			ret = append(ret, chaos)
		}
	}
	return ret, cobra.ShellCompDirectiveNoFileComp
}
