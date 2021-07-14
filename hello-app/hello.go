package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/httpbin/hello", func(c *gin.Context){

		c.String(200, "hello, World!")
	})
	
	r.Run(":3000")
}
