package main

import (
	"k8s.io/clientpgo/kubernetes"
	applisters "k8s.io/client.go/listers/apps/v1"

)
type controller struct{
	clientset kubernetes.Interface
	deplister applisters.DeploymentLister
}
