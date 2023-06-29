//package main
//
//import (
//	"log"
//	"os"
//
//	"github.com/marmotedu/gopractise-demo/log/cuslog"
//)
//
//func main() {
//	cuslog.Info("std log")
//	cuslog.SetOptions(cuslog.WithLevel(cuslog.DebugLevel))
//	cuslog.Debug("change std log to debug level")
//
//	// 输出到文件
//	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		log.Fatalln("create file test.log failed")
//	}
//	defer fd.Close()
//
//	l := cuslog.New(cuslog.WithLevel(cuslog.InfoLevel),
//		cuslog.WithOutput(fd),
//		cuslog.WithFormatter(&cuslog.JsonFormatter{IgnoreBasicFields: false}),
//	)
//	l.Info("custom log with json formatter")
//}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marmotedu/gopractise-demo/log/cuslog"
)

func main() {
	cuslog.Info("std log")
	fmt.Println("11111111111111")
	cuslog.SetOptions(cuslog.WithLevel(cuslog.DebugLevel))

	fmt.Println("222222222")

	//cuslog.Debug("change std log to debug level")
	cuslog.SetOptions(cuslog.WithFormatter(&cuslog.JsonFormatter{IgnoreBasicFields: false}))
	//cuslog.Debug("log in json format")

	fmt.Println("333333333333333333")

	cuslog.Info("another log in json format")

	fmt.Println("44444444444444444444")

	// 输出到文件
	fd, err := os.OpenFile("test.log", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	//localStdoutFile, err := os.OpenFile(localStdoutFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)

	if err != nil {
		log.Printf("create file error: ", err)
		//log.Fatalln("create file test.log failed")
	}
	defer fd.Close()

	l := cuslog.New(cuslog.WithLevel(cuslog.InfoLevel),
		cuslog.WithOutput(fd),
		cuslog.WithFormatter(&cuslog.JsonFormatter{IgnoreBasicFields: false}),
	)
	l.Info("custom log with json formatter")
}
