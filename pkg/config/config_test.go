package config

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(pwd)
	m.Run()
}

func TestPath(t *testing.T) {
	c, err := New("./config.yml")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(c.GetString("http.port"))
}
