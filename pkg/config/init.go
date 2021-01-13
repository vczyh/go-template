package config

import (
	"bufio"
	"go-template/pkg/log"
	"os"
	"strings"
)

func LoadEnvFile(profile string) {
	file, err := os.Open(profile)
	if err != nil {
		log.Warnw("Failed to open .env", "err", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, "=")

		if len(strs) != 2 {
			log.Fatalw("Failed to parse ENV", "env", line)
		}

		key, val := strings.TrimSpace(strs[0]), strings.TrimSpace(strs[1])
		log.Debugw("Setting ENV", "Key", key, "Val", val)
		os.Setenv(key, val)
	}
}
