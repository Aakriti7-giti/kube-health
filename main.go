package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	fmt.Println("🚀 kube-health v2 - K8s DevTools Debugger")
	fmt.Println("----------------------------------------")

	home, _ := os.UserHomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Printf("❌ No cluster found. Start Docker Desktop K8s or run kind create cluster\nError: %v\n", err)
		os.Exit(1)
	}

	clientset, _ := kubernetes.NewForConfig(config)

	fmt.Println("\n  Namespaces:")
	nsList, _ := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	for _, ns := range nsList.Items {
		fmt.Printf("  🔹 %s (%s)\n", ns.Name, ns.Status.Phase)
	}

	fmt.Println("\n  Nodes Health:")
	nodes, _ := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	for _, node := range nodes.Items {
		fmt.Printf("  🔹 %s\n", node.Name)
	}

	fmt.Println("\n Check complete")
}