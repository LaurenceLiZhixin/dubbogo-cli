// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation and Dapr Contributors.
// Licensed under the MIT License.
// ------------------------------------------------------------

package cmd

import (
	"github.com/LaurenceLiZhixin/dubbogo-cli/pkg/k8s"
	"github.com/spf13/cobra"
)
var (
	imageName    string
	appName      string
	appPath      string
)

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy dubbogo user service",
	Run: func(cmd *cobra.Command, args []string) {
		imageName := "image"

		// build docker

		// publish docker

		// set

		// deploy to k8s
		if err := k8s.DeployUserSerivce(appPath, appName, imageName); err != nil{
			panic(err)
		}
	},
}

func init() {
	DeployCmd.Flags().StringVarP(&imageName, "imageName", "", "latest", "image name")
	DeployCmd.Flags().StringVarP(&appName, "appName", "", "latest", "container name")
	DeployCmd.Flags().StringVarP(&appPath, "appPath", "", "latest", "container name")
	//InitCmd.Flags().StringVarP(&initNamespace, "namespace", "n", "dapr-system", "The Kubernetes namespace to install Dapr in")
	//InitCmd.Flags().BoolVarP(&enableMTLS, "enable-mtls", "", true, "Enable mTLS in your cluster")
	//InitCmd.Flags().BoolVarP(&enableHA, "enable-ha", "", false, "Enable high availability (HA) mode")
	//InitCmd.Flags().String("network", "", "The Docker network on which to deploy the Dapr runtime")
	//InitCmd.Flags().BoolP("help", "h", false, "Print this help message")
	//InitCmd.Flags().StringArrayVar(&values, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	dubbogoRootCMD.AddCommand(DeployCmd)
}
