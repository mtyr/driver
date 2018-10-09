package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/driver/rpio"
)

func main() {
	err := rpio.Open()

	if err != nil {
		defer rpio.Close()
		panic(err)
	}

	pin := rpio.Pin(10)
	pin.Mode(rpio.Pwm)
	pin.Freq(64000)
	pin.DutyCycle(0, 32)

	r := gin.Default()
	r.GET("/@:deg", func(c *gin.Context) {
		deg, err := strconv.Atoi(c.Param("deg"))

		if err != nil {
			defer rpio.Close()
			panic(err)
		}

		fmt.Println(deg)
		for {
			for i := uint32(0); i < 32; i++ {
				pin.DutyCycle(i, 32)
				time.Sleep(time.Second / 32)
			}
			for i := uint32(32); i > 0; i-- {
				pin.DutyCycle(i, 32)
				time.Sleep(time.Second / 32)
			}
		}
		c.JSON(200, "ok!")
	})
	r.Run(":8000")
}
