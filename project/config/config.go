package config

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	configs Configs
	once    sync.Once
)

// Configs consists all configuration
type Configs struct {
	Migrate bool
	Server  ServerConfig
	Mysql   MySQLConfig
}

// ServerConfig consists server configuration
type ServerConfig struct {
	Address        string
	RequestTimeout int
}

// MySQLConfig consists MySQL database configuration
type MySQLConfig struct {
	Host            string
	Port            string
	User            string
	Pass            string
	Name            string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime int
}

// GetConfig return Configs object read from config.json file
func GetConfig() *Configs {
	once.Do(func() {
		conf := viper.New()
		conf.SetConfigFile("./config/config.json")

		err := conf.ReadInConfig()
		if err != nil {
			log.Fatalf("failed to read config file: %s", err)
		}

		if err := conf.Unmarshal(&configs); err != nil {
			log.Fatalf("failed to unmarshal config: %s", err)
		}

		args := os.Args[1:]
		if len(args) > 0 {
			if args[0] == "migrate" {
				configs.Migrate = true
			}
		}
	})

	return &configs
}
