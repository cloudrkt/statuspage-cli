package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Manipulate statuspage components",
}

var componentListCmd = &cobra.Command{
	Use:     "list",
	Example: "statuspage component list",
	Short:   "list component",
	Long:    `Lists all components`,
	Run: func(cmd *cobra.Command, args []string) {
		components, err := app.Client.GetAllComponents()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, component := range components {
			fmt.Println(&component)
		}

	},
}

func init() {
	RootCmd.AddCommand(componentCmd)
	componentCmd.AddCommand(componentListCmd)
}
