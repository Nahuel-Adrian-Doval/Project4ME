package models

type DBConfig struct {
	DATABASE_PORT string

	DATABASE_HOST     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string

	DATABASE_SSLMODE  string
	DATABASE_TIMEZONE string
}

type ServerConfig struct {
	SERVER_PORT   string
	SERVER_PREFIX string
}
