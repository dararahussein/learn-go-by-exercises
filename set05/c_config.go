package set05

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Address string `json:"address"`
	Debug   bool   `json:"debug"`
}

// LoadConfig should add the path and operation to read/decode errors while
// preserving the original error for errors.Is/errors.As.
func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config %q: %w", path, err)
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("decode config %q: %w", path, err)
	}
	return cfg, nil
}
