package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

var (
//flagvar = pflag.Int("flagname", 1234, "help message for flagname")
//flagvar = pflag.IntP("flagname", "f", 1234, "help message for flagname")
)

func main() {
	
	var version bool
	var name string
	
	//1、调用NewFlagSet 创建一个 FlagSet
	//flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	//flagSet.BoolVar(&version, "version", true, "Print version information and quit.")
	//flagSet.StringVar(&name, "name", "shuaige", "please input a userName")
	//flagSet.Parse(os.Args)
	//
	//fmt.Println("version: ", version)
	//fmt.Println("name: ", name)
	////
	//getBool, err := flagSet.GetBool("version")
	//if err != nil {
	//	return
	//}
	//getString, err := flagSet.GetString("name")
	//if err != nil {
	//	return
	//}
	////
	//fmt.Println("version: ", getBool)
	//fmt.Println("name: ", getString)
	
	//2、使用全局 FlagSet
	pflag.BoolVar(&version, "common.version", true, "Print version information and quit.")
	pflag.StringVar(&name, "common.name", "dashuaige", "input your name please")
	
	//var name = pflag.String("name", "colin", "Input Your Name")
	//var name = pflag.StringP("name", "n", "colin", "Input Your Name")
	
	var port = pflag.IntP("hello.port", "p", 1234, "help message")
	var port2 = pflag.IntP("hello.port2", "d", 1234, "help message")
	//pflag.Lookup("flagname").NoOptDefVal = "4321"
	////
	//pflag.CommandLine.MarkDeprecated("version", "please use --new_version instead")
	//
	//pflag.CommandLine.MarkShorthandDeprecated("port", "please use --port instead")
	//
	//pflag.Lookup("flagname").NoOptDefVal = "250"
	
	//fmt.Println("name: ", *name)
	//fmt.Println("name: ", str)
	
	pflag.Parse()
	fmt.Println("version: ", version)
	fmt.Println("name: ", name)
	
	//
	fmt.Println("port is: ", *port)
	fmt.Println("port is: ", *port2)
	//fmt.Printf("flagvar is: %v\n", *flagvar)
	fmt.Printf("argument number is: %v\n", pflag.NArg())
	fmt.Printf("argument list is: %v\n", pflag.Args())
	fmt.Printf("the first argument is: %v\n", pflag.Arg(0))
	fmt.Printf("the second argument is: %v\n", pflag.Arg(1))
	
}
