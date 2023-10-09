package routers

import (
	"github.com/elysiamori/assignment3-09/requestapi"
	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	r := gin.Default()
	r.PUT("/weather", requestapi.UpdateWeatherData)

	return r
}
