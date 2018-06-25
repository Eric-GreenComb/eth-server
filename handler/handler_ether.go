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
	"github.com/Eric-GreenComb/eth-server/ethereum"
)

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

// TransferErc20 TransferErc20
func TransferErc20(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_conaddr := _formParams.Address
	_from := _formParams.From
	_to := _formParams.To
	_amount := _formParams.Amount
	_pwd := _formParams.Pwd

	_int64, err := strconv.ParseInt(_amount, 10, 64)
	if err != nil {
		c.String(200, err.Error())
		return
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

	_bigint := big.NewInt(_int64)

	_tx, err := ts.Transfer(txOpt, common.HexToAddress(_to), _bigint)
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

	_int64, err := strconv.ParseInt(_amount, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "errinfo": err.Error()})
		return
	}

	_value, err := badger.NewRead().Get(_from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "errinfo": err.Error()})
		return
	}

	var _keystore string
	_keystore = strings.Replace(string(_value), "\\\"", "\"", -1)
	_key, err := keystore.DecryptKey([]byte(_keystore), _pwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "errinfo": err.Error()})
		return
	}

	_amountBigInt := big.NewInt(_int64)
	// _chainIDBigInt := big.NewInt(config.EthereumConfig.ChainID)

	_nonce, err := ethereum.PendingNonce(_from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "errinfo": err.Error()})
		return
	}

	_txid, err := ethereum.SendEthTokens(_conaddr, _to, _nonce, _amountBigInt, _key.PrivateKey, nil)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "errinfo": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "txid": _txid})
}
