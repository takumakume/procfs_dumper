package procfsdumper

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/prometheus/procfs"
)

type ProcFSDumper struct {
	fs          procfs.FS
	concurrency int
	Out, Err    io.Writer
}

func NewProcFSDumper(c Config) (*ProcFSDumper, error) {
	p := ProcFSDumper{}
	fs, err := procfs.NewFS(c.Path)
	if err != nil {
		return &p, err
	}
	p.fs = fs
	p.concurrency = c.Concurrency

	return &p, nil
}

func (self *ProcFSDumper) ProcessByPid(pid int) {
	proc, err := self.fs.NewProc(pid)
	if err != nil {
		self.outputError(err)
	}

	process, warnings := self.fetchProcess(proc)
	if err := self.outputJSON(process); err != nil {
		self.outputError(err)
	}

	if len(warnings) == 0 {
		self.outputErrors(warnings)
	}

}

func (self *ProcFSDumper) AllProcesses() {
	procs, err := self.fs.AllProcs()
	if err != nil {
		self.outputError(err)
	}

	processes, warnings := self.fetchProcesses(procs)
	if err := self.outputJSON(processes); err != nil {
		self.outputError(err)
	}

	if len(warnings) == 0 {
		self.outputErrors(warnings)
	}
}

func (self *ProcFSDumper) fetchProcess(proc procfs.Proc) (Process, []error) {
	warnings := []error{}

	p := Process{
		PID: proc.PID,
	}

	var err error
	if p.CmdLine, err = proc.CmdLine(); err != nil {
		warnings = append(warnings, err)
	}
	if p.Comm, err = proc.Comm(); err != nil {
		warnings = append(warnings, err)
	}
	if p.Cwd, err = proc.Cwd(); err != nil {
		warnings = append(warnings, err)
	}
	if p.Executable, err = proc.Executable(); err != nil {
		warnings = append(warnings, err)
	}
	if p.FileDescriptorTargets, err = proc.FileDescriptorTargets(); err != nil {
		warnings = append(warnings, err)
	}
	if p.IO, err = proc.IO(); err != nil {
		warnings = append(warnings, err)
	}
	if p.Limits, err = proc.Limits(); err != nil {
		warnings = append(warnings, err)
	}
	if p.MountStats, err = proc.MountStats(); err != nil {
		warnings = append(warnings, err)
	}
	if p.Namespaces, err = proc.Namespaces(); err != nil {
		warnings = append(warnings, err)
	}
	if p.NetDev, err = proc.NetDev(); err != nil {
		warnings = append(warnings, err)
	}
	if p.RootDir, err = proc.RootDir(); err != nil {
		warnings = append(warnings, err)
	}
	if p.Stat, err = proc.Stat(); err != nil {
		warnings = append(warnings, err)
	}
	if p.Environ, err = proc.Environ(); err != nil {
		warnings = append(warnings, err)
	}

	return p, warnings
}

func (self *ProcFSDumper) fetchProcesses(procs []procfs.Proc) (Processes, []error) {
	wargings := []error{}
	processes := Processes{}
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, self.concurrency)

	for _, proc := range procs {
		wg.Add(1)
		proc := proc
		go func(proc procfs.Proc) {
			ch <- struct{}{}
			defer wg.Done()
			process, errs := self.fetchProcess(proc)
			processes = append(processes, process)
			wargings = append(wargings, errs...)
			<-ch
		}(proc)
	}
	wg.Wait()

	return processes, wargings
}

func (self *ProcFSDumper) outputJSON(o interface{}) error {
	json, err := json.Marshal(o)
	if err != nil {
		return err
	}
	fmt.Fprintln(self.Out, string(json))
	return nil
}

func (self *ProcFSDumper) outputError(e error) {
	fmt.Fprintln(self.Err, e.Error)
}

func (self *ProcFSDumper) outputErrors(es []error) {
	for _, e := range es {
		self.outputError(e)
	}
}
