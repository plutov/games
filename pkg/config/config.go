package config

import (
	"encoding/json"
	"os"
)

// Structure struct
type Structure struct {
	Addr    string
	AddrAPI string
}

// Get func
func Get() (*Structure, error) {
	env := os.Getenv("ENV")
	if 0 == len(env) {
		env = "dev"
	}

	file, err := os.Open("config-" + env + ".json")
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	c := new(Structure)
	err = decoder.Decode(&c)

	return c, err
}
