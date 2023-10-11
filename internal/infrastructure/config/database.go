package config

type (
	DBConfig interface {
		GetHost() string
		GetUser() string
		GetPassword() string
		GetDatabaseName() string
		GetPort() int
		GetSSLMode() bool
	}

	dbConfig struct{}
)

func NewDBConfig() DBConfig {
	return dbConfig{}
}

func (dbc dbConfig) GetHost() string {
	return getEnvWithDefaults("DB_HOST", "localhost")
}

func (dbc dbConfig) GetUser() string {
	return getEnvWithDefaults("DB_USER", "mightyblade")
}

func (dbc dbConfig) GetPassword() string {
	return getEnvWithDefaults("DB_PASSWORD", "mightyblade")
}

func (dbc dbConfig) GetDatabaseName() string {
	return getEnvWithDefaults("DB_NAME", "mightyblade")
}

func (dbc dbConfig) GetPort() int {
	return getIntEnvValueWithDefaults("DB_PORT", 5432)
}

func (dbc dbConfig) GetSSLMode() bool {
	return getBoolEnvValueWithDefaults("DB_SSL_MODE", false)
}
