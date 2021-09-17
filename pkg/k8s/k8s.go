package k8s

import (
	"context"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/LaurenceLiZhixin/dubbogo-cli/pkg/config"
	"github.com/LaurenceLiZhixin/dubbogo-cli/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"time"
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

func EnvRemove()error{
	return removeNamespace()
}

func Deploy(image string){

}

func removeNamespace()error{
	kubeConfigEnv := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigEnv)
	if err != nil {
		return  err
	}
	client := kubernetes.NewForConfigOrDie(config)
	client.CoreV1().Namespaces().Delete(context.TODO(), "dubbogo", metav1.DeleteOptions{})
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
	// deploy nacos
	_ , err := utils.RunCmdAndWait("kubectl", "apply", "-f", "https://raw.githubusercontent.com/LaurenceLiZhixin/dubbogo-cli/main/resource/deployment/deployment.yml")

	// todo deploy pixiu

	return err
}


func DeployUserSerivce(appPath, appName, imageName string)error{
	// port-forward nacos
	if err := utils.RunCmd("kubectl", "port-forward", "service/dubbogo-nacos", "8848:nacos", "-n", "dubbogo"); err != nil{
		return err
	}
	time.Sleep(time.Second*3)
	// publish to nacos
	config.PublishDubbogoConfigToNacos(appPath + "/conf/dubbogo.yml", appName, "dubbogo")

	// build docker
	os.Setenv("GO111MODULE", "on")
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "amd64")
	if _, err := utils.RunCmdAndWait("go", "build", "-o", appPath + "/dubbogoApp", appPath+"/cmd"); err != nil{
		return  err
	}

	if _, err := utils.RunCmdAndWait("docker", "build", "--platform", "linux/amd64", "--no-cache", "-t", imageName,"-f", "", appPath); err != nil{
		return  err
	}
	// publish docker
	if _, err := utils.RunCmdAndWait("docker", "push",imageName); err != nil{
		return  err
	}


	// deploy instance
	if err := deployPod(appName, imageName); err != nil{
		return err
	}

	return nil
}


func deployPod(appName, imageName string) error{

	kubeConfigEnv := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigEnv)
	if err != nil {
		return  err
	}
	client := kubernetes.NewForConfigOrDie(config)

	pod := &v1.Pod{
		TypeMeta:metav1.TypeMeta{
			Kind: "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      appName,
			Namespace: "dubbogo",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            appName,
					Image:           imageName,
					ImagePullPolicy: v1.PullAlways,
					Ports: []v1.ContainerPort{
						{
							ContainerPort: 20000,
							Name: "dubbogo",
						},
						{
							ContainerPort: 9090,
							Name: "metrics",
						},
						{
							ContainerPort: 6060,
							Name: "pprof",
						},
					},
					Command: []string{"/dubbogoApp"},
					Env: []v1.EnvVar{
						{
							Name: "DUBBO_GO_CONIFG_PATH",
							Value: "/conf/dubbogo.yml",
						},
						{
							Name: "DUBBO_GO_K8S_MOD",
							Value: "true",
						},
						{
							Name: "DUBBO_GO_APP_NAME",
							Value: appName,
						},
					},
				},
			},
		},
	}

	client.CoreV1().Pods("dubbogo").Create(context.TODO(), pod, metav1.CreateOptions{})
	return nil
}