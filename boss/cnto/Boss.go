package cnto

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/driver/boss/service"
)

func Boss(c *gin.Context) {
	deg, err := strconv.Atoi(c.Param("deg"))
	fmt.Println(deg)

	if err != nil || deg < 0 || deg > 180 {
		c.JSON(400, "Bad reqest")
		return
	}

	service.Boss.Pow(deg)
	c.JSON(http.StatusOK, "OK")
}
