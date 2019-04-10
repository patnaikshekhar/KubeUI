package main

import (
	"encoding/json"
	"log"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetCronJobs(clientset *kubernetes.Clientset) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cronjobs, err := clientset.BatchV1beta1().CronJobs("").List(metav1.ListOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var results []CronJobResult
		for _, cronjob := range cronjobs.Items {

			log.Print("Fetching images")
			var images []string
			for _, container := range cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers {
				images = append(images, container.Image)
			}

			log.Print("Setting Cron Result")
			result := CronJobResult{
				Name:      cronjob.ObjectMeta.Name,
				Namespace: cronjob.ObjectMeta.Namespace,
				Schedule:  cronjob.Spec.Schedule,
				Images:    images,
			}

			results = append(results, result)
		}

		data, err := json.Marshal(GetCronJobsResult{
			Items:   results,
			Headers: []string{"Name", "Namespace", "Schedule", "Images"},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(data)
	}
}

type GetCronJobsResult struct {
	Items   []CronJobResult
	Headers []string
}

type CronJobResult struct {
	Name      string
	Namespace string
	Schedule  string
	Images    []string
}
