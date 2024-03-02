package config

import (
	"bufio"
	"github.com/goccy/go-yaml"
	"io"
	"os"
)

type FullConfig struct {
	Gin   GinConfig   `yaml:"gin"`
	Auth  AuthConfig  `yaml:"auth"`
	Cache CacheConfig `yaml:"cache"`
	Mongo MongoConfig `yaml:"mongo"`
}

type GinConfig struct {
	Domain  string   `yaml:"domain"`
	Port    string   `yaml:"port"`
	Env     string   `yaml:"env"`
	Origins []string `yaml:"origins"`
}

type AuthConfig struct {
	AccessTokenPub  string `yaml:"access_token_pub"`
	AccessTokenTTL  int    `yaml:"access_token_ttl"`
	RefreshTokenPub string `yaml:"refresh_token_pub"`
	RefreshTokenTTL int    `yaml:"refresh_token_ttl"`
}

type CacheConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DBID     int    `yaml:"dbid"`
}

type MongoConfig struct {
	URI          string `yaml:"uri"`
	DatabaseName string `yaml:"database_name"`
}

// GetConfig reads all configurable values
// located in /bin/config.toml in to a FullConfig object
func GetConfig() *FullConfig {
	file, err := os.Open("bin/config.yaml")
	if err != nil {
		panic("failed to read config: " + err.Error())
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("failed to close config buffer: " + err.Error())
		}
	}(file)

	stat, err := file.Stat()
	if err != nil {
		panic("failed to parse config: " + err.Error())
	}

	b := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(b)
	if err != nil && err != io.EOF {
		panic("failed to read config byte data: " + err.Error())
	}

	var conf FullConfig
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		panic("failed to unmarshal config yaml data: " + err.Error())
	}

	return &conf
}
