package config

import (
	"os"
	"strconv"
)

type Env struct{}

func (e *Env) getString(key string) string {
	return os.Getenv(key)
}

func (e *Env) getInt(key string) int {
	val, _ := strconv.Atoi(os.Getenv(key))
	return val
}
