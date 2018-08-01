package handler

import (
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/eth-server/badger"
	"github.com/Eric-GreenComb/eth-server/bean"
	"github.com/Eric-GreenComb/eth-server/cc"
	"github.com/Eric-GreenComb/eth-server/ethereum"
)

// DeployTeam DeployTeam
func DeployTeam(c *gin.Context) {
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

	_objAddress, _tx, _, err := cc.DeployTeamJust(txOpt, _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "address": _objAddress, "tx": _tx})
}

// DeployPlayerBook DeployPlayerBook
func DeployPlayerBook(c *gin.Context) {
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

	_objAddress, _tx, _, err := cc.DeployPlayerBook(txOpt, _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "address": _objAddress, "tx": _tx})
}

// DeployF3d DeployF3d
func DeployF3d(c *gin.Context) {
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

	_objAddress, _tx, _, err := cc.DeployFoMo3Dlong(txOpt, _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "address": _objAddress, "tx": _tx})
}
