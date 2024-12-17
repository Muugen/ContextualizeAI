package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	IgnorePatterns []string `json:"ignore_patterns"`
	OutputFile     string   `json:"output_file"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	if cfg.OutputFile == "" {
		cfg.OutputFile = "get_code_context.txt"
	}

	return &cfg, nil
}
