package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/@:status", func(c *gin.Context) {
		status := c.Param("status")
		n, err := s.Write([]byte(status))
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, n)
	})
	r.Run(":8000")
}
