package info

import (
	"go-template/pkg/log"
	"os"
)

func PrintInfo() error {
	if err := PWD(); err != nil {
		return err
	}
	return nil
}

func PWD() error {
	pwd, err := os.Getwd()
	log.Info("pwd:", pwd)
	return err
}
