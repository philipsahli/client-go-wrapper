// Package wrapper provides a way to setup a Clientset quickly.
package wrapper

import (
	"log"

	v1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodList []v1.Pod

func GetPods() (PodList, error) {
	podList := make(PodList, 0)
	clientset, err := GetClientSet()
	if err != nil {
		return nil, err
	}
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error in getting pods: %s", err)
	}

	for _, pod := range pods.Items {
		podList = append(podList, pod)
	}

	return podList, nil
}
