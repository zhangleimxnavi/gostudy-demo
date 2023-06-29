package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func fileServer(c *gin.Context) {
	path := string(http.Dir("."))
	fmt.Println("path:", path)
	//fileName := path + c.Query("name")
	fileName := path + c.Param("name")
	c.File(fileName)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/file/:name", fileServer)

	// 静态文件服务
	// 显示当前文件夹下面的所有文件 / 或者指定文件
	// 页面返回：服务器当前路径下地文件信息
	r.StaticFS("/showDir", http.Dir("./gin"))

	// 页面返回：服务器./packages目录下地文件信息
	//r.Static("/files", "D:\\workspace\\go-work\\gopractise-demo\\gin")

	// 页面返回：服务器./images/1.jpg图片
	//r.StaticFile("/image", "./images/1.jpg")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
