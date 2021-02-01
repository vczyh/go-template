package config

import (
	"bufio"
	"go-template/pkg/log"
	"os"
	"strings"
)

func LoadEnvFile(profile string) error {
	file, err := os.Open(profile)
	if err != nil {
		log.Warnw("Failed to open .env", "err", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, "=")
		if len(strs) != 2 {
			log.Errorw("failed to parse environment value", "env", line)
			return err
		}

		key, val := strings.TrimSpace(strs[0]), strings.TrimSpace(strs[1])
		if v := os.Getenv(key); v != "" {
			log.Warnw("environment value has existed so not set", "key", key, "currentVal", v, "configVal", val)
			continue
		}
		err := os.Setenv(key, val)
		if err != nil {
			log.Errorw("setting environment value Error", "key", key, "val", val)
			return err
		}
		log.Debugw("setting environment value", "key", key, "val", val)
	}

	return nil
}
