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
package main

import (
	"github.com/marmotedu/gopractise-demo/cobra/demoapp/cmd"
)

func main() {
	//rootCmd.Execute() 是命令的执行入口，其内部会解析 os.Args[1:] 参数列表
	//（默认情况下是这样，也可以通过 Command.SetArgs 方法设置参数），然后遍历命令树，为命令找到合适的匹配项和对应的标志。
	cmd.Execute()
}
