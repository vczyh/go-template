package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	viper *viper.Viper
}

func New(file string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(file)

	if err := v.ReadInConfig(); err != nil {
		return nil,err
	}

	return &Config{
		viper: v,
	}, nil
}

// Set
func (c *Config) Set(key string, value interface{}) {
	c.viper.Set(key, value)
}

func (c *Config) SetDefault(key string, value interface{}) {
	c.viper.SetDefault(key, value)
}

// Get
func (c *Config) Get(key string) interface{} {
	return c.viper.Get(key)
}

func (c *Config) GetString(key string) string {
	return c.viper.GetString(key)
}

func (c *Config) GetInt(key string) int {
	return c.viper.GetInt(key)
}

func (c *Config) GetFloat64(key string) float64 {
	return c.viper.GetFloat64(key)
}

func (c *Config) GetBool(key string) bool {
	return c.viper.GetBool(key)
}

// time
func (c *Config) GetDuGetTimeration(key string) time.Time {
	return c.viper.GetTime(key)
}

func (c *Config) GetDuration(key string) time.Duration {
	return c.viper.GetDuration(key)
}

// slice
func (c *Config) GetStringSlice(key string) []string {
	return c.viper.GetStringSlice(key)
}

func (c *Config) GetIntSlice(key string) []int {
	return c.viper.GetIntSlice(key)
}
