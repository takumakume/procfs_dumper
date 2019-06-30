package procfsdumper

import "github.com/prometheus/procfs"

type Config struct {
	Path        string
	Concurrency int
}

func NewConfig() Config {
	c := Config{}
	defaultConfig(&c)

	return c
}

func defaultConfig(c *Config) {
	c.Path = procfs.DefaultMountPoint
	c.Concurrency = 10
}
