package main

import (
	"encoding/json"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetServices(clientset *kubernetes.Clientset) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		services, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var results []ServiceResult
		for _, service := range services.Items {
			result := ServiceResult{
				Name:           service.ObjectMeta.Name,
				Namespace:      service.ObjectMeta.Namespace,
				Type:           string(service.Spec.Type),
				ClusterIP:      service.Spec.ClusterIP,
				LoadBalancerIP: service.Spec.LoadBalancerIP,
			}

			results = append(results, result)
		}

		data, err := json.Marshal(GetServicesResult{
			Items:   results,
			Headers: []string{"Name", "Namespace", "Type", "ClusterIP", "LoadBalancerIP"},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(data)
	}
}

type GetServicesResult struct {
	Items   []ServiceResult
	Headers []string
}

type ServiceResult struct {
	Name           string
	Namespace      string
	Type           string
	ClusterIP      string
	LoadBalancerIP string
}
