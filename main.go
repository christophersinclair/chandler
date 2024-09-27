package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(greeting)

	// Get API config
	chandlerConfig, err := getChandlerConfig()
	if err != nil || chandlerConfig == (&ChandlerConfig{}) {
		fmt.Fprintf(os.Stderr, "Faults detected.\nUnrecoverable error: %v\nExiting program...\n\n", err.Error())
		os.Exit(1)
	}

	// Initialize Kubernetes API connection
	k8sAPI, cancel, err := buildK8SAPI()
	if err != nil {
		fmt.Printf("Error establishing a connection to the Kubernetes API: %v\n", err)
		fmt.Println("Unrecoverable error found. Exiting program...\n")
		os.Exit(1)
	}

	defer cancel()
}
