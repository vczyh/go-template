package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConfigEnv(activeProfile string) error {
	err := LoadEnvFile(activeProfile)
	if err != nil {
		return err
	}
	return nil
}

func LoadEnvFile(profile string) error {
	file, err := os.Open(profile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if index := strings.Index(line, "#"); index != -1 {
			line = line[:index]
		}
		strs := strings.Split(line, "=")
		if len(strs) != 2 {
			return fmt.Errorf("failed to parse environment variable: %s\n", line)
		}

		key, val := strings.TrimSpace(strs[0]), strings.TrimSpace(strs[1])
		if v := os.Getenv(key); v != "" {
			continue
		}
		if strings.HasPrefix(val, `"`) && strings.HasSuffix(val, `"`) {
			v := ""
			if len(val) != 2 {
				v = val[1 : len(val)-1]
			}
			val = v
		}
		err := os.Setenv(key, val)
		if err != nil {
			return err
		}
	}

	return nil
}
