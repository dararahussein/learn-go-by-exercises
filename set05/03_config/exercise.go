package config

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
	return Config{}, fmt.Errorf("TODO: load config %q", path)
}

// Keep the required packages visible while LoadConfig is a learner-owned stub.
var (
	_ = json.Unmarshal
	_ = os.ReadFile
)
