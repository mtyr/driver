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
	deg0usec   = 750
	deg180usec = 2250
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
	pin.DutyCycle(0, 1280)

	r := gin.Default()
	r.GET("/@:deg", func(c *gin.Context) {
		deg, err := strconv.Atoi(c.Param("deg"))
		fmt.Println(deg)

		if err != nil {
			defer rpio.Close()
			panic(err)
		}

		duty := deg0usec + deg/180*(deg180usec-deg0usec)
		dutyLen := uint32(duty / 20000 * 1280)
		pin.DutyCycle(dutyLen, 1280)

		c.JSON(200, "ok!")
	})
	r.Run(":8000")
}
