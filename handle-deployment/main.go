package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/aloysZY/MyOperatorProjects/client-go-examples/handle-deployment/deployment"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	homePath := homedir.HomeDir()
	if homePath == "" {
		log.Fatalf("failed to get the home directory: %v", os.ErrNotExist)
	}

	kubeconfig := filepath.Join(homePath, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("%s", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("%s", err)
	}
	// Now you can use the clientset to interact with Kubernetes resources

	// 传入namespace获取这个namespace下的所有deployment接口，这个接口有所有deploy方法
	dpClient := clientset.AppsV1().Deployments(corev1.NamespaceDefault)

	log.Printf("Create Deployment.")
	if err := deployment.CreateDeployment(dpClient); err != nil {
		log.Fatalf("%s", err)
	}
	<-time.Tick(1 * time.Minute)

	log.Printf("Update Deployment.")
	if err := deployment.UpdateDeployment(dpClient); err != nil {
		log.Fatalf("%s", err)
	}
	<-time.Tick(1 * time.Minute)

	log.Printf("Delete Deployment.")
	if err := deployment.DeleteDeployment(dpClient); err != nil {
		log.Fatalf("%s", err)
	}
	<-time.Tick(1 * time.Minute)
}
