/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var usageTemplate = fmt.Sprintf(`%s{{if .Runnable}}
  %s{{end}}{{if .HasAvailableSubCommands}}
  %s{{end}}{{if gt (len .Aliases) 0}}

%s
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

%s
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

%s{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  %s {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

%s
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

%s
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

%s{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "%s --help" for more information about a command.{{end}}
`,
	color.CyanString("Usage:"),
	color.GreenString("{{.UseLine}}"),
	color.GreenString("{{.CommandPath}} [command]"),
	color.CyanString("Aliases:"),
	color.CyanString("Examples:"),
	color.CyanString("Available Commands:"),
	color.GreenString("{{rpad .Name .NamePadding }}"),
	color.CyanString("Flags:"),
	color.CyanString("Global Flags:"),
	color.CyanString("Additional help topics:"),
	color.GreenString("{{.CommandPath}} [command]"),
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{

	Use:   "demoapp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	//默认情况下，Cobra 仅解析目标命令上的本地标志，忽略父命令上的本地标志。
	//通过在父命令上启用 Command.TraverseChildren 属性，Cobra 将在执行目标命令之前解析每个命令的本地标志。
	TraverseChildren: true,

	//在执行命令行程序时，我们可能需要对命令参数进行合法性验证，cobra.Command 的 Args 属性提供了此功能
	// 使用内置的验证函数，位置参数多于 2 个则报错
	//Args: cobra.MaximumNArgs(2),

	// 使用自定义验证函数
	//Args: func(cmd *cobra.Command, args []string) error {
	//	if len(args) < 1 {
	//		return errors.New("requires at least one arg")
	//	}
	//	if len(args) > 4 {
	//		return errors.New("the number of args cannot exceed 4")
	//	}
	//	if args[0] != "a" {
	//		return errors.New("first argument must be 'a'")
	//	}
	//	return nil
	//},

	//Run：执行命令时调用的函数，用来编写命令的业务逻辑。
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("root cmd called")
	//},

	//在执行 Run 函数前后，我么可以执行一些钩子函数，其作用和执行顺序如下：
	//PersistentPreRun：在 PreRun 函数执行之前执行，对此命令的子命令同样生效。
	//PreRun：在 Run 函数执行之前执行。
	//PostRun：在 Run 函数执行之后执行。
	//PersistentPostRun：在 PostRun 函数执行之后执行，对此命令的子命令同样生效。

	//以上几个函数都有对应的 <Hooks>E 版本，E 表示 Error，即函数执行出错将会返回 Error，执行顺序不变：
	//如果定义了 <Hooks>E 函数，则 <Hooks> 函数不会执行。比如同时定义了 Run 和 RunE，则只会执行 RunE，不会执行 Run，其他 Hooks 函数同理。
	//PersistentPreRunE
	//PreRunE
	//RunE
	//PostRunE
	//PersistentPostRunE
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("root cmd called")
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	//我们可以使用 cobra.OnInitialize() 来初始化配置文件
	//传递给 cobra.OnInitialize() 的 initConfig 函数将在调用命令的 Execute 方法时运行。
	cobra.OnInitialize(initConfig)

	//如果你对 Cobra 自动生成的帮助命令不满意，我们可以自定义帮助命令或模板
	//Cobra 提供了三个方法来实现自定义帮助命令，后两者也适用于任何子命令。
	//cmd.SetHelpCommand(cmd *Command)
	//cmd.SetHelpFunc(f func(*Command, []string))
	//cmd.SetHelpTemplate(s string)

	//默认情况下，我们可以使用 hugo help command 语法查看子命令的帮助信息，也可以使用 hugo command -h/--help 查看。
	//二者唯一的区别是，使用 help 命令查看帮助信息时会执行钩子函数。
	//我们可以使用 rootCmd.SetHelpCommand 来控制 help 命令输出
	//可以发现，使用 help 命令查看帮助信息输出结果是 rootCmd.SetHelpCommand 中 Run 函数的执行输出。使用 -h 查看帮助信息输出结果是 rootCmd.SetHelpFunc 函数的执行输出
	//rootCmd.SetHelpCommand(&cobra.Command{
	//	Use:    "help",
	//	Short:  "Custom help command",
	//	Hidden: true,
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("1111111111111111111111")
	//	},
	//})
	////使用 rootCmd.SetHelpFunc 来控制 -h/--help 输出。
	//rootCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
	//	fmt.Println("22222222222222222222222222")
	//})

	//rootCmd.SetHelpTemplate 的作用，它用来设置帮助信息模板，支持标准的 Go Template 语法
	//无论使用 help 命令查看帮助信息，还是使用 -h 查看帮助信息，其输出内容都遵循我们自定义的模版格式。
	//rootCmd.SetHelpTemplate(`Custom Help Template:
	//Usage:
	//  {{.UseLine}}
	//Description:
	//  {{.Short}}
	//Commands:
	//{{- range .Commands}}
	//  {{.Name}}: {{.Short}}
	//{{- end}}
	//`)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.demoapp.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//
	//rootCmd.SetUsageTemplate(usageTemplate)
	rootCmd.SetUsageFunc(func(command *cobra.Command) error {
		fmt.Println("11111111")
		return nil
	})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".demoapp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".demoapp")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
