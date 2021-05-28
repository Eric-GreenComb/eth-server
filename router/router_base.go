package router

import (
	"github.com/gin-gonic/gin"

	myhandler "github.com/Eric-GreenComb/eth-server/handler"
)

// SetupBaseRouter SetupBaseRouter
func SetupBaseRouter(g *gin.Engine) {
	r0 := g.Group("/")
	{
		r0.GET("", myhandler.Index)
		r0.GET("health", myhandler.Health)
	}
}
