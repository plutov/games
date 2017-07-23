package config

import (
	"encoding/json"
	"os"
)

// Structure struct
type Structure struct {
	Addr string
}

// Get func
func Get() (*Structure, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	c := new(Structure)
	err = decoder.Decode(&c)

	return c, err
}
