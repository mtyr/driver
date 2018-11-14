package main

import (
	"github.com/gin-gonic/gin"

	"github.com/hal-ms/driver/boss/cnto"
)

func main() {

	r := gin.Default()
	r.GET("/@:deg", cnto.Boss)
	go r.Run(":8000")

	for {
	}
}
