package main

import (
	"encoding/json"
	"log"
	"net/http"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func GetDeployments(clientset *kubernetes.Clientset) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		deployments, err := clientset.AppsV1().Deployments("").List(metav1.ListOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var results []DeploymentResult
		for _, deployment := range deployments.Items {
			result := DeploymentResult{
				Name:              deployment.ObjectMeta.Name,
				Namespace:         deployment.ObjectMeta.Namespace,
				Replicas:          *deployment.Spec.Replicas,
				ReadyReplicas:     deployment.Status.ReadyReplicas,
				AvailableReplicas: deployment.Status.AvailableReplicas,
			}

			results = append(results, result)
		}

		data, err := json.Marshal(GetDeploymentsResult{
			Items:   results,
			Headers: []string{"Name", "Namespace", "Replicas", "ReadyReplicas", "AvailableReplicas"},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(data)
	}
}

func CreateDeployment(clientset *kubernetes.Clientset) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("deployments.CreateDeployment started")
		var request CreateDeploymentRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			log.Printf("deployments.CreateDeployment Error: ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		deploymentsClient := clientset.AppsV1().Deployments(request.Namespace)

		deployment := &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name: request.Name,
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: &request.Replicas,
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"app": request.Name,
					},
				},
				Template: apiv1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app": request.Name,
						},
					},
					Spec: apiv1.PodSpec{
						Containers: []apiv1.Container{
							{
								Name:  request.Name,
								Image: request.Image,
							},
						},
					},
				},
			},
		}

		_, err = deploymentsClient.Create(deployment)
		if err != nil {
			log.Printf("deployments.CreateDeployment Error: ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("deployments.CreateDeployment Deployment created successfully")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Created"))
	}
}

type GetDeploymentsResult struct {
	Items   []DeploymentResult
	Headers []string
}

type DeploymentResult struct {
	Name              string
	Namespace         string
	Replicas          int32
	ReadyReplicas     int32
	AvailableReplicas int32
}

type CreateDeploymentRequest struct {
	Name      string
	Namespace string
	Image     string
	Replicas  int32
}
