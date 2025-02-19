package config

type Config struct {
	Server Server `mapstructure:"server"`
	Infra  Infra  `mapstructure:"infra"`
	Secret Secret `mapstructure:"secret"`
}

type Secret struct {
	Jwt string `mapstructure:"jwt"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type Infra struct {
	DB DB `mapstructure:"db"`
}

type DB struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	DBName   string `mapstructure:"dbName"`
	Port     string `mapstructure:"port"`
}
