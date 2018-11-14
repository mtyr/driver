package cnto

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/driver/building/service"
)

func Building(c *gin.Context) {
	service.Building.Bar()
	c.JSON(http.StatusOK, "OK")
}
