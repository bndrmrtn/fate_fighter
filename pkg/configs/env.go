package configs

import (
	"fmt"
	"os"

	"github.com/golobby/dotenv"
)

type dbEnv struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
	Database string `env:"DB_NAME"`
}

func (d dbEnv) GetUrl() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Port, d.Database)
}

type env struct {
	DB dbEnv
}

var Env env

func LoadEnv(path ...string) error {
	p := "./.env"
	if len(path) > 0 {
		p = path[0]
	}

	file, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
	}

	err = dotenv.NewDecoder(file).Decode(&Env)
	if err != nil {
		return fmt.Errorf("failed to parse environment variables: %w", err)
	}

	return nil
}
