package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	rpio "github.com/stianeikeland/go-rpio"
)

func main() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	pin := rpio.Pin(10)
	pin.Output()
	r := gin.Default()
	r.GET("/@:deg", func(c *gin.Context) {
		deg := c.Param("deg")
		fmt.Println(deg)
		if deg == "on" {
			pin.High()
		} else {
			pin.Low()
		}
		c.JSON(200, "ok!")
	})
	r.Run(":8000")
}
