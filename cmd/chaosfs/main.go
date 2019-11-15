package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/ethercflow/hookfs/hookfs"
	"github.com/golang/glog"

	"github.com/pingcap/chaos-operator/pkg/chaosfs"
	"github.com/pingcap/chaos-operator/pkg/pidfile"
	"github.com/pingcap/chaos-operator/pkg/signals"
	"github.com/pingcap/chaos-operator/pkg/version"
)

var (
	addr         = flag.String("addr", ":65534", "The address to bind to")
	pidFile      = flag.String("pidfile", "", "PidFile")
	original     = flag.String("original", "", "ORIGINAL")
	mountpoint   = flag.String("mountpoint", "", "MOUNTPOINT")
	printVersion = flag.Bool("version", false, "print version information and exit")

	pf *pidfile.PIDFile
)

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
}

func main() {
	if *printVersion {
		fmt.Printf("ChaosFS verion: %#v\n", version.Get())
		os.Exit(0)
	}

	if *original == "" || *mountpoint == "" {
		glog.Fatal("invalid original or mountpoint")
	}

	if *pidFile == "" {
		glog.Fatal("invalid pidfile")
	}

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	go func() {
		sig := <-stopCh
		fmt.Printf("\nGot signal [%v] to exit.\n", sig)

		select {
		case <-stopCh:
			fmt.Printf("\nGot signal [%v] again to exit.\n", sig)
		case <-time.After(10 * time.Second):
			fmt.Print("\nWait 10s for closed, force exit\n")
		}

		if pf != nil {
			if err := pf.Remove(); err != nil {
				glog.Error(err)
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := exec.CommandContext(ctx, "fusermount", "-u", *mountpoint).Run()
		if err != nil {
			if err1 := exec.CommandContext(ctx, "umount", "-l", *mountpoint).Run(); err1 != nil {
				glog.Errorln("fusermount failed: ", err, " even umount failed too: ", err1)
			}
			glog.Error(err)
		}
		os.Exit(0)
	}()

	fs, err := hookfs.NewHookFs(*original, *mountpoint, &chaosfs.InjuredHook{Addr: *addr})
	if err != nil {
		glog.Error(err)
	}

	pf, err := pidfile.New(*pidFile)
	if err != nil {
		glog.Fatalln("Create pidfile failed: ", err)
	}
	defer func() {
		if err := pf.Remove(); err != nil {
			glog.Error(err)
		}
	}()

	if err = fs.Serve(); err != nil {
		glog.Error(err)
	}
}
