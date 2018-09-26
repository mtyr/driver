package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/:deg", func(c *gin.Context) {
		fmt.Println(c.Param("deg"))
		c.JSON(200, "ok!")
	})
	r.Run(":8000")
}
