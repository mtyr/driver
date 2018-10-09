package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/driver/rpio"
)

const (
	freq = 64000
)

const (
	deg0msec   = 650
	deg180msec = 2350
)

func main() {
	err := rpio.Open()

	if err != nil {
		defer rpio.Close()
		panic(err)
	}

	pin := rpio.Pin(19)
	pin.Pwm()
	pin.Freq(freq)
	pin.DutyCycle(0, 32)

	r := gin.Default()
	r.GET("/@:deg", func(c *gin.Context) {
		deg, err := strconv.Atoi(c.Param("deg"))

		if err != nil {
			defer rpio.Close()
			panic(err)
		}

		dutyLen := freq / (deg0msec + deg/180*(deg180msec-deg0msec))
		fmt.Println(deg)
		pin.DutyCycle(dutyLen, 32)

		c.JSON(200, "ok!")
	})
	r.Run(":8000")
}
