package main

import (
	"encoding/json"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetJobs(clientset *kubernetes.Clientset) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jobs, err := clientset.BatchV1().Jobs("").List(metav1.ListOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var results []JobResult
		for _, job := range jobs.Items {

			var images []string
			for _, container := range job.Spec.Template.Spec.Containers {
				images = append(images, container.Image)
			}

			result := JobResult{
				Name:        job.ObjectMeta.Name,
				Namespace:   job.ObjectMeta.Namespace,
				Completions: *job.Spec.Completions,
				Parallelism: *job.Spec.Parallelism,
				Images:      images,
			}

			results = append(results, result)
		}

		data, err := json.Marshal(GetJobsResult{
			Items:   results,
			Headers: []string{"Name", "Namespace", "Completions", "Parallelism", "Images"},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(data)
	}
}

type GetJobsResult struct {
	Items   []JobResult
	Headers []string
}

type JobResult struct {
	Name        string
	Namespace   string
	Completions int32
	Parallelism int32
	Images      []string
}
