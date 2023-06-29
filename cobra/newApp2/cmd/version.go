package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name string

func init() {
	rootCmd.AddCommand(versionCmd)
	//cmd.Flags().StringVarP(Name, "name", "n", "", "delete by name")
	versionCmd.Flags().StringVarP(&name, "name", "n", "", "delete by name")

}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
