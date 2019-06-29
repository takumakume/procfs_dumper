package main

import (
	"log"
	"os"

	"github.com/takumakume/procfs_dumper"
	"github.com/urfave/cli"
)

var version = ""

type params struct {
	pid        int
	allPids    bool
	procfsPath string
}

func main() {
	opts := params{}

	app := cli.NewApp()
	app.Name = "procfs_dump"
	app.Usage = "Dump a procfs"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "pid, p",
			Usage:       "Process ID",
			Destination: &opts.pid,
		},
		cli.BoolFlag{
			Name:        "all-pids, P",
			Usage:       "All Process ID",
			Destination: &opts.allPids,
		},
		cli.StringFlag{
			Name:        "procfs-path",
			Usage:       "Specify path of procfs",
			Destination: &opts.procfsPath,
		},
	}

	config := procfsdumper.NewConfig()
	if opts.procfsPath != "" {
		config.Path = opts.procfsPath
	}

	pd, err := procfsdumper.NewProcFSDumper(config)
	if err != nil {
		return
	}

	pd.Out = os.Stdout
	pd.Err = os.Stderr

	app.Action = func(ctx *cli.Context) error {
		if opts.pid != 0 {
			pd.ProcessByPid(opts.pid)
		} else if opts.allPids {
			pd.AllProcesses()
		} else {
			cli.ShowCommandHelp(ctx, "")
			return cli.NewExitError("", 1)
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
