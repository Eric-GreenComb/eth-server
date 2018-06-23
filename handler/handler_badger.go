package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/eth-server/badger"
	"github.com/Eric-GreenComb/eth-server/bean"
)

// SetBadgerKey SetBadgerKey
func SetBadgerKey(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	badger.NewWrite().Set(_formParams.Key, []byte(_formParams.Value))

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "key": _formParams.Key, "value": _formParams.Value})
}

// SetBadgerKeyTTL SetBadgerKeyTTL
func SetBadgerKeyTTL(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_nTTL, _ := strconv.Atoi(_formParams.TTL)
	second := int(time.Second)

	badger.NewWrite().SetWithTTL(_formParams.Key, []byte(_formParams.Value), time.Duration(_nTTL*second))

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "key": _formParams.Key, "value": _formParams.Value, "ttl": _formParams.TTL})
}

// GetBadgerKey GetBadgerKey
func GetBadgerKey(c *gin.Context) {

	_key := c.Params.ByName("key")

	_value, err := badger.NewRead().Get(_key)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "errinfo": err.Error()})
		return
	}

	fmt.Println(strings.Replace(string(_value), "\\\"", "\"", -1))

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "key": _key, "value": string(_value)})
}
