package main

import (
	"os"
	"testing"
)

func Test_getChandlerConfig_NoEmpty(t *testing.T) {
	config, _ := getChandlerConfig()

	if config == (&ChandlerConfig{}) {
		t.Errorf("Expected populated config, got empty")
	}
}

func Test_getChandlerConfig_DefaultAPIHost(t *testing.T) {
	config, _ := getChandlerConfig()

	if config.APIHost != "0.0.0.0" {
		t.Errorf("Expected APIHost to be %v, got %v", "0.0.0.0", config.APIHost)
	}
}

func Test_getChandlerConfig_DefaultAPIPort(t *testing.T) {
	config, _ := getChandlerConfig()

	if config.APIPort != "41024" {
		t.Errorf("Expected APIPort to be %v, got %v", "41024", config.APIPort)
	}
}

func Test_getChandlerConfig_DefaultCacheFile(t *testing.T) {
	os.Setenv("ADMIN_USER", "foo")
	os.Setenv("ADMIN_PASSWORD", "bar")

	config, _ := getChandlerConfig()

	if config.CacheFilePath != "chandler/tokens.txt" {
		t.Errorf("Expected CacheFilePath to be %v, got %v", "chandler/tokens.txt", config.CacheFilePath)
	}
}

func Test_getChandlerConfig_TokenTTL(t *testing.T) {
	config, _ := getChandlerConfig()

	if config.TokenTTL != "3600" {
		t.Errorf("Expected TokenTTL to be %v, got %v", "3600", config.TokenTTL)
	}
}
