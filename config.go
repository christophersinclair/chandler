package main

import (
	"fmt"
	"os"
)

type ChandlerConfig struct {
	APIHost       string
	APIPort       string
	CacheFilePath string
	TokenTTL      string
}

var (
	NullConfig = ChandlerConfig{}
)

func getChandlerConfig() (*ChandlerConfig, error) {
	apiHost := os.Getenv("API_HOST")
	if apiHost == "" {
		fmt.Println("Could not find API_HOST set. Assuming 0.0.0.0...")
		apiHost = "0.0.0.0"
	}

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		fmt.Println("Could not find API_PORT set. Assuming 41024...")
		apiPort = "41024"
	}

	cacheFilePath := os.Getenv("CACHE_FILE")
	if cacheFilePath == "" {
		fmt.Println("Could not find CACHE_FILE set. Assuming chandler/tokens.txt")
		cacheFilePath = "chandler/tokens.txt"
	}

	tokenTTL := os.Getenv("TOKEN_TTL")
	if tokenTTL == "" {
		fmt.Println("Could not find TOKEN_TTL set. Assuming 3600s...")
		tokenTTL = "3600"
	}

	return &ChandlerConfig{
		APIHost:       apiHost,
		APIPort:       apiPort,
		CacheFilePath: cacheFilePath,
		TokenTTL:      tokenTTL,
	}, nil

}
