package config

import (
	"github.com/spf13/viper"
	"time"
)

type File struct {
	viper *viper.Viper
}

func LoadFile(file string) (*File, error) {
	v := viper.New()
	v.SetConfigFile(file)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return &File{
		viper: v,
	}, nil
}

func (c *File) Get(key string) interface{} {
	return c.viper.Get(key)
}

func (c *File) getString(key string) string {
	return c.viper.GetString(key)
}

func (c *File) getInt(key string) int {
	return c.viper.GetInt(key)
}

func (c *File) GetFloat64(key string) float64 {
	return c.viper.GetFloat64(key)
}

func (c *File) GetBool(key string) bool {
	return c.viper.GetBool(key)
}

func (c *File) GetTime(key string) time.Time {
	return c.viper.GetTime(key)
}

func (c *File) GetDuration(key string) time.Duration {
	return c.viper.GetDuration(key)
}

func (c *File) GetStringSlice(key string) []string {
	return c.viper.GetStringSlice(key)
}

func (c *File) GetIntSlice(key string) []int {
	return c.viper.GetIntSlice(key)
}
