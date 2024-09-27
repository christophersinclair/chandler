package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type KubernetesAPI struct {
	clientset *kubernetes.Clientset
	dynamic   *dynamic.DynamicClient
	ctx       *context.Context
}

var (
	// Safe-return empty configs
	EmptyConfig    = rest.Config{}
	EmptyClientset = kubernetes.Clientset{}
	EmptyDynamic   = dynamic.DynamicClient{}
	EmptyK8sAPI    = KubernetesAPI{}

	ErrorEmptyConfig    = errors.New("could not generate cluster configuration")
	ErrorEmptyClientset = errors.New("could not generate clientset")
	ErrorEmptyDynamic   = errors.New("could not generate dynamic client")
	ErrorEmptyK8sAPI    = errors.New("could not generate Kubernetes API")
)

func generateConfig() (*rest.Config, error) {
	// Initialize configuration (running on a pod in the cluster)
	// Uses the service account token mounted in the pod at /var/run/secrets
	// Of course, whatever the SA token has access to, the clientset will have access to
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("Error generating in-cluster config: %v\n", err)
		return &EmptyConfig, ErrorEmptyConfig
	}

	return config, nil
}

func generateClientset() (*kubernetes.Clientset, error) {
	// Get cluster config
	config, err := generateConfig()
	if err != nil {
		fmt.Printf("Error generating in-cluster config: %v\n", err)
		return &EmptyClientset, ErrorEmptyConfig
	}

	// Clientset is used to communicate with the API server (for core resources)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error generating clientset: %v\n", err)
		return &EmptyClientset, ErrorEmptyClientset
	}

	return clientset, nil
}

func generateDynamic() (*dynamic.DynamicClient, error) {
	// Get cluster config
	config, err := generateConfig()
	if err != nil {
		fmt.Printf("Error generating in-cluster config: %v\n", err)
		return &EmptyDynamic, ErrorEmptyConfig
	}

	// Dynamic is used to communicate with the API server (for CRDs)
	dynamic, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error generating dynamic clientset: %v\n", err)
		return &EmptyDynamic, ErrorEmptyDynamic
	}

	return dynamic, nil
}

// Build the K8sAPI struct
// NOTE: cancel() should be deferred by the caller to clean up the context
func buildK8SAPI() (KubernetesAPI, context.CancelFunc, error) {
	// Initialize Kubernetes API connection
	clientset, err := generateClientset()
	if err != nil {
		fmt.Printf("Error generating clientset: %v\n", err)
		return EmptyK8sAPI, nil, ErrorEmptyClientset
	}

	// Initialize Dynamic client
	dynamicClient, err := generateDynamic()
	if err != nil {
		fmt.Printf("Error generating dynamic clientset: %v\n", err)
		return EmptyK8sAPI, nil, ErrorEmptyDynamic
	}

	// Create a root context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// Build KubernetesAPI struct
	k8sAPI := KubernetesAPI{
		clientset: clientset,
		dynamic:   dynamicClient,
		ctx:       &ctx,
	}

	// Caller should defer cancel
	return k8sAPI, cancel, nil
}
