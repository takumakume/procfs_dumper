package procfsdumper

import "github.com/prometheus/procfs"

// Process is a /proc/<pid>
type Process struct {
	// process ID
	PID int
	// command line of a process
	CmdLine []string
	// command name of a process
	Comm string
	// absolute path to the current working directory of a process
	Cwd string
	// process environments from /proc/<pid>/environ
	Environ []string
	// absolute path of the executable command of a process
	Executable string
	// targets of all file descriptors of a process
	// If a file descriptor is not a symlink to a file (like a socket), that value will be the empty string
	FileDescriptorTargets []string
	// IO directory of a process
	IO procfs.ProcIO
	// current soft limits of a process
	Limits procfs.ProcLimits
	// retrieves statistics and configuration for mount points in a process's namespace
	MountStats []*procfs.Mount
	// from /proc/<pid>/ns/* to get the namespaces of which the process is a member
	Namespaces procfs.Namespaces
	// kernel/system statistics read from /proc/[pid]/net/dev
	NetDev procfs.NetDev
	// absolute path to the process's root directory (as set by chroot)
	RootDir string
	// current status information of a process
	Stat procfs.ProcStat
}

// Processes is some process
type Processes []Process
