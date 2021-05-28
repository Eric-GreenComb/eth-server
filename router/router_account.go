package router

import (
	"github.com/gin-gonic/gin"

	myhandler "github.com/Eric-GreenComb/eth-server/handler"
)

// SetupAccountRouter SetupAccountRouter
func SetupAccountRouter(g *gin.Engine) {
	r := g.Group("/account")
	{
		r.POST("/create/:passphrase", myhandler.CreateAccount)
		r.POST("/bip39/create", myhandler.CreateBIP39)
		r.POST("/bip39/keystore/create", myhandler.CreateBIP39Keysore)

		r.GET("/bip39/keystore/:address", myhandler.GetBIP39Keysore)
		r.POST("/checkpwd", myhandler.CheckPassphrase)
	}
}
