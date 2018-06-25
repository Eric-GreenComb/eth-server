package handler

import (
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/eth-server/badger"
	"github.com/Eric-GreenComb/eth-server/bean"
	"github.com/Eric-GreenComb/eth-server/cc"
	"github.com/Eric-GreenComb/eth-server/config"
	"github.com/Eric-GreenComb/eth-server/ethereum"
)

// DeployErc20 DeployErc20
func DeployErc20(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_name := _formParams.Name
	_symbol := _formParams.Symbol
	_total := _formParams.Total
	_decimals := _formParams.Decimals

	_from := _formParams.From
	_pwd := _formParams.Pwd

	_int, err := strconv.Atoi(_decimals)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 2, "msg": err.Error()})
	}

	_value, err := badger.NewRead().Get(_from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "errinfo": err.Error()})
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

	_totalBigInt := ethereum.StringToWei(_total, _int)

	_tokenAddress, _tx, _, err := cc.DeployHumanStandardToken(txOpt, _client, _totalBigInt, _name, uint8(_int), _symbol)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "address": _tokenAddress, "tx": _tx})
}

// TransferErc20 TransferErc20
func TransferErc20(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_conaddr := _formParams.Address
	_from := _formParams.From
	_to := _formParams.To
	_amount := _formParams.Amount
	_decimals := _formParams.Decimals
	_pwd := _formParams.Pwd

	_int, err := strconv.Atoi(_decimals)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 2, "msg": err.Error()})
	}

	_value, err := badger.NewRead().Get(_from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "errinfo": err.Error()})
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
	ts, _ := cc.NewHumanStandardTokenTransactor(common.HexToAddress(_conaddr), _client)

	_amountBigInt := ethereum.StringToWei(_amount, _int)
	_tx, err := ts.Transfer(txOpt, common.HexToAddress(_to), _amountBigInt)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "errinfo": _tx})
}

// RawTransferErc20 RawTransferErc20
func RawTransferErc20(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_conaddr := _formParams.Address
	_from := _formParams.From
	_to := _formParams.To
	_amount := _formParams.Amount
	_pwd := _formParams.Pwd
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
	_key, err := keystore.DecryptKey([]byte(_keystore), _pwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_amountBigInt := ethereum.StringToWei(_amount, _int)
	_chainIDBigInt := big.NewInt(config.EthereumConfig.ChainID)

	_nonce, err := ethereum.PendingNonce(_from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_txid, err := ethereum.SendEthTokens(_conaddr, _to, _nonce, _amountBigInt, _key.PrivateKey, _chainIDBigInt)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _txid})
}

// GetBalance GetBalance
func GetBalance(c *gin.Context) {

	_conaddr := c.Params.ByName("conaddr")
	_addr := c.Params.ByName("addr")

	_client := ethclient.NewClient(ethereum.Clients.Eth)
	defer _client.Close()

	_caller, err := cc.NewHumanStandardTokenCaller(common.HexToAddress(_conaddr), _client)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	_bigint, err := _caller.BalanceOf(&bind.CallOpts{Pending: true}, common.HexToAddress(_addr))
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, _bigint.String())
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
