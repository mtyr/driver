package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/driver/building/cnto"
)

func main() {
	r := gin.Default()
	r.GET("/", cnto.Building)
	go r.Run(":8000")
	//go send()

	for {
	}
}
