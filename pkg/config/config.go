package config

import (
	"fmt"
)

var config = &Config{}

type Config struct {
	httpPort      int
	httpMode      string
	httpAccessLog string
	httpErrorLog  string

	logLevel string
	logPath  string
}

// source maybe nil
func ReadFromMultiSources(env, file Source) error {
	var c Config

	if err := c.getKeys(env, file); err != nil {
		return err
	}
	config = &c
	return nil
}

func (c *Config) getKeys(env, file Source) error {
	if file == nil {
		return fmt.Errorf("file source can't be nil")
	}

	// http
	c.httpPort = file.getInt("http.port")
	c.httpMode = file.getString("http.mode")
	c.httpAccessLog = file.getString("http.access-log")
	c.httpErrorLog = file.getString("http.error-log")

	// log
	c.logLevel = file.getString("log.level")
	c.logPath = file.getString("log.path")

	return nil
}

func Info() string {
	return fmt.Sprintf("%+v", *config)
}

func GetHttpPort() int {
	return config.httpPort
}

func GetHttpMode() string {
	return config.httpMode
}

func GetHttpAccessLog() string {
	return config.httpAccessLog
}

func GetHttpErrorLog() string {
	return config.httpErrorLog
}

func GetLogLevel() string {
	return config.logLevel
}

func GetLogPath() string {
	return config.logPath
}
