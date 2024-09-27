package main

import (
	"testing"
)

func Test_generateConfig_OutsideCluster(t *testing.T) {
	_, err := generateConfig()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if err != nil && err.Error() != "could not generate cluster configuration" {
		t.Errorf("Expected error, got %v", err)
	}
}
