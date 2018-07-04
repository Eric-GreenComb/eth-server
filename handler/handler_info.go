package handler

import (
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/eth-server/badger"
	"github.com/Eric-GreenComb/eth-server/bean"
	"github.com/Eric-GreenComb/eth-server/cc"
	"github.com/Eric-GreenComb/eth-server/ethereum"
)

// DeployInfo DeployInfo
func DeployInfo(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_from := _formParams.From
	_pwd := _formParams.Pwd

	_value, err := badger.NewRead().Get(_from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	var _keystore string
	_keystore = strings.Replace(string(_value), "\\\"", "\"", -1)

	txOpt, err := bind.NewTransactor(strings.NewReader(_keystore), _pwd)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	_client := ethclient.NewClient(ethereum.Clients.Eth)
	defer _client.Close()

	_objAddress, _tx, _, err := cc.DeployInfoContract(txOpt, _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "address": _objAddress, "tx": _tx})
}

// SetInfoChannel SetInfoChannel
func SetInfoChannel(c *gin.Context) {
	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_from := _formParams.From
	_pwd := _formParams.Pwd

	_address := _formParams.Address
	_channel := _formParams.Value

	_jsonKey, err := badger.NewRead().Get(_from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	var _keystore string
	_keystore = strings.Replace(string(_jsonKey), "\\\"", "\"", -1)

	txOpt, err := bind.NewTransactor(strings.NewReader(_keystore), _pwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_client := ethclient.NewClient(ethereum.Clients.Eth)
	defer _client.Close()

	_tx, err := cc.NewInfoContractTransactor(common.HexToAddress(_address), _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_int64, err := strconv.ParseInt(_channel, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_valueBigInt := big.NewInt(_int64)

	tx, err := _tx.SetChannel(txOpt, _valueBigInt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": tx})
}

// GetInfoChannel GetInfoChannel
func GetInfoChannel(c *gin.Context) {
	_address := c.Params.ByName("address")

	_client := ethclient.NewClient(ethereum.Clients.Eth)
	defer _client.Close()

	tc, err := cc.NewInfoContractCaller(common.HexToAddress(_address), _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_channel, err := tc.GetChannel(&bind.CallOpts{Pending: true})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _channel})
}

// SetInfo SetInfo
func SetInfo(c *gin.Context) {
	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_from := _formParams.From
	_pwd := _formParams.Pwd

	_address := _formParams.Address
	_name := _formParams.Name
	_value := _formParams.Value

	_jsonKey, err := badger.NewRead().Get(_from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	var _keystore string
	_keystore = strings.Replace(string(_jsonKey), "\\\"", "\"", -1)

	txOpt, err := bind.NewTransactor(strings.NewReader(_keystore), _pwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_client := ethclient.NewClient(ethereum.Clients.Eth)
	defer _client.Close()

	_tx, err := cc.NewInfoContractTransactor(common.HexToAddress(_address), _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_int64, err := strconv.ParseInt(_value, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_valueBigInt := big.NewInt(_int64)

	tx, err := _tx.SetInfo(txOpt, _name, _valueBigInt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": tx})
}

// GetInfo GetInfo
func GetInfo(c *gin.Context) {
	_address := c.Params.ByName("address")

	_client := ethclient.NewClient(ethereum.Clients.Eth)
	defer _client.Close()

	tc, err := cc.NewInfoContractCaller(common.HexToAddress(_address), _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_sn, _name, _value, err := tc.GetInfo(&bind.CallOpts{Pending: true})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "sn": _sn, "name": _name, "value": _value})
}
