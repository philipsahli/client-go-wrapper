package wrapper

import (
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ExampleGetClientSet() {
	clientset, err := GetClientSet()
	if err != nil {
		log.Fatal(err)
	}

	// Get all Pods
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	fmt.Println(pods)
}
