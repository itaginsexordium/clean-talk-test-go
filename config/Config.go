package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config struct used for configuration of app with env variables
type Config struct {
	MemcacheURL      string `envconfig:"MEMCACHE_URL" required:"false" default:"0.0.0.0:7800"`
	HTTPBindAddr     string `envconfig:"HTTP_BIND_ADDR" required:"false" default:"8080"`
	GeoIpPath       string `envconfig:"GEOIP_PATH" required:"false" default:"source/db/GeoLite2-Country.mmdb"`
} 

func Get() (*Config, error) {
	c := &Config{}
	err := envconfig.Process("", c)
	return c, err
}
