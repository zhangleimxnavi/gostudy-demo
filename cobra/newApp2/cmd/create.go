package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	apptype    string
	appversion string
	appname    string
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&appname, "name", "n", "", "delete by name")

	//rootCmd.PersistentFlags().StringVar(&apptype, "apptype", "", "config file (default is $HOME/.cobra.yaml)")
	//rootCmd.PersistentFlags().StringVarP(&appversion, "appversion", "d", "", "base project directory eg. github.com/spf13/")
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Print the createCmd number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	//Args:  cobra.MinimumNArgs(1), // 使用内置的验证函数
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return fmt.Errorf("invalid color specified: %s", args[0])
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create aHugo Static Site Generator v0.9 -- HEAD")
	},
}
