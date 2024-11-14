package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// 获取用户的主目录
	homePath := homedir.HomeDir()
	if homePath == "" {
		log.Fatalf("failed to get the home directory: %v", os.ErrNotExist)
	}

	// 构建kubeconfig文件路径
	kubeconfig := filepath.Join(homePath, ".kube", "config")

	// 从kubeconfig文件构建API服务器的配置
	// masterUrl是从配置文件中获取的，只传入一个配置文件就可以
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("failed to build config from flags: %v", err)
	}

	// 创建新的Kubernetes客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("failed to create new clientset: %v", err)
	}

	// 列出所有命名空间
	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsForbidden(err) {
			log.Fatalf("permission denied to list namespaces: %v", err)
		}
		log.Fatalf("failed to list namespaces: %v", err)
	}

	// 打印所有命名空间的名字
	for _, namespace := range namespaceList.Items {
		log.Printf("Namespace: %s", namespace.Name)
	}
}
