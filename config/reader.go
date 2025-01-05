package config

import (
	"encoding/json"
	"os"
)

func ReadConfigFromFile(file string) (conf Config, err error) {
	raw, err := os.ReadFile(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(raw, &conf)
	return
}