// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	RegistryFile string `config:"registry_file"`
}

var DefaultConfig = Config{
	RegistryFile: ".errlogbeat.yml",
}
