package ethereum

import (
	"crypto/ecdsa"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// SendRawTransaction SendRawTransaction
func SendRawTransaction(address string, inputData []byte, nonce uint64, priv *ecdsa.PrivateKey, chainID, gasPrice *big.Int) (string, error) {

	client := Clients.Eth

	var err error

	tx := types.NewTransaction(nonce, ethcommon.HexToAddress(address), nil, gasDefaultLimit, gasPrice, []byte(inputData))

	var signed *types.Transaction

	if chainID != nil {
		signed, _ = types.SignTx(tx, types.NewEIP155Signer(chainID), priv)
	} else {
		signed, _ = types.SignTx(tx, types.HomesteadSigner{}, priv)
	}

	txData, err := rlp.EncodeToBytes(signed)
	if err != nil {
		return "", err
	}

	txID := &ethcommon.Hash{}
	err = client.Call(&txID, "eth_sendRawTransaction", hexutil.Encode(txData))
	if err != nil {
		return "", err
	}

	txid := txID.String()

	return txid, nil
}
