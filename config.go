package shh

import (
	"io/ioutil"

	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v2"
)

type StoreConfig struct {
	Type     string `yaml:"type" env:"STORE_TYPE"`
	Address  string `yaml:"address" env:"STORE_ADDRESS"`
	Username string `yaml:"username" env:"STORE_USERNAME"`
	Password string `yaml:"password" env:"STORE_PASSWORD"`
}

type Config struct {
	APIPort      string      `yaml:"api_port" env:"API_PORT"`
	RPCPort      string      `yaml:"rpc_port" env:"RPC_PORT"`
	StoreBackend StoreConfig `yaml:"store_backend" env:"STORE_BACKEND"`
}

var defaultConfig = Config{
	APIPort: "8080",
	RPCPort: "7777",
	StoreBackend: StoreConfig{
		Type:    "redis",
		Address: "localhost:6379",
	},
}

func processConfig(filename ...string) (*Config, error) {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	if len(filename) > 0 {
		configFile := filename[0]
		configData, err := ioutil.ReadFile(configFile)
		if err != nil {
			return nil, err
		}

		if err = yaml.Unmarshal(configData, cfg); err != nil {
			return nil, err
		}
	}

	if cfg.APIPort == "" {
		cfg.APIPort = defaultConfig.APIPort
	}
	if cfg.RPCPort == "" {
		cfg.RPCPort = defaultConfig.RPCPort
	}
	if cfg.StoreBackend.Type == "" {
		cfg.StoreBackend.Type = defaultConfig.StoreBackend.Type
	}
	if cfg.StoreBackend.Address == "" {
		cfg.StoreBackend.Address = defaultConfig.StoreBackend.Address
	}
	if cfg.StoreBackend.Username == "" {
		cfg.StoreBackend.Username = defaultConfig.StoreBackend.Username
	}
	if cfg.StoreBackend.Password == "" {
		cfg.StoreBackend.Password = defaultConfig.StoreBackend.Password
	}

	return cfg, nil
}

func LoadConfig(filename ...string) (*Config, error) {
	return processConfig(filename...)
}