package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/eth-server/bean"
	"github.com/Eric-GreenComb/eth-server/ethereum"
)

// PendingNonce PendingNonce
func PendingNonce(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_nonce, err := ethereum.PendingNonce(_formParams.Params)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _nonce})
}
