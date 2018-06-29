package handler

import (
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

// DeployBank DeployBank
func DeployBank(c *gin.Context) {

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

	_bankAddress, _tx, _, err := cc.DeployBank(txOpt, _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "address": _bankAddress, "tx": _tx})
}

// DepositBank DepositBank
func DepositBank(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_conaddr := _formParams.Address

	_from := _formParams.From
	_pwd := _formParams.Pwd

	_amount := _formParams.Amount
	_decimals := _formParams.Decimals

	_int, err := strconv.Atoi(_decimals)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 2, "msg": err.Error()})
	}

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

	_amountBigInt := ethereum.StringToWei(_amount, _int)

	_transactor, err := cc.NewBankTransactor(common.HexToAddress(_conaddr), _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	txOpt.Value = _amountBigInt

	_tx, err := _transactor.Deposit(txOpt, _amountBigInt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "tx": _tx})
}

// WithdrawBank WithdrawBank
func WithdrawBank(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_conaddr := _formParams.Address

	_from := _formParams.From
	_pwd := _formParams.Pwd

	_amount := _formParams.Amount
	_decimals := _formParams.Decimals

	_int, err := strconv.Atoi(_decimals)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 2, "msg": err.Error()})
	}

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

	_amountBigInt := ethereum.StringToWei(_amount, _int)

	_transactor, err := cc.NewBankTransactor(common.HexToAddress(_conaddr), _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_tx, err := _transactor.Withdraw(txOpt, _amountBigInt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "tx": _tx})
}

// GetBankBalance GetBankBalance
func GetBankBalance(c *gin.Context) {

	_conaddr := c.Params.ByName("conaddr")
	_address := c.Params.ByName("address")

	_client := ethclient.NewClient(ethereum.Clients.Eth)
	defer _client.Close()

	tc, err := cc.NewBankCaller(common.HexToAddress(_conaddr), _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_value, err := tc.BalanceOf(&bind.CallOpts{Pending: true}, common.HexToAddress(_address))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _value})
}
