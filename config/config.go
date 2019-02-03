// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"time"
)

type Config struct {
	RegistryFile         string        `config:"registry_file"`
	CheckpointMaxUpdates int           `config:"checkpoint_max_updates" validate:"min=1"`
	CheckpointInterval   time.Duration `config:"checkpoint_interval" validate:"min=0,nonzero"`
	Once                 bool          `config:"once"`
	PollingInterval      time.Duration `config:"polling_interval" validate:"min=0,nonzero"`
	WaitClose            time.Duration `config:"wait_close" validate:"min=0"`
}

var DefaultConfig = Config{
	RegistryFile:         ".errlogbeat.yml",
	CheckpointMaxUpdates: 10,
	CheckpointInterval:   15 * time.Second,
	Once:                 false,
	PollingInterval:      15 * time.Second,
	WaitClose:            15 * time.Second,
}
