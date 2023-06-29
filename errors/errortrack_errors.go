//package main
//
//import (
//	"fmt"
//
//	"github.com/marmotedu/errors"
//
//	code "github.com/marmotedu/sample-code"
//)
//
//func main() {
//	if err := getUser(); err != nil {
//		//fmt.Printf("%+v\n", err)
//		fmt.Printf("%+v\n", err)
//	}
//}
//
//func getUser() error {
//	if err := queryDatabase(); err != nil {
//		return errors.Wrap(err, "1111111get user failed.111111")
//	}
//
//	return nil
//}
//
//func queryDatabase() error {
//	return errors.WithCode(code.ErrDatabase, "0000000user 'Lingfei Kong' not found.00000000")
//}
