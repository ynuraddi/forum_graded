package config

import "sync"

type Config struct {
	HTTP struct {
		HOST string `json:"host"`
		PORT string `json:"port"`
	} `json:"http"`

	DATABASE struct {
		HOST string `json:"host"`
		PORT string `json:"port"`
	} `json:"database"`
}

var once sync.Once

func Init() (config *Config) {
	once.Do(func() {
	})
	return
}
