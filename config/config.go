// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"time"
)

type Config struct {
	RegistryFile    string        `config:"registry_file"`
	Once            bool          `config:"once"`
	PollingInterval time.Duration `config:"polling_interval" validate:"min=0,nonzero"`
}

var DefaultConfig = Config{
	RegistryFile:    ".errlogbeat.yml",
	Once:            false,
	PollingInterval: 15 * time.Second,
}
