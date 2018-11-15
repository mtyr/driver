package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/driver/building/cnto"
)

func main() {
	r := gin.Default()
	r.GET("/:scene", cnto.Building)
	r.POST("/:scene", cnto.Scene)
	go r.Run(":8000")
	//go send()

	for {
	}
}
