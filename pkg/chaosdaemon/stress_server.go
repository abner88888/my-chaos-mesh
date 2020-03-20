// Copyright 2020 PingCAP, Inc.
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

package chaosdaemon

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/containerd/cgroups"
	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/pingcap/chaos-mesh/pkg/chaosdaemon/pb"
)

var (
	stressorLocker = new(sync.Mutex)
	podStressors   = make(map[string]*exec.Cmd)

	// Possible cgroup subsystems
	cgroup_subsys = []string{"cpu", "memory", "systemd", "net_cls",
		"net_prio", "freezer", "blkio", "perf_event", "devices",
		"cpuset", "cpuacct", "pids", "hugetlb"}
)

func (s *daemonServer) ExecPodStressors(ctx context.Context,
	req *pb.StressRequest) (*empty.Empty, error) {
	log.Info("executing stressors", "request", req)
	pid, err := s.crClient.GetPidFromContainerID(ctx, req.Target)
	if err != nil {
		return nil, err
	}
	path := cgroups.PidPath(int(pid))
	cgroup, err := findValidCgroup(path, req.Target)
	if err != nil {
		return nil, err
	}
	dir, _ := filepath.Split(cgroup)
	control, err := cgroups.Load(cgroups.V1, cgroups.StaticPath(dir))
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("stress-ng", strings.Fields(req.Stressors)...)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	if err = control.Add(cgroups.Process{Pid: cmd.Process.Pid}); err != nil {
		if err := cmd.Process.Kill(); err != nil {
			return nil, err
		}
	}
	stressorLocker.Lock()
	defer stressorLocker.Unlock()
	podStressors[req.Target] = cmd
	go func() {
		if err, ok := cmd.Wait().(*exec.ExitError); ok {
			status := err.Sys().(syscall.WaitStatus)
			if status.Signaled() && status.Signal() == syscall.SIGKILL {
				log.Info("stressors cancelled", "request", req)
			} else {
				log.Error(err, "stressors exited accidentally", "request", req)
			}
		}
		stressorLocker.Lock()
		defer stressorLocker.Unlock()
		delete(podStressors, req.Target)
	}()

	return &empty.Empty{}, nil
}

func (s *daemonServer) CancelPodStressors(ctx context.Context,
	req *pb.StressRequest) (*empty.Empty, error) {
	log.Info("canceling stressors", "request", req)
	if cmd, ok := podStressors[req.Target]; ok {
		if err := cmd.Process.Kill(); err != nil {
			log.Error(err, "fail to exit stressors", "pid", cmd.Process.Pid)
			return nil, err
		}
	}
	return &empty.Empty{}, nil
}

func findValidCgroup(path cgroups.Path, target string) (string, error) {
	for _, subsys := range cgroup_subsys {
		if p, _ := path(cgroups.Name(subsys));
			strings.Contains(p, target) {
			return p, nil
		}
	}
	return "", fmt.Errorf("never found cgroup for %s", target)
}
