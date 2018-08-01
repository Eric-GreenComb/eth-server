// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cc

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MSFunABI is the input ABI used to generate the binding from.
const MSFunABI = "[]"

// MSFunBin is the compiled bytecode used for deploying new contracts.
const MSFunBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058208fbca7a6f9504d7838903c0e56e71b90eb842daba112fb83f09cf53f241fea790029`

// DeployMSFun deploys a new Ethereum contract, binding an instance of MSFun to it.
func DeployMSFun(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MSFun, error) {
	parsed, err := abi.JSON(strings.NewReader(MSFunABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MSFunBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MSFun{MSFunCaller: MSFunCaller{contract: contract}, MSFunTransactor: MSFunTransactor{contract: contract}, MSFunFilterer: MSFunFilterer{contract: contract}}, nil
}

// MSFun is an auto generated Go binding around an Ethereum contract.
type MSFun struct {
	MSFunCaller     // Read-only binding to the contract
	MSFunTransactor // Write-only binding to the contract
	MSFunFilterer   // Log filterer for contract events
}

// MSFunCaller is an auto generated read-only Go binding around an Ethereum contract.
type MSFunCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MSFunTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MSFunTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MSFunFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MSFunFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MSFunSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MSFunSession struct {
	Contract     *MSFun            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MSFunCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MSFunCallerSession struct {
	Contract *MSFunCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MSFunTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MSFunTransactorSession struct {
	Contract     *MSFunTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MSFunRaw is an auto generated low-level Go binding around an Ethereum contract.
type MSFunRaw struct {
	Contract *MSFun // Generic contract binding to access the raw methods on
}

// MSFunCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MSFunCallerRaw struct {
	Contract *MSFunCaller // Generic read-only contract binding to access the raw methods on
}

// MSFunTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MSFunTransactorRaw struct {
	Contract *MSFunTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMSFun creates a new instance of MSFun, bound to a specific deployed contract.
func NewMSFun(address common.Address, backend bind.ContractBackend) (*MSFun, error) {
	contract, err := bindMSFun(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MSFun{MSFunCaller: MSFunCaller{contract: contract}, MSFunTransactor: MSFunTransactor{contract: contract}, MSFunFilterer: MSFunFilterer{contract: contract}}, nil
}

// NewMSFunCaller creates a new read-only instance of MSFun, bound to a specific deployed contract.
func NewMSFunCaller(address common.Address, caller bind.ContractCaller) (*MSFunCaller, error) {
	contract, err := bindMSFun(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MSFunCaller{contract: contract}, nil
}

// NewMSFunTransactor creates a new write-only instance of MSFun, bound to a specific deployed contract.
func NewMSFunTransactor(address common.Address, transactor bind.ContractTransactor) (*MSFunTransactor, error) {
	contract, err := bindMSFun(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MSFunTransactor{contract: contract}, nil
}

// NewMSFunFilterer creates a new log filterer instance of MSFun, bound to a specific deployed contract.
func NewMSFunFilterer(address common.Address, filterer bind.ContractFilterer) (*MSFunFilterer, error) {
	contract, err := bindMSFun(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MSFunFilterer{contract: contract}, nil
}

// bindMSFun binds a generic wrapper to an already deployed contract.
func bindMSFun(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MSFunABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MSFun *MSFunRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MSFun.Contract.MSFunCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MSFun *MSFunRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MSFun.Contract.MSFunTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MSFun *MSFunRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MSFun.Contract.MSFunTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MSFun *MSFunCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MSFun.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MSFun *MSFunTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MSFun.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MSFun *MSFunTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MSFun.Contract.contract.Transact(opts, method, params...)
}

// TeamJustABI is the input ABI used to generate the binding from.
const TeamJustABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"isDev\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_whatFunction\",\"type\":\"bytes32\"}],\"name\":\"deleteAnyProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"adminCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_howMany\",\"type\":\"uint256\"}],\"name\":\"changeRequiredDevSignatures\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_isDev\",\"type\":\"bool\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_howMany\",\"type\":\"uint256\"}],\"name\":\"changeRequiredSignatures\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requiredSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_whatFunction\",\"type\":\"bytes32\"}],\"name\":\"checkData\",\"outputs\":[{\"name\":\"message_data\",\"type\":\"bytes32\"},{\"name\":\"signature_count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"adminName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"devCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_whatFunction\",\"type\":\"bytes32\"},{\"name\":\"_signerA\",\"type\":\"uint256\"},{\"name\":\"_signerB\",\"type\":\"uint256\"},{\"name\":\"_signerC\",\"type\":\"uint256\"}],\"name\":\"checkSignersByName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requiredDevSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// TeamJustBin is the compiled bytecode used for deploying new contracts.
const TeamJustBin = `0x608060405260008054600160a060020a031916905534801561002057600080fd5b506040805160608181018352600180835260208084018281527f696e76656e746f720000000000000000000000000000000000000000000000008587019081527328c0f6142d1232c9e663e29cc0a6f8f0872693736000908152600280855296517fbd63e4667d9cae038f97c233744e4e0107e4dc99d6d45d181f27607b529ccf4a80549451151561010090810261ff001993151560ff199788161784161790915592517fbd63e4667d9cae038f97c233744e4e0107e4dc99d6d45d181f27607b529ccf4b5588518088018a528681528086018781527f6d616e74736f0000000000000000000000000000000000000000000000000000828c0190815273e54c005c9ef185cfe70209ad825301f9a84534a885528a885291517f70bc3b14fcde8e9d318942ddc202321789750d66624073988dd617c1bb44d5e78054925115158702911515928816929092178416179055517f70bc3b14fcde8e9d318942ddc202321789750d66624073988dd617c1bb44d5e85588518088018a528681528086018781527f6a7573746f000000000000000000000000000000000000000000000000000000828c01908152733d3b33b8f50ab9e8f5a9ff369853f0e638450adb85528a885291517f048e034a4b4d6a0035c4393a6ac9c2748159a972a6b908c8ec055d67c32dea9b8054925115158702911515928816929092178416179055517f048e034a4b4d6a0035c4393a6ac9c2748159a972a6b908c8ec055d67c32dea9c5588518088018a528681528086018781527f73756d70756e6b00000000000000000000000000000000000000000000000000828c0190815273dbeb69c655b666b3e17b8061df7ea4cc2399df1185528a885291517fce0b0aff3f03938bcd39f6f0181a6cb0740f7a7e659dba2dd830e9e7b916b74f8054925115158702911515928816929092178416179055517fce0b0aff3f03938bcd39f6f0181a6cb0740f7a7e659dba2dd830e9e7b916b75055885196870189528587528685018681527f6465706c6f796572000000000000000000000000000000000000000000000000998801998a52736b9e7c45622832a12f728ca87e23fa3a6b512fe29092529690935293517f248c391ec2cbcb20087b53b7e4e8dd9e6e1880159f1b0f97f8070ccfdecfc067805493511515909502901515929091169190911790931692909217905590517f248c391ec2cbcb20087b53b7e4e8dd9e6e1880159f1b0f97f8070ccfdecfc06855600560038190556004819055819055600655611525806103c86000396000f3006080604052600436106100c15763ffffffff60e060020a6000350416630c3f64bf81146101515780630efcf295146101865780631785f53c146101a057806324d7806c146101c15780632b7832b3146101e25780632c29665614610209578063372cd1831461022157806339f636ab1461024a57806366d38203146102625780638d06804314610283578063a553506e14610298578063af1c084d146102c9578063cebc141a146102ea578063ed3643d6146102ff578063fcf2f85f1461033e575b600054604080517fd0e30db00000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169163d0e30db091303191600480830192602092919082900301818588803b15801561012257600080fd5b505af1158015610136573d6000803e3d6000fd5b50505050506040513d602081101561014d57600080fd5b5050005b34801561015d57600080fd5b50610172600160a060020a0360043516610353565b604080519115158252519081900360200190f35b34801561019257600080fd5b5061019e600435610376565b005b3480156101ac57600080fd5b5061019e600160a060020a03600435166103f8565b3480156101cd57600080fd5b50610172600160a060020a03600435166107d8565b3480156101ee57600080fd5b506101f76107f6565b60408051918252519081900360200190f35b34801561021557600080fd5b5061019e6004356107fc565b34801561022d57600080fd5b5061019e600160a060020a0360043516602435604435151561098c565b34801561025657600080fd5b5061019e600435610b27565b34801561026e57600080fd5b5061019e600160a060020a0360043516610cb7565b34801561028f57600080fd5b506101f7610d70565b3480156102a457600080fd5b506102b0600435610d76565b6040805192835260208301919091528051918290030190f35b3480156102d557600080fd5b506101f7600160a060020a0360043516610e29565b3480156102f657600080fd5b506101f7610e47565b34801561030b57600080fd5b50610320600435602435604435606435610e4d565b60408051938452602084019290925282820152519081900360600190f35b34801561034a57600080fd5b506101f761108e565b600160a060020a0316600090815260026020526040902054610100900460ff1690565b3360009081526002602052604090205460ff6101009091041615156001146103ea576040805160e560020a62461bcd02815260206004820152602960248201526000805160206114ba83398151915260448201526000805160206114da833981519152606482015290519081900360840190fd5b6103f5600182611094565b50565b3360009081526002602052604090205460ff61010090910416151560011461046c576040805160e560020a62461bcd02815260206004820152602960248201526000805160206114ba83398151915260448201526000805160206114da833981519152606482015290519081900360840190fd5b6003546001106104da576040805160e560020a62461bcd028152602060048201526033602482015260008051602061149a83398151915260448201527f206c657373207468616e20322061646d696e7300000000000000000000000000606482015290519081900360840190fd5b6005546003541015610570576040805160e560020a62461bcd02815260206004820152604f602482015260008051602061149a83398151915260448201527f206c6573732061646d696e73207468616e206e756d626572206f66207265717560648201527f69726564207369676e6174757265730000000000000000000000000000000000608482015290519081900360a40190fd5b600160a060020a03811660009081526002602052604090205460ff610100909104161515600114156106a05760045460011061060a576040805160e560020a62461bcd028152602060048201526031602482015260008051602061149a83398151915260448201527f206c657373207468616e20322064657673000000000000000000000000000000606482015290519081900360840190fd5b60065460045410156106a0576040805160e560020a62461bcd028152602060048201526051602482015260008051602061149a83398151915260448201527f206c6573732064657673207468616e206e756d626572206f662072657175697260648201527f656420646576207369676e617475726573000000000000000000000000000000608482015290519081900360a40190fd5b6106ce60016006547f72656d6f766541646d696e000000000000000000000000000000000000000000611146565b1515600114156103f55761070360017f72656d6f766541646d696e000000000000000000000000000000000000000000611094565b600160a060020a03811660009081526002602052604090205460ff1615156001141561076957600160a060020a0381166000908152600260205260409020805460ff19169055600380546000190190556005546001101561076957600580546000190190555b600160a060020a03811660009081526002602052604090205460ff610100909104161515600114156103f557600160a060020a0381166000908152600260205260409020805461ff001916905560048054600019019055600654600110156103f5576006805460001901905550565b600160a060020a031660009081526002602052604090205460ff1690565b60035490565b3360009081526002602052604090205460ff610100909104161515600114610870576040805160e560020a62461bcd02815260206004820152602960248201526000805160206114ba83398151915260448201526000805160206114da833981519152606482015290519081900360840190fd5b60008111801561088257506004548111155b1515610924576040805160e560020a62461bcd02815260206004820152604960248201527f6368616e676552657175697265644465765369676e617475726573206661696c60448201527f6564202d206d757374206265206265747765656e203120616e64206e756d626560648201527f72206f6620646576730000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b61095260016006547f6368616e676552657175697265644465765369676e6174757265730000000000611146565b1515600114156103f55761098760017f6368616e676552657175697265644465765369676e6174757265730000000000611094565b600655565b3360009081526002602052604090205460ff610100909104161515600114610a00576040805160e560020a62461bcd02815260206004820152602960248201526000805160206114ba83398151915260448201526000805160206114da833981519152606482015290519081900360840190fd5b610a2e60016006547f61646441646d696e000000000000000000000000000000000000000000000000611146565b151560011415610b0757610a6360017f61646441646d696e000000000000000000000000000000000000000000000000611094565b600160a060020a03831660009081526002602052604090205460ff161515610abd57600160a060020a0383166000908152600260205260409020805460ff1916600190811790915560038054820190556005805490910190555b60018115151415610b0757600160a060020a0383166000908152600260205260409020805461ff001916610100831515021790556004805460019081019091556006805490910190555b50600160a060020a03909116600090815260026020526040902060010155565b3360009081526002602052604090205460ff610100909104161515600114610b9b576040805160e560020a62461bcd02815260206004820152602960248201526000805160206114ba83398151915260448201526000805160206114da833981519152606482015290519081900360840190fd5b600081118015610bad57506003548111155b1515610c4f576040805160e560020a62461bcd02815260206004820152604860248201527f6368616e676552657175697265645369676e617475726573206661696c65642060448201527f2d206d757374206265206265747765656e203120616e64206e756d626572206f60648201527f662061646d696e73000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b610c7d60016006547f6368616e676552657175697265645369676e6174757265730000000000000000611146565b1515600114156103f557610cb260017f6368616e676552657175697265645369676e6174757265730000000000000000611094565b600555565b3360009081526002602052604090205460ff610100909104161515600114610d2b576040805160e560020a62461bcd02815260206004820152602960248201526000805160206114ba83398151915260448201526000805160206114da833981519152606482015290519081900360840190fd5b600054600160a060020a031615610d4157600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60055490565b33600090815260026020526040812054819060ff161515600114610e0a576040805160e560020a62461bcd02815260206004820152602e60248201527f6f6e6c7941646d696e73206661696c6564202d206d73672e73656e646572206960448201527f73206e6f7420616e2061646d696e000000000000000000000000000000000000606482015290519081900360840190fd5b610e156001846112fd565b610e20600185611321565b91509150915091565b600160a060020a031660009081526002602052604090206001015490565b60045490565b336000908152600260205260408120548190819060ff161515600114610ee3576040805160e560020a62461bcd02815260206004820152602e60248201527f6f6e6c7941646d696e73206661696c6564202d206d73672e73656e646572206960448201527f73206e6f7420616e2061646d696e000000000000000000000000000000000000606482015290519081900360840190fd5b3063af1c084d610ef560018a8a611348565b6040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b158015610f4057600080fd5b505af1158015610f54573d6000803e3d6000fd5b505050506040513d6020811015610f6a57600080fd5b50513063af1c084d610f7e60018b8a611348565b6040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b158015610fc957600080fd5b505af1158015610fdd573d6000803e3d6000fd5b505050506040513d6020811015610ff357600080fd5b50513063af1c084d61100760018c8a611348565b6040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b15801561105257600080fd5b505af1158015611066573d6000803e3d6000fd5b505050506040513d602081101561107c57600080fd5b50519199909850909650945050505050565b60065490565b60008060006110a284611405565b9250600090505b60008381526020869052604090206001015481101561112b57600083815260208681526040808320848452600381018084528285208054600160a060020a031680875260029093018552928520805460ff191690559385905292909152805473ffffffffffffffffffffffffffffffffffffffff1916905591506001016110a9565b50506000908152602092909252506040812081815560010155565b600080600080600061115786611405565b600081815260208a905260408082206001015490519296509450339350903690808383808284376040519201829003909120945050508415159150611221905057600084815260208981526040808320848155600160a060020a038616808552600282018452828520805460ff19166001908117909155888652600383018552928520805473ffffffffffffffffffffffffffffffffffffffff1916909117905592879052908a905290810180549091019081905587141561121c57600194506112f2565b6112f2565b6000848152602089905260409020548114156112f257600084815260208981526040808320600160a060020a038616845260020190915290205460ff1615156112d457600084815260208981526040808320600160a060020a038616808552600282018452828520805460ff19166001908117909155888652600383018552928520805473ffffffffffffffffffffffffffffffffffffffff1916909117905592879052908a9052908101805490910190555b6000848152602089905260409020600101548714156112f257600194505b505050509392505050565b60008061130983611405565b60009081526020949094525050604090912054919050565b60008061132d83611405565b60009081526020949094525050604090912060010154919050565b6000808083116113c8576040805160e560020a62461bcd02815260206004820152602860248201527f4d5346756e20636865636b5369676e6572206661696c6564202d2030206e6f7460448201527f20616c6c6f776564000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6113d184611405565b60008181526020878152604080832060001988018452600301909152902054600160a060020a031692509050509392505050565b6040805160208082018490526c01000000000000000000000000300282840152825160348184030181526054909201928390528151600093918291908401908083835b602083106114675780518252601f199092019160209182019101611448565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912095945050505050560072656d6f766541646d696e206661696c6564202d2063616e6e6f7420686176656f6e6c7944657673206661696c6564202d206d73672e73656e646572206973206e6f742061206465760000000000000000000000000000000000000000000000a165627a7a72305820939a77bc200c43ff1bc532bdca43c6a9dd0c69de4937a1df5192243a8f211d340029`

// DeployTeamJust deploys a new Ethereum contract, binding an instance of TeamJust to it.
func DeployTeamJust(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TeamJust, error) {
	parsed, err := abi.JSON(strings.NewReader(TeamJustABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TeamJustBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TeamJust{TeamJustCaller: TeamJustCaller{contract: contract}, TeamJustTransactor: TeamJustTransactor{contract: contract}, TeamJustFilterer: TeamJustFilterer{contract: contract}}, nil
}

// TeamJust is an auto generated Go binding around an Ethereum contract.
type TeamJust struct {
	TeamJustCaller     // Read-only binding to the contract
	TeamJustTransactor // Write-only binding to the contract
	TeamJustFilterer   // Log filterer for contract events
}

// TeamJustCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeamJustCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamJustTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeamJustTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamJustFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeamJustFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamJustSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeamJustSession struct {
	Contract     *TeamJust         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeamJustCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeamJustCallerSession struct {
	Contract *TeamJustCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TeamJustTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeamJustTransactorSession struct {
	Contract     *TeamJustTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TeamJustRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeamJustRaw struct {
	Contract *TeamJust // Generic contract binding to access the raw methods on
}

// TeamJustCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeamJustCallerRaw struct {
	Contract *TeamJustCaller // Generic read-only contract binding to access the raw methods on
}

// TeamJustTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeamJustTransactorRaw struct {
	Contract *TeamJustTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeamJust creates a new instance of TeamJust, bound to a specific deployed contract.
func NewTeamJust(address common.Address, backend bind.ContractBackend) (*TeamJust, error) {
	contract, err := bindTeamJust(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeamJust{TeamJustCaller: TeamJustCaller{contract: contract}, TeamJustTransactor: TeamJustTransactor{contract: contract}, TeamJustFilterer: TeamJustFilterer{contract: contract}}, nil
}

// NewTeamJustCaller creates a new read-only instance of TeamJust, bound to a specific deployed contract.
func NewTeamJustCaller(address common.Address, caller bind.ContractCaller) (*TeamJustCaller, error) {
	contract, err := bindTeamJust(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeamJustCaller{contract: contract}, nil
}

// NewTeamJustTransactor creates a new write-only instance of TeamJust, bound to a specific deployed contract.
func NewTeamJustTransactor(address common.Address, transactor bind.ContractTransactor) (*TeamJustTransactor, error) {
	contract, err := bindTeamJust(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeamJustTransactor{contract: contract}, nil
}

// NewTeamJustFilterer creates a new log filterer instance of TeamJust, bound to a specific deployed contract.
func NewTeamJustFilterer(address common.Address, filterer bind.ContractFilterer) (*TeamJustFilterer, error) {
	contract, err := bindTeamJust(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeamJustFilterer{contract: contract}, nil
}

// bindTeamJust binds a generic wrapper to an already deployed contract.
func bindTeamJust(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TeamJustABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeamJust *TeamJustRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TeamJust.Contract.TeamJustCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeamJust *TeamJustRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeamJust.Contract.TeamJustTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeamJust *TeamJustRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeamJust.Contract.TeamJustTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeamJust *TeamJustCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TeamJust.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeamJust *TeamJustTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeamJust.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeamJust *TeamJustTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeamJust.Contract.contract.Transact(opts, method, params...)
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() constant returns(uint256)
func (_TeamJust *TeamJustCaller) AdminCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TeamJust.contract.Call(opts, out, "adminCount")
	return *ret0, err
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() constant returns(uint256)
func (_TeamJust *TeamJustSession) AdminCount() (*big.Int, error) {
	return _TeamJust.Contract.AdminCount(&_TeamJust.CallOpts)
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() constant returns(uint256)
func (_TeamJust *TeamJustCallerSession) AdminCount() (*big.Int, error) {
	return _TeamJust.Contract.AdminCount(&_TeamJust.CallOpts)
}

// AdminName is a free data retrieval call binding the contract method 0xaf1c084d.
//
// Solidity: function adminName(_who address) constant returns(bytes32)
func (_TeamJust *TeamJustCaller) AdminName(opts *bind.CallOpts, _who common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TeamJust.contract.Call(opts, out, "adminName", _who)
	return *ret0, err
}

// AdminName is a free data retrieval call binding the contract method 0xaf1c084d.
//
// Solidity: function adminName(_who address) constant returns(bytes32)
func (_TeamJust *TeamJustSession) AdminName(_who common.Address) ([32]byte, error) {
	return _TeamJust.Contract.AdminName(&_TeamJust.CallOpts, _who)
}

// AdminName is a free data retrieval call binding the contract method 0xaf1c084d.
//
// Solidity: function adminName(_who address) constant returns(bytes32)
func (_TeamJust *TeamJustCallerSession) AdminName(_who common.Address) ([32]byte, error) {
	return _TeamJust.Contract.AdminName(&_TeamJust.CallOpts, _who)
}

// CheckData is a free data retrieval call binding the contract method 0xa553506e.
//
// Solidity: function checkData(_whatFunction bytes32) constant returns(message_data bytes32, signature_count uint256)
func (_TeamJust *TeamJustCaller) CheckData(opts *bind.CallOpts, _whatFunction [32]byte) (struct {
	MessageData    [32]byte
	SignatureCount *big.Int
}, error) {
	ret := new(struct {
		MessageData    [32]byte
		SignatureCount *big.Int
	})
	out := ret
	err := _TeamJust.contract.Call(opts, out, "checkData", _whatFunction)
	return *ret, err
}

// CheckData is a free data retrieval call binding the contract method 0xa553506e.
//
// Solidity: function checkData(_whatFunction bytes32) constant returns(message_data bytes32, signature_count uint256)
func (_TeamJust *TeamJustSession) CheckData(_whatFunction [32]byte) (struct {
	MessageData    [32]byte
	SignatureCount *big.Int
}, error) {
	return _TeamJust.Contract.CheckData(&_TeamJust.CallOpts, _whatFunction)
}

// CheckData is a free data retrieval call binding the contract method 0xa553506e.
//
// Solidity: function checkData(_whatFunction bytes32) constant returns(message_data bytes32, signature_count uint256)
func (_TeamJust *TeamJustCallerSession) CheckData(_whatFunction [32]byte) (struct {
	MessageData    [32]byte
	SignatureCount *big.Int
}, error) {
	return _TeamJust.Contract.CheckData(&_TeamJust.CallOpts, _whatFunction)
}

// CheckSignersByName is a free data retrieval call binding the contract method 0xed3643d6.
//
// Solidity: function checkSignersByName(_whatFunction bytes32, _signerA uint256, _signerB uint256, _signerC uint256) constant returns(bytes32, bytes32, bytes32)
func (_TeamJust *TeamJustCaller) CheckSignersByName(opts *bind.CallOpts, _whatFunction [32]byte, _signerA *big.Int, _signerB *big.Int, _signerC *big.Int) ([32]byte, [32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _TeamJust.contract.Call(opts, out, "checkSignersByName", _whatFunction, _signerA, _signerB, _signerC)
	return *ret0, *ret1, *ret2, err
}

// CheckSignersByName is a free data retrieval call binding the contract method 0xed3643d6.
//
// Solidity: function checkSignersByName(_whatFunction bytes32, _signerA uint256, _signerB uint256, _signerC uint256) constant returns(bytes32, bytes32, bytes32)
func (_TeamJust *TeamJustSession) CheckSignersByName(_whatFunction [32]byte, _signerA *big.Int, _signerB *big.Int, _signerC *big.Int) ([32]byte, [32]byte, [32]byte, error) {
	return _TeamJust.Contract.CheckSignersByName(&_TeamJust.CallOpts, _whatFunction, _signerA, _signerB, _signerC)
}

// CheckSignersByName is a free data retrieval call binding the contract method 0xed3643d6.
//
// Solidity: function checkSignersByName(_whatFunction bytes32, _signerA uint256, _signerB uint256, _signerC uint256) constant returns(bytes32, bytes32, bytes32)
func (_TeamJust *TeamJustCallerSession) CheckSignersByName(_whatFunction [32]byte, _signerA *big.Int, _signerB *big.Int, _signerC *big.Int) ([32]byte, [32]byte, [32]byte, error) {
	return _TeamJust.Contract.CheckSignersByName(&_TeamJust.CallOpts, _whatFunction, _signerA, _signerB, _signerC)
}

// DevCount is a free data retrieval call binding the contract method 0xcebc141a.
//
// Solidity: function devCount() constant returns(uint256)
func (_TeamJust *TeamJustCaller) DevCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TeamJust.contract.Call(opts, out, "devCount")
	return *ret0, err
}

// DevCount is a free data retrieval call binding the contract method 0xcebc141a.
//
// Solidity: function devCount() constant returns(uint256)
func (_TeamJust *TeamJustSession) DevCount() (*big.Int, error) {
	return _TeamJust.Contract.DevCount(&_TeamJust.CallOpts)
}

// DevCount is a free data retrieval call binding the contract method 0xcebc141a.
//
// Solidity: function devCount() constant returns(uint256)
func (_TeamJust *TeamJustCallerSession) DevCount() (*big.Int, error) {
	return _TeamJust.Contract.DevCount(&_TeamJust.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(_who address) constant returns(bool)
func (_TeamJust *TeamJustCaller) IsAdmin(opts *bind.CallOpts, _who common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TeamJust.contract.Call(opts, out, "isAdmin", _who)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(_who address) constant returns(bool)
func (_TeamJust *TeamJustSession) IsAdmin(_who common.Address) (bool, error) {
	return _TeamJust.Contract.IsAdmin(&_TeamJust.CallOpts, _who)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(_who address) constant returns(bool)
func (_TeamJust *TeamJustCallerSession) IsAdmin(_who common.Address) (bool, error) {
	return _TeamJust.Contract.IsAdmin(&_TeamJust.CallOpts, _who)
}

// IsDev is a free data retrieval call binding the contract method 0x0c3f64bf.
//
// Solidity: function isDev(_who address) constant returns(bool)
func (_TeamJust *TeamJustCaller) IsDev(opts *bind.CallOpts, _who common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TeamJust.contract.Call(opts, out, "isDev", _who)
	return *ret0, err
}

// IsDev is a free data retrieval call binding the contract method 0x0c3f64bf.
//
// Solidity: function isDev(_who address) constant returns(bool)
func (_TeamJust *TeamJustSession) IsDev(_who common.Address) (bool, error) {
	return _TeamJust.Contract.IsDev(&_TeamJust.CallOpts, _who)
}

// IsDev is a free data retrieval call binding the contract method 0x0c3f64bf.
//
// Solidity: function isDev(_who address) constant returns(bool)
func (_TeamJust *TeamJustCallerSession) IsDev(_who common.Address) (bool, error) {
	return _TeamJust.Contract.IsDev(&_TeamJust.CallOpts, _who)
}

// RequiredDevSignatures is a free data retrieval call binding the contract method 0xfcf2f85f.
//
// Solidity: function requiredDevSignatures() constant returns(uint256)
func (_TeamJust *TeamJustCaller) RequiredDevSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TeamJust.contract.Call(opts, out, "requiredDevSignatures")
	return *ret0, err
}

// RequiredDevSignatures is a free data retrieval call binding the contract method 0xfcf2f85f.
//
// Solidity: function requiredDevSignatures() constant returns(uint256)
func (_TeamJust *TeamJustSession) RequiredDevSignatures() (*big.Int, error) {
	return _TeamJust.Contract.RequiredDevSignatures(&_TeamJust.CallOpts)
}

// RequiredDevSignatures is a free data retrieval call binding the contract method 0xfcf2f85f.
//
// Solidity: function requiredDevSignatures() constant returns(uint256)
func (_TeamJust *TeamJustCallerSession) RequiredDevSignatures() (*big.Int, error) {
	return _TeamJust.Contract.RequiredDevSignatures(&_TeamJust.CallOpts)
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_TeamJust *TeamJustCaller) RequiredSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TeamJust.contract.Call(opts, out, "requiredSignatures")
	return *ret0, err
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_TeamJust *TeamJustSession) RequiredSignatures() (*big.Int, error) {
	return _TeamJust.Contract.RequiredSignatures(&_TeamJust.CallOpts)
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_TeamJust *TeamJustCallerSession) RequiredSignatures() (*big.Int, error) {
	return _TeamJust.Contract.RequiredSignatures(&_TeamJust.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x372cd183.
//
// Solidity: function addAdmin(_who address, _name bytes32, _isDev bool) returns()
func (_TeamJust *TeamJustTransactor) AddAdmin(opts *bind.TransactOpts, _who common.Address, _name [32]byte, _isDev bool) (*types.Transaction, error) {
	return _TeamJust.contract.Transact(opts, "addAdmin", _who, _name, _isDev)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x372cd183.
//
// Solidity: function addAdmin(_who address, _name bytes32, _isDev bool) returns()
func (_TeamJust *TeamJustSession) AddAdmin(_who common.Address, _name [32]byte, _isDev bool) (*types.Transaction, error) {
	return _TeamJust.Contract.AddAdmin(&_TeamJust.TransactOpts, _who, _name, _isDev)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x372cd183.
//
// Solidity: function addAdmin(_who address, _name bytes32, _isDev bool) returns()
func (_TeamJust *TeamJustTransactorSession) AddAdmin(_who common.Address, _name [32]byte, _isDev bool) (*types.Transaction, error) {
	return _TeamJust.Contract.AddAdmin(&_TeamJust.TransactOpts, _who, _name, _isDev)
}

// ChangeRequiredDevSignatures is a paid mutator transaction binding the contract method 0x2c296656.
//
// Solidity: function changeRequiredDevSignatures(_howMany uint256) returns()
func (_TeamJust *TeamJustTransactor) ChangeRequiredDevSignatures(opts *bind.TransactOpts, _howMany *big.Int) (*types.Transaction, error) {
	return _TeamJust.contract.Transact(opts, "changeRequiredDevSignatures", _howMany)
}

// ChangeRequiredDevSignatures is a paid mutator transaction binding the contract method 0x2c296656.
//
// Solidity: function changeRequiredDevSignatures(_howMany uint256) returns()
func (_TeamJust *TeamJustSession) ChangeRequiredDevSignatures(_howMany *big.Int) (*types.Transaction, error) {
	return _TeamJust.Contract.ChangeRequiredDevSignatures(&_TeamJust.TransactOpts, _howMany)
}

// ChangeRequiredDevSignatures is a paid mutator transaction binding the contract method 0x2c296656.
//
// Solidity: function changeRequiredDevSignatures(_howMany uint256) returns()
func (_TeamJust *TeamJustTransactorSession) ChangeRequiredDevSignatures(_howMany *big.Int) (*types.Transaction, error) {
	return _TeamJust.Contract.ChangeRequiredDevSignatures(&_TeamJust.TransactOpts, _howMany)
}

// ChangeRequiredSignatures is a paid mutator transaction binding the contract method 0x39f636ab.
//
// Solidity: function changeRequiredSignatures(_howMany uint256) returns()
func (_TeamJust *TeamJustTransactor) ChangeRequiredSignatures(opts *bind.TransactOpts, _howMany *big.Int) (*types.Transaction, error) {
	return _TeamJust.contract.Transact(opts, "changeRequiredSignatures", _howMany)
}

// ChangeRequiredSignatures is a paid mutator transaction binding the contract method 0x39f636ab.
//
// Solidity: function changeRequiredSignatures(_howMany uint256) returns()
func (_TeamJust *TeamJustSession) ChangeRequiredSignatures(_howMany *big.Int) (*types.Transaction, error) {
	return _TeamJust.Contract.ChangeRequiredSignatures(&_TeamJust.TransactOpts, _howMany)
}

// ChangeRequiredSignatures is a paid mutator transaction binding the contract method 0x39f636ab.
//
// Solidity: function changeRequiredSignatures(_howMany uint256) returns()
func (_TeamJust *TeamJustTransactorSession) ChangeRequiredSignatures(_howMany *big.Int) (*types.Transaction, error) {
	return _TeamJust.Contract.ChangeRequiredSignatures(&_TeamJust.TransactOpts, _howMany)
}

// DeleteAnyProposal is a paid mutator transaction binding the contract method 0x0efcf295.
//
// Solidity: function deleteAnyProposal(_whatFunction bytes32) returns()
func (_TeamJust *TeamJustTransactor) DeleteAnyProposal(opts *bind.TransactOpts, _whatFunction [32]byte) (*types.Transaction, error) {
	return _TeamJust.contract.Transact(opts, "deleteAnyProposal", _whatFunction)
}

// DeleteAnyProposal is a paid mutator transaction binding the contract method 0x0efcf295.
//
// Solidity: function deleteAnyProposal(_whatFunction bytes32) returns()
func (_TeamJust *TeamJustSession) DeleteAnyProposal(_whatFunction [32]byte) (*types.Transaction, error) {
	return _TeamJust.Contract.DeleteAnyProposal(&_TeamJust.TransactOpts, _whatFunction)
}

// DeleteAnyProposal is a paid mutator transaction binding the contract method 0x0efcf295.
//
// Solidity: function deleteAnyProposal(_whatFunction bytes32) returns()
func (_TeamJust *TeamJustTransactorSession) DeleteAnyProposal(_whatFunction [32]byte) (*types.Transaction, error) {
	return _TeamJust.Contract.DeleteAnyProposal(&_TeamJust.TransactOpts, _whatFunction)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(_who address) returns()
func (_TeamJust *TeamJustTransactor) RemoveAdmin(opts *bind.TransactOpts, _who common.Address) (*types.Transaction, error) {
	return _TeamJust.contract.Transact(opts, "removeAdmin", _who)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(_who address) returns()
func (_TeamJust *TeamJustSession) RemoveAdmin(_who common.Address) (*types.Transaction, error) {
	return _TeamJust.Contract.RemoveAdmin(&_TeamJust.TransactOpts, _who)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(_who address) returns()
func (_TeamJust *TeamJustTransactorSession) RemoveAdmin(_who common.Address) (*types.Transaction, error) {
	return _TeamJust.Contract.RemoveAdmin(&_TeamJust.TransactOpts, _who)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(_addr address) returns()
func (_TeamJust *TeamJustTransactor) Setup(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _TeamJust.contract.Transact(opts, "setup", _addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(_addr address) returns()
func (_TeamJust *TeamJustSession) Setup(_addr common.Address) (*types.Transaction, error) {
	return _TeamJust.Contract.Setup(&_TeamJust.TransactOpts, _addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(_addr address) returns()
func (_TeamJust *TeamJustTransactorSession) Setup(_addr common.Address) (*types.Transaction, error) {
	return _TeamJust.Contract.Setup(&_TeamJust.TransactOpts, _addr)
}
