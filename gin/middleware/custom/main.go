package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置变量example
		c.Set("example", "12345")

		// 请求之前
		//c.Next() 之前的操作是在 Handler 执行之前就执行；
		//c.Next() 之后的操作是在 Handler 执行之后再执行；
		fmt.Println("Hello Before;")
		c.Next()
		fmt.Println("Hello After;")

		// 请求之后
		latency := time.Since(t)
		log.Print(latency)

		// 访问我们发送的状态
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.Use(gin.BasicAuth(gin.Accounts{"admin": "admin"}))

	// RequestID
	r.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return "testRequestid"
		},
	}))

	// 跨域
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println("i am test function")
		log.Println(example)

	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
