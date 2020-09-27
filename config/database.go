package config

// DatabaseConfig Wrapper struct around postgres configuration.
type DatabaseConfig struct {
	Host string
	Port int
	User string
	Password string
	Database string
}