package k8s

import (
	"context"
	"github.com/LaurenceLiZhixin/dubbogo-cli/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

// EnvInit create namespace and deploy resources
func EnvInit()error{
	if err := createNamespace(); err != nil{
		return err
	}
	if err := applyDependency(); err != nil{
		return err
	}
	return nil
}


func createNamespace()error{
	kubeConfigEnv := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigEnv)
	if err != nil {
		return  err
	}
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "dubbogo",
		},
	}
	client := kubernetes.NewForConfigOrDie(config)
	client.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	return nil
}

func applyDependency()error{
	_ , err := utils.RunCmdAndWait("kubectl", "apply", "-f", "")
	return err
}