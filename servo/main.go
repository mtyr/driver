package main

import (
	"fmt"
	"strconv"
	"time"

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

		if err != nil || deg < 0 || deg > 180 {
			c.JSON(400, "Bad reqest")
			return
		}

		duty := deg0usec + deg*(deg180usec-deg0usec)/180
		dutyLen := uint32(duty * 1280 / 20000)
		fmt.Println(dutyLen)

		pin.DutyCycle(dutyLen, 1280)
		time.Sleep(time.Second / 1280)

		c.JSON(200, "ok!")
	})
	r.Run(":8000")
}
