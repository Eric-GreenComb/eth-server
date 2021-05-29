package handler

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/eth-server/bean"
	"github.com/Eric-GreenComb/eth-server/cached"
	"github.com/Eric-GreenComb/eth-server/contracts"
	"github.com/Eric-GreenComb/eth-server/ethereum"
)

// DeployInbox DeployInbox
func DeployInbox(c *gin.Context) {

	var _params bean.FormParams
	c.BindJSON(&_params)

	_from := _params.From
	_pwd := _params.Pwd

	_value, err := cached.Manager.Get("eth.account." + _from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	// ethclient.Client
	_client := ethclient.NewClient(ethereum.Clients.Eth)
	defer _client.Close()
	fmt.Println(_client.NetworkID(context.Background()))
	fmt.Println(_client.ChainID(context.Background()))

	_nonce, err := _client.PendingNonceAt(context.Background(), common.HexToAddress(_from))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	fmt.Println(_from, "nonce:", _nonce)

	_gasPrice, err := _client.SuggestGasPrice(context.Background())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	fmt.Println("gasPrice:", _gasPrice)

	_keystore := strings.Replace(_value.(string), "\\\"", "\"", -1)
	_txOpt, err := bind.NewTransactor(strings.NewReader(_keystore), _pwd)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	_txOpt.Nonce = big.NewInt(int64(_nonce))
	_txOpt.Value = big.NewInt(0) // in wei
	_txOpt.GasLimit = uint64(300000)
	_txOpt.GasPrice = _gasPrice

	_objAddress, _tx, _, err := contracts.DeployInbox(_txOpt, _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	fmt.Println(_objAddress.Hex()) // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(_tx.Hash().Hex())  // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	// _ = _instance

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "address": _objAddress.Hex(), "tx": _tx.Hash().Hex()})
	return
}

// SetInboxValue SetInboxValue
func SetInboxValue(c *gin.Context) {

	var _params bean.FormParams
	c.BindJSON(&_params)

	_from := _params.From
	_pwd := _params.Pwd
	_smartContractAddress := _params.Address
	_int64 := _params.Int64

	_value, err := cached.Manager.Get("eth.account." + _from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	// ethclient.Client
	_client := ethclient.NewClient(ethereum.Clients.Eth)
	defer _client.Close()
	fmt.Println(_client.NetworkID(context.Background()))
	fmt.Println(_client.ChainID(context.Background()))

	_nonce, err := _client.PendingNonceAt(context.Background(), common.HexToAddress(_from))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	fmt.Println(_from, "nonce:", _nonce)

	_gasPrice, err := _client.SuggestGasPrice(context.Background())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	fmt.Println("gasPrice:", _gasPrice)

	_keystore := strings.Replace(_value.(string), "\\\"", "\"", -1)
	_txOpt, err := bind.NewTransactor(strings.NewReader(_keystore), _pwd)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	_txOpt.Nonce = big.NewInt(int64(_nonce))
	_txOpt.Value = big.NewInt(0)
	_txOpt.GasLimit = uint64(300000)
	_txOpt.GasPrice = _gasPrice
	// 1000000000

	address := common.HexToAddress(_smartContractAddress)

	fmt.Println("contract address:", address)

	_instance, err := contracts.NewInbox(address, _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_inputInt64 := big.NewInt(_int64)

	_tx, err := _instance.Set(_txOpt, _inputInt64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "tx": _tx.Hash().Hex()})
	return
}

// GetInboxValue GetInboxValue
func GetInboxValue(c *gin.Context) {

	var _params bean.FormParams
	c.BindJSON(&_params)

	_address := _params.Address

	_client := ethclient.NewClient(ethereum.Clients.Eth)
	defer _client.Close()

	_caller, err := contracts.NewInbox(common.HexToAddress(_address), _client)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_outputInt64, err := _caller.Get(nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _outputInt64})
}
