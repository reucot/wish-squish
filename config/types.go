package config

type Config struct {
	Mode string `json:"mode" env:"SQUISHY_ENV"`
	Host string `json:"host" env:"SQUISHY_HOST"`
	Port string `json:"port" env:"SQUISHY_PORT"`
	Cert string `json:"cert" env:"SQUISHY_CERT"`
	Log  Log    `json:"log"`
	DB   DB     `json:"db"`
	Smtp Smtp   `json:"smtp"`
}

type Log struct {
	Output string `json:"output" env:"SQUISHY_LOG_OUTPUT"`
	Mode   string `json:"mode" env:"SQUISHY_LOG_MODE"`
}

type DB struct {
	Driver string `json:"driver" env:"SQUISHY_DB_DRIVER"`
	Host   string `json:"host" env:"SQUISHY_DB_URL"`
}

type Smtp struct {
	Host     string `json:"host" env:"SQUISHY_SMTP_HOST"`
	Port     string `json:"port" env:"SQUISHY_SMTP_PORT"`
	User     string `json:"user" env:"SQUISHY_SMTP_USER"`
	Password string `json:"password" env:"SQUISHY_SMTP_PASSWORD"`
}
