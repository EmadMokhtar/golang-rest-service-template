package config

type Database struct {
	Host     string `fig:"host" default:"localhost"`
	Port     string `fig:"port" default:"3306"`
	User     string `fig:"user" default:"root"`
	Password string `fig:"password" default:""`
	DBName   string `fig:"dbname" default:"golang_rest_service_template"`
	SSLMode  string `fig:"sslmode" default:"disable"`
}
