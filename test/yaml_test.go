package test

import (
	"gopkg.in/yaml.v2"
	"testing"
)

func TestMarshal(t *testing.T) {
	var data = struct {
		Log struct {
			AccessLog string
		}
	}{
		Log: struct{ AccessLog string }{AccessLog: "logs/access.log"},
	}
	bytes, err := yaml.Marshal(&data)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(bytes))
}
