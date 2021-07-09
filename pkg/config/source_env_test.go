package config

import (
	"os"
	"testing"
)

func TestGetStringFromEnv(t *testing.T) {
	err := os.Setenv("log.path", "111")
	if err != nil {
		t.Fatal(err)
	}

	var env Env
	t.Log(env.getString("log.path"))
}