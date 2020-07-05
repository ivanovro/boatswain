package main

import (
	"flag"
	"fmt"
	"github.com/alauda/helm-crds/pkg/apis/app/v1alpha1"
	"github.com/alauda/helm-crds/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	captainClient := versioned.NewForConfigOrDie(kubeConfig())
	handleHttpRequests(captainClient)
}

func handleHttpRequests(client *versioned.Clientset) {
	http.HandleFunc("/", rootHandler(client))
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func rootHandler(client *versioned.Clientset) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		chartName := r.RequestURI[1:]
		switch r.Method {
		case "POST":
			helmRequest := makeHelmRequest(chartName)
			_, err := captainRequestCreate(client, helmRequest)
			if err != nil {
				replyError(w, err)
			} else {
				confirmRequest(w, chartName, "installation")
			}
		case "DELETE":
			err := captainRequestDelete(client, chartName)
			if err != nil {
				replyError(w, err)
			} else {
				confirmRequest(w, chartName, "deletion")
			}
		default:
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message": "Can't find method requested"}`))
		}
	}
}

func confirmRequest(w http.ResponseWriter, chartName string, operation string) {
	w.WriteHeader(http.StatusOK)
	status := fmt.Sprintf(`{"status": "'%s' %s in progress"}`, chartName, operation)
	_, _ = w.Write([]byte(status))
}

func replyError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	status := fmt.Sprintf(`{"status": "%s"}`, err)
	_, _ = w.Write([]byte(status))
}

func captainRequestCreate(c *versioned.Clientset, helmRequest *v1alpha1.HelmRequest) (*v1alpha1.HelmRequest, error) {
	return c.AppV1alpha1().HelmRequests("default").Create(helmRequest)
}
func captainRequestDelete(c *versioned.Clientset, chartName string) error {
	return c.AppV1alpha1().HelmRequests("default").Delete(chartName, nil)
}

func makeHelmRequest(chartName string) *v1alpha1.HelmRequest {
	helmRequest := &v1alpha1.HelmRequest{
		ObjectMeta: v1.ObjectMeta{Name: chartName},
		Spec: v1alpha1.HelmRequestSpec{
			Chart:     "stable/" + chartName,
			Namespace: "default",
		},
	}
	return helmRequest
}

func kubeConfig() *rest.Config {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	return config
}
