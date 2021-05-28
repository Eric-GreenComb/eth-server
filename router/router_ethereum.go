package router

import (
	"github.com/gin-gonic/gin"

	myhandler "github.com/Eric-GreenComb/eth-server/handler"
)

// SetupEthereumRouter SetupEthereumRouter
func SetupEthereumRouter(g *gin.Engine) {
	r := g.Group("/ethereum")
	{
		r.GET("/chainid", myhandler.GetChainID)
		r.GET("/nonce", myhandler.PendingNonce)
		r.POST("/send", myhandler.SendEthCoin)
		r.GET("/balance/:addr", myhandler.GetBalance)

		r.GET("/string2wei/:val/:decimals", myhandler.StringToWei)
	}
}
