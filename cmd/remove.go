package cmd

import (
	"github.com/LaurenceLiZhixin/dubbogo-cli/pkg/k8s"
	"github.com/spf13/cobra"
)


var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove dubbogo relies on k8s",
	Run: func(cmd *cobra.Command, args []string) {
		k8s.EnvRemove()
	},
}

func init() {
	//InitCmd.Flags().BoolVarP(&kubernetesMode, "kubernetes", "k", false, "Deploy Dapr to a Kubernetes cluster")
	//InitCmd.Flags().BoolVarP(&wait, "wait", "", false, "Wait for Kubernetes initialization to complete")
	//InitCmd.Flags().UintVarP(&timeout, "timeout", "", 300, "The wait timeout for the Kubernetes installation")
	//InitCmd.Flags().BoolVarP(&slimMode, "slim", "s", false, "Exclude placement service, Redis and Zipkin containers from self-hosted installation")
	//InitCmd.Flags().StringVarP(&runtimeVersion, "runtime-version", "", "latest", "The version of the Dapr runtime to install, for example: 1.0.0")
	//InitCmd.Flags().StringVarP(&dashboardVersion, "dashboard-version", "", "latest", "The version of the Dapr dashboard to install, for example: 1.0.0")
	//InitCmd.Flags().StringVarP(&initNamespace, "namespace", "n", "dapr-system", "The Kubernetes namespace to install Dapr in")
	//InitCmd.Flags().BoolVarP(&enableMTLS, "enable-mtls", "", true, "Enable mTLS in your cluster")
	//InitCmd.Flags().BoolVarP(&enableHA, "enable-ha", "", false, "Enable high availability (HA) mode")
	//InitCmd.Flags().String("network", "", "The Docker network on which to deploy the Dapr runtime")
	//InitCmd.Flags().BoolP("help", "h", false, "Print this help message")
	//InitCmd.Flags().StringArrayVar(&values, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	dubbogoRootCMD.AddCommand(removeCmd)
}