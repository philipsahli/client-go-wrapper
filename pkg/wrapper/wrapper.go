// Package wrapper provides a way to setup a Clientset quickly.
package wrapper

import (
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// GetClientSet returns a clientset to work with Kubernetes API
func GetClientSet() (*kubernetes.Clientset, error) {
	var kubeconfig string

	// Look for KUBECONFIG environment variable
	if kubeconfig = os.Getenv("KUBECONFIG"); kubeconfig != "" {
		if _, err := os.Stat(kubeconfig); err != nil {
			log.Fatalf("file '%s' referenced in KUBECONFIG variable but does not exist", kubeconfig)
		}
	} else {
		// Look for config file in homedirectory
		if home := homeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		}
	}

	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
