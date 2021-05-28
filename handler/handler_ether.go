package handler

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/eth-server/bean"
	"github.com/Eric-GreenComb/eth-server/cached"
	"github.com/Eric-GreenComb/eth-server/config"
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

// SendEthCoin SendEthCoin
func SendEthCoin(c *gin.Context) {
	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_from := _formParams.From
	_to := _formParams.To
	_amount := _formParams.Amount
	_pwd := _formParams.Pwd
	_decimals := _formParams.Decimals
	_desc := _formParams.Desc

	_int, err := strconv.Atoi(_decimals)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
	}

	_value, err := cached.Manager.Get("eth.account." + _from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_key, err := keystore.DecryptKey(_value.([]byte), _pwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_amountBigInt := ethereum.StringToWei(_amount, _int)
	_chainIDBigInt := big.NewInt(config.Ethereum.ChainID)

	_nonce, err := ethereum.PendingNonce(_from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_inputData := []byte(_desc)
	_txid, err := ethereum.SendEthCoins(_to, _nonce, _amountBigInt, _key.PrivateKey, _chainIDBigInt, _inputData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _txid})
}

// GetChainID GetChainID
func GetChainID(c *gin.Context) {

	_client := ethclient.NewClient(ethereum.Clients.Eth)
	_ctx := context.Background()
	_chainid, err := _client.NetworkID(_ctx)
	if err != nil {
		fmt.Println("NetworkID error", err.Error())
	} else {
		fmt.Println(_chainid)
	}

	defer _client.Close()

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "ChainID": _chainid})
}

// GetBalance GetBalance
func GetBalance(c *gin.Context) {

	_addr := c.Params.ByName("addr")

	_ethCoin, err := ethereum.GetBalance(_addr)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	n := new(big.Int)
	m := new(big.Int)
	k := big.NewInt(1000000000000)
	n, ok := n.SetString(Remove0x(_ethCoin), 16)
	var _balance int64
	if !ok {
		fmt.Println("SetString: error")
		_balance = 0
	} else {
		_balance = m.Div(n, k).Int64()
	}
	fmt.Println("balance", _balance)

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "hex": _ethCoin, "balance": n})
}

// Remove0x Remove0x
func Remove0x(s string) string {
	if len(s) < 2 {
		return s
	}
	if s[0:2] == "0x" || s[0:2] == "0X" {
		return s[2:]
	}
	return s
}

// StringToWei StringToWei
func StringToWei(c *gin.Context) {

	_val := c.Params.ByName("val")
	_decimals := c.Params.ByName("decimals")

	_int, err := strconv.Atoi(_decimals)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
	}

	_wei := ethereum.StringToWei(_val, _int)

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _wei.String()})
}

// FloatToWei FloatToWei
func FloatToWei(c *gin.Context) {

	_val := c.Params.ByName("val")
	_decimals := c.Params.ByName("decimals")

	_float64, err := strconv.ParseFloat(_val, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
	}

	_int, err := strconv.Atoi(_decimals)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 2, "msg": err.Error()})
	}

	_wei := ethereum.FloatToWei(_float64, _int)

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _wei})
}
