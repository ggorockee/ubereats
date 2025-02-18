package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Vars struct {
	Profile string
}

var (
	Profile string
)

func NewVars() *Vars {
	return &Vars{
		Profile: Profile,
	}
}

func init() {
	Profile = os.Getenv("env")
	if Profile == "" {
		Profile = "dev"
	}
}

func NewConfig(vars *Vars) (*Config, error) {
	var config *Config
	path := filepath.Join(fmt.Sprintf("config/%s/config.yaml", vars.Profile))

	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read %s", path)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}
