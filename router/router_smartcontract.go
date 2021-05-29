package router

import (
	"github.com/gin-gonic/gin"

	myhandler "github.com/Eric-GreenComb/eth-server/handler"
)

// SetupSmartContractRouter SetupSmartContractRouter
func SetupSmartContractRouter(g *gin.Engine) {
	_rinbox := g.Group("/smartcontract/inbox")
	{
		_rinbox.POST("/deploy", myhandler.DeployInbox)
		_rinbox.POST("/set", myhandler.SetInboxValue)
		_rinbox.POST("/set/raw", myhandler.SetRawInboxValue)
		_rinbox.POST("/get", myhandler.GetInboxValue)
	}
}
