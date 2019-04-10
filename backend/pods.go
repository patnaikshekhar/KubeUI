package main

import (
	"encoding/json"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetPods(clientset *kubernetes.Clientset) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var results []PodResult
		for _, pod := range pods.Items {

			var images []string
			for _, container := range pod.Spec.Containers {
				images = append(images, container.Image)
			}

			result := PodResult{
				Name:      pod.ObjectMeta.Name,
				Namespace: pod.ObjectMeta.Name,
				Images:    images,
				Status:    string(pod.Status.Phase),
			}

			results = append(results, result)
		}

		data, err := json.Marshal(GetPodsResult{
			Items:   results,
			Headers: []string{"Name", "Namespace", "Images", "Status"},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(data)
	}
}

type GetPodsResult struct {
	Items   []PodResult
	Headers []string
}

type PodResult struct {
	Name      string
	Namespace string
	Images    []string
	Status    string
}
