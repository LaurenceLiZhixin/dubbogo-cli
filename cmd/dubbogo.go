package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var dubbogoRootCMD = &cobra.Command{

}

func Init(){
	if err := dubbogoRootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
