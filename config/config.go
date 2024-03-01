package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type FullConfig struct {
	Gin   GinConfig   `toml:"gin"`
	Auth  AuthConfig  `toml:"auth"`
	Cache CacheConfig `toml:"cache"`
	Mongo MongoConfig `toml:"mongo"`
}

type GinConfig struct {
	Domain  string   `toml:"domain"`
	Port    string   `toml:"port"`
	Env     string   `toml:"env"`
	Origins []string `toml:"origins"`
}

type AuthConfig struct {
	AccessTokenPub  string `toml:"access_token_pub"`
	AccessTokenTTL  int64  `toml:"access_token_ttl"`
	RefreshTokenPub string `toml:"refresh_token_pub"`
	RefreshTokenTTL int64  `toml:"refresh_token_ttl"`
}

type CacheConfig struct {
	Address  string `toml:"address"`
	Password string `toml:"password"`
	DBID     int    `toml:"id"`
}

type MongoConfig struct {
	URI          string `toml:"uri"`
	DatabaseName string `toml:"database_name"`
}

// GetConfig reads all configurable values
// located in /bin/config.toml in to a FullConfig object
func GetConfig() *FullConfig {
	f := "bin/config.toml"
	if _, err := os.Stat(f); err != nil {
		panic("failed to read config: " + err.Error())
	}

	var conf FullConfig
	if _, err := toml.Decode(f, &conf); err != nil {
		panic("failed to decode config: " + err.Error())
	}

	return &conf
}
