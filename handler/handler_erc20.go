package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/eth-server/bean"
	"github.com/Eric-GreenComb/eth-server/ethereum"
)

// GetNonceAt GetNonceAt
func GetNonceAt(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	fmt.Println(_formParams.Params)

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": ethereum.GetNonceAt(_formParams.Params)})

}
