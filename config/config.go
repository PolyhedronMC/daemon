package config

import (
	"github.com/spf13/viper"
)

// DaemonConfig Wrapper struct around daemon configuration.
type DaemonConfig struct {
	Docker DockerConfig
	Database DatabaseConfig
}

// DockerConfig Wrapper struct around Docker engine configuration.
type DockerConfig struct {
	Host string
	Port int
}

// DatabaseConfig Wrapper struct around postgres configuration.
type DatabaseConfig struct {
	Host string
	Port int
	User string
	Password string
	Database string
}

// GetConfig Create and read the configuration file.
func GetConfig() DaemonConfig {
	viper.SetConfigName("daemon")
	viper.SetConfigType("toml")
	// viper.AddConfigPath("/etc/polyhedron/daemon")
	viper.AddConfigPath(".")

	// Docker configuration
	viper.SetDefault("docker.host", "172.0.0.1")
	viper.SetDefault("docker.port", 80)

	// PostgreSQL configuration
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", 54321)
	viper.SetDefault("postgres.user", "postgres")
	viper.SetDefault("postgres.password", "")
	viper.SetDefault("postgres.database", "polyhedron")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = viper.SafeWriteConfig()
			if err != nil {
				panic("failed to write sample config")
			}
		} else {
			panic("failed to read config")
		}
	}

	return DaemonConfig {
		Docker: DockerConfig {
			Host: viper.GetString("docker.host"),
			Port: viper.GetInt("docker.port"),
		},
		Database: DatabaseConfig {
			Host: viper.GetString("postgres.host"),
			Port: viper.GetInt("postgres.port"),
			User: viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
			Database: viper.GetString("postgres.database"),
		},
	}
}
