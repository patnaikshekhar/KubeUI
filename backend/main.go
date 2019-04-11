package main

import (
	"net/http"
	"os"

	"k8s.io/client-go/rest"

	"github.com/gorilla/mux"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	configLocation := os.Getenv("CONFIG_LOCATION")

	var config *rest.Config
	var err error

	if configLocation == "" {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", configLocation)
	}

	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public/static"))))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})
	r.HandleFunc("/api/pods", GetPods(clientset)).Methods("GET")
	r.HandleFunc("/api/deployments", GetDeployments(clientset)).Methods("GET")
	r.HandleFunc("/api/services", GetServices(clientset)).Methods("GET")
	r.HandleFunc("/api/jobs", GetJobs(clientset)).Methods("GET")
	r.HandleFunc("/api/cronjobs", GetCronJobs(clientset)).Methods("GET")
	r.HandleFunc("/api/deployments", CreateDeployment(clientset)).Methods("POST")

	http.ListenAndServe(":80", r)
}
