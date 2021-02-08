package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var C = struct {
	Http struct {
		Port int
		Mode string
	}
	Log struct {
		App struct {
			Path  string
			Level string
		}
		AccessLog string `yaml:"access-log"`
	}
}{}

func LoadConfig(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err

	}

	err = yaml.Unmarshal(bytes, &C)
	if err != nil {
		return err
	}
	return nil
}
