// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wxmr

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WxmrABI is the input ABI used to generate the binding from.
const WxmrABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_burnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_moneroAddressHash\",\"type\":\"bytes32\"}],\"name\":\"CoinsBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"}],\"name\":\"CoinsMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_new\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_previous\",\"type\":\"bytes\"}],\"name\":\"ReserveProofUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_moneroAddressHash\",\"type\":\"bytes32\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestReserveProof\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"postReserveProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setExchangeRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WxmrBin is the compiled bytecode used for deploying new contracts.
var WxmrBin = "0x60806040526000600255600060035534801561001a57600080fd5b5060008054336001600160a01b0319918216811783556001805490921617905561115a90819061004a90396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c80638da5cb5b116100b8578063d87063531161007c578063d8706353146103b2578063db068e0e14610458578063dd62ed3e14610475578063df7b5c3f146104a3578063f2fde38b146104ab578063f851a440146104d157610137565b80638da5cb5b1461031357806395d89b411461013c578063a9059cbb14610337578063bcf64e0514610363578063d73dd6231461038657610137565b80633385fdc3116100ff5780633385fdc31461026757806340c10f191461026f578063661884631461029b578063704b6c02146102c757806370a08231146102ed57610137565b806306fdde031461013c578063095ea7b3146101b957806318160ddd146101f957806323b872dd14610213578063313ce56714610249575b600080fd5b6101446104d9565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561017e578181015183820152602001610166565b50505050905090810190601f1680156101ab5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101e5600480360360408110156101cf57600080fd5b506001600160a01b0381351690602001356104f9565b604080519115158252519081900360200190f35b61020161055f565b60408051918252519081900360200190f35b6101e56004803603606081101561022957600080fd5b506001600160a01b03813581169160208101359091169060400135610565565b6102516106a6565b6040805160ff9092168252519081900360200190f35b6102016106ab565b6101e56004803603604081101561028557600080fd5b506001600160a01b0381351690602001356106b1565b6101e5600480360360408110156102b157600080fd5b506001600160a01b0381351690602001356107a1565b6101e5600480360360208110156102dd57600080fd5b50356001600160a01b031661088a565b6102016004803603602081101561030357600080fd5b50356001600160a01b031661096e565b61031b610989565b604080516001600160a01b039092168252519081900360200190f35b6101e56004803603604081101561034d57600080fd5b506001600160a01b038135169060200135610998565b6101e56004803603604081101561037957600080fd5b5080359060200135610a84565b6101e56004803603604081101561039c57600080fd5b506001600160a01b038135169060200135610b24565b6101e5600480360360208110156103c857600080fd5b8101906020810181356401000000008111156103e357600080fd5b8201836020820111156103f557600080fd5b8035906020019184600183028401116401000000008311171561041757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610bb7945050505050565b6101e56004803603602081101561046e57600080fd5b5035610d96565b6102016004803603604081101561048b57600080fd5b506001600160a01b0381358116916020013516610d9e565b610144610dc9565b6101e5600480360360208110156104c157600080fd5b50356001600160a01b0316610e57565b61031b610f28565b604051806040016040528060048152602001633bac26a960e11b81525081565b3360008181526006602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60025481565b6000826001600160a01b0381166105be576040805162461bcd60e51b81526020600482015260186024820152776d757374206265206e6f6e207a65726f206164647265737360401b604482015290519081900360640190fd5b6001600160a01b03851660009081526006602090815260408083203384529091529020546105ec9084610f37565b6001600160a01b0386166000818152600660209081526040808320338452825280832094909455918152600590915220546106279084610f37565b6001600160a01b0380871660009081526005602052604080822093909355908616815220546106569084610f80565b6001600160a01b03808616600081815260056020908152604091829020949094558051878152905191939289169260008051602061110583398151915292918290030190a3506001949350505050565b600c81565b60035481565b600080546001600160a01b03163314806106d557506001546001600160a01b031633145b6106de57600080fd5b6001600160a01b0383166000908152600560205260409020546107019083610f80565b6001600160a01b0384166000908152600560205260409020556002546107279083610f80565b6002556040805183815290516001600160a01b038516916000916000805160206111058339815191529181900360200190a36040805183815290516001600160a01b038516917fc9ac6f335722ba5012b311e37c325450961c6e224ea391135b78d95118e190c9919081900360200190a250600192915050565b3360009081526006602090815260408083206001600160a01b03861684529091528120548083106107f5573360009081526006602090815260408083206001600160a01b0388168452909152812055610824565b6107ff8184610f37565b3360009081526006602090815260408083206001600160a01b03891684529091529020555b3360008181526006602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600080546001600160a01b031633146108a257600080fd5b816001600160a01b0381166108f9576040805162461bcd60e51b81526020600482015260186024820152776d757374206265206e6f6e207a65726f206164647265737360401b604482015290519081900360640190fd5b6001546001600160a01b038481169116141561091457600080fd5b600180546001600160a01b0385166001600160a01b0319909116811790915560408051918252517f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c9181900360200190a150600192915050565b6001600160a01b031660009081526005602052604090205490565b6000546001600160a01b031681565b6000826001600160a01b0381166109f1576040805162461bcd60e51b81526020600482015260186024820152776d757374206265206e6f6e207a65726f206164647265737360401b604482015290519081900360640190fd5b33600090815260056020526040902054610a0b9084610f37565b33600090815260056020526040808220929092556001600160a01b03861681522054610a379084610f80565b6001600160a01b0385166000818152600560209081526040918290209390935580518681529051919233926000805160206111058339815191529281900390910190a35060019392505050565b33600090815260056020526040812054610a9e9084610f37565b33600090815260056020526040902055600254610abb9084610f37565b60025560408051848152905160009133916000805160206111058339815191529181900360200190a36040805184815260208101849052815133927fbfe5be076e85df6616ae2ba97cc602211875bbf1c4458fd10012104b899d8f90928290030190a292915050565b3360009081526006602090815260408083206001600160a01b0386168452909152812054610b529083610f80565b3360008181526006602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600080546001600160a01b0316331480610bdb57506001546001600160a01b031633145b610be457600080fd5b60048054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610c705780601f10610c4557610100808354040283529160200191610c70565b820191906000526020600020905b815481529060010190602001808311610c5357829003601f168201915b50508651939450610c8c93600493506020880192509050611071565b507f566a252338bdd070ece239450e646cb7d583c610ec27f166eae07cf15ab833148382604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610cf1578181015183820152602001610cd9565b50505050905090810190601f168015610d1e5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610d51578181015183820152602001610d39565b50505050905090810190601f168015610d7e5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a150600192915050565b600355600190565b6001600160a01b03918216600090815260066020908152604080832093909416825291909152205490565b6004805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610e4f5780601f10610e2457610100808354040283529160200191610e4f565b820191906000526020600020905b815481529060010190602001808311610e3257829003601f168201915b505050505081565b600080546001600160a01b03163314610e6f57600080fd5b816001600160a01b038116610ec6576040805162461bcd60e51b81526020600482015260186024820152776d757374206265206e6f6e207a65726f206164647265737360401b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b03851690811790915560408051338152602081019290925280517f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09281900390910190a150600192915050565b6001546001600160a01b031681565b6000610f7983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610fda565b9392505050565b600082820183811015610f79576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600081848411156110695760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561102e578181015183820152602001611016565b50505050905090810190601f16801561105b5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110b257805160ff19168380011785556110df565b828001600101855582156110df579182015b828111156110df5782518255916020019190600101906110c4565b506110eb9291506110ef565b5090565b5b808211156110eb57600081556001016110f056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa2646970667358221220a8b219169752891f125367d34eb0d3f3974df42bed7d4daf5845dd84a2a7189064736f6c63430007010033"

// DeployWxmr deploys a new Ethereum contract, binding an instance of Wxmr to it.
func DeployWxmr(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Wxmr, error) {
	parsed, err := abi.JSON(strings.NewReader(WxmrABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WxmrBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wxmr{WxmrCaller: WxmrCaller{contract: contract}, WxmrTransactor: WxmrTransactor{contract: contract}, WxmrFilterer: WxmrFilterer{contract: contract}}, nil
}

// Wxmr is an auto generated Go binding around an Ethereum contract.
type Wxmr struct {
	WxmrCaller     // Read-only binding to the contract
	WxmrTransactor // Write-only binding to the contract
	WxmrFilterer   // Log filterer for contract events
}

// WxmrCaller is an auto generated read-only Go binding around an Ethereum contract.
type WxmrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WxmrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WxmrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WxmrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WxmrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WxmrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WxmrSession struct {
	Contract     *Wxmr             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WxmrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WxmrCallerSession struct {
	Contract *WxmrCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WxmrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WxmrTransactorSession struct {
	Contract     *WxmrTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WxmrRaw is an auto generated low-level Go binding around an Ethereum contract.
type WxmrRaw struct {
	Contract *Wxmr // Generic contract binding to access the raw methods on
}

// WxmrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WxmrCallerRaw struct {
	Contract *WxmrCaller // Generic read-only contract binding to access the raw methods on
}

// WxmrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WxmrTransactorRaw struct {
	Contract *WxmrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWxmr creates a new instance of Wxmr, bound to a specific deployed contract.
func NewWxmr(address common.Address, backend bind.ContractBackend) (*Wxmr, error) {
	contract, err := bindWxmr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wxmr{WxmrCaller: WxmrCaller{contract: contract}, WxmrTransactor: WxmrTransactor{contract: contract}, WxmrFilterer: WxmrFilterer{contract: contract}}, nil
}

// NewWxmrCaller creates a new read-only instance of Wxmr, bound to a specific deployed contract.
func NewWxmrCaller(address common.Address, caller bind.ContractCaller) (*WxmrCaller, error) {
	contract, err := bindWxmr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WxmrCaller{contract: contract}, nil
}

// NewWxmrTransactor creates a new write-only instance of Wxmr, bound to a specific deployed contract.
func NewWxmrTransactor(address common.Address, transactor bind.ContractTransactor) (*WxmrTransactor, error) {
	contract, err := bindWxmr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WxmrTransactor{contract: contract}, nil
}

// NewWxmrFilterer creates a new log filterer instance of Wxmr, bound to a specific deployed contract.
func NewWxmrFilterer(address common.Address, filterer bind.ContractFilterer) (*WxmrFilterer, error) {
	contract, err := bindWxmr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WxmrFilterer{contract: contract}, nil
}

// bindWxmr binds a generic wrapper to an already deployed contract.
func bindWxmr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WxmrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wxmr *WxmrRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wxmr.Contract.WxmrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wxmr *WxmrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wxmr.Contract.WxmrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wxmr *WxmrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wxmr.Contract.WxmrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wxmr *WxmrCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wxmr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wxmr *WxmrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wxmr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wxmr *WxmrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wxmr.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Wxmr *WxmrCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wxmr.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Wxmr *WxmrSession) Admin() (common.Address, error) {
	return _Wxmr.Contract.Admin(&_Wxmr.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Wxmr *WxmrCallerSession) Admin() (common.Address, error) {
	return _Wxmr.Contract.Admin(&_Wxmr.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Wxmr *WxmrCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wxmr.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Wxmr *WxmrSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Wxmr.Contract.Allowance(&_Wxmr.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Wxmr *WxmrCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Wxmr.Contract.Allowance(&_Wxmr.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _holder) view returns(uint256)
func (_Wxmr *WxmrCaller) BalanceOf(opts *bind.CallOpts, _holder common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wxmr.contract.Call(opts, out, "balanceOf", _holder)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _holder) view returns(uint256)
func (_Wxmr *WxmrSession) BalanceOf(_holder common.Address) (*big.Int, error) {
	return _Wxmr.Contract.BalanceOf(&_Wxmr.CallOpts, _holder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _holder) view returns(uint256)
func (_Wxmr *WxmrCallerSession) BalanceOf(_holder common.Address) (*big.Int, error) {
	return _Wxmr.Contract.BalanceOf(&_Wxmr.CallOpts, _holder)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Wxmr *WxmrCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Wxmr.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Wxmr *WxmrSession) Decimals() (uint8, error) {
	return _Wxmr.Contract.Decimals(&_Wxmr.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Wxmr *WxmrCallerSession) Decimals() (uint8, error) {
	return _Wxmr.Contract.Decimals(&_Wxmr.CallOpts)
}

// ExchangeRateUSD is a free data retrieval call binding the contract method 0x3385fdc3.
//
// Solidity: function exchangeRateUSD() view returns(uint256)
func (_Wxmr *WxmrCaller) ExchangeRateUSD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wxmr.contract.Call(opts, out, "exchangeRateUSD")
	return *ret0, err
}

// ExchangeRateUSD is a free data retrieval call binding the contract method 0x3385fdc3.
//
// Solidity: function exchangeRateUSD() view returns(uint256)
func (_Wxmr *WxmrSession) ExchangeRateUSD() (*big.Int, error) {
	return _Wxmr.Contract.ExchangeRateUSD(&_Wxmr.CallOpts)
}

// ExchangeRateUSD is a free data retrieval call binding the contract method 0x3385fdc3.
//
// Solidity: function exchangeRateUSD() view returns(uint256)
func (_Wxmr *WxmrCallerSession) ExchangeRateUSD() (*big.Int, error) {
	return _Wxmr.Contract.ExchangeRateUSD(&_Wxmr.CallOpts)
}

// LatestReserveProof is a free data retrieval call binding the contract method 0xdf7b5c3f.
//
// Solidity: function latestReserveProof() view returns(bytes)
func (_Wxmr *WxmrCaller) LatestReserveProof(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Wxmr.contract.Call(opts, out, "latestReserveProof")
	return *ret0, err
}

// LatestReserveProof is a free data retrieval call binding the contract method 0xdf7b5c3f.
//
// Solidity: function latestReserveProof() view returns(bytes)
func (_Wxmr *WxmrSession) LatestReserveProof() ([]byte, error) {
	return _Wxmr.Contract.LatestReserveProof(&_Wxmr.CallOpts)
}

// LatestReserveProof is a free data retrieval call binding the contract method 0xdf7b5c3f.
//
// Solidity: function latestReserveProof() view returns(bytes)
func (_Wxmr *WxmrCallerSession) LatestReserveProof() ([]byte, error) {
	return _Wxmr.Contract.LatestReserveProof(&_Wxmr.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wxmr *WxmrCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Wxmr.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wxmr *WxmrSession) Name() (string, error) {
	return _Wxmr.Contract.Name(&_Wxmr.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wxmr *WxmrCallerSession) Name() (string, error) {
	return _Wxmr.Contract.Name(&_Wxmr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wxmr *WxmrCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wxmr.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wxmr *WxmrSession) Owner() (common.Address, error) {
	return _Wxmr.Contract.Owner(&_Wxmr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wxmr *WxmrCallerSession) Owner() (common.Address, error) {
	return _Wxmr.Contract.Owner(&_Wxmr.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wxmr *WxmrCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Wxmr.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wxmr *WxmrSession) Symbol() (string, error) {
	return _Wxmr.Contract.Symbol(&_Wxmr.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wxmr *WxmrCallerSession) Symbol() (string, error) {
	return _Wxmr.Contract.Symbol(&_Wxmr.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wxmr *WxmrCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wxmr.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wxmr *WxmrSession) TotalSupply() (*big.Int, error) {
	return _Wxmr.Contract.TotalSupply(&_Wxmr.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wxmr *WxmrCallerSession) TotalSupply() (*big.Int, error) {
	return _Wxmr.Contract.TotalSupply(&_Wxmr.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Wxmr *WxmrTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Wxmr.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Wxmr *WxmrSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.Approve(&_Wxmr.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Wxmr *WxmrTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.Approve(&_Wxmr.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0xbcf64e05.
//
// Solidity: function burn(uint256 _amount, bytes32 _moneroAddressHash) returns(bool)
func (_Wxmr *WxmrTransactor) Burn(opts *bind.TransactOpts, _amount *big.Int, _moneroAddressHash [32]byte) (*types.Transaction, error) {
	return _Wxmr.contract.Transact(opts, "burn", _amount, _moneroAddressHash)
}

// Burn is a paid mutator transaction binding the contract method 0xbcf64e05.
//
// Solidity: function burn(uint256 _amount, bytes32 _moneroAddressHash) returns(bool)
func (_Wxmr *WxmrSession) Burn(_amount *big.Int, _moneroAddressHash [32]byte) (*types.Transaction, error) {
	return _Wxmr.Contract.Burn(&_Wxmr.TransactOpts, _amount, _moneroAddressHash)
}

// Burn is a paid mutator transaction binding the contract method 0xbcf64e05.
//
// Solidity: function burn(uint256 _amount, bytes32 _moneroAddressHash) returns(bool)
func (_Wxmr *WxmrTransactorSession) Burn(_amount *big.Int, _moneroAddressHash [32]byte) (*types.Transaction, error) {
	return _Wxmr.Contract.Burn(&_Wxmr.TransactOpts, _amount, _moneroAddressHash)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_Wxmr *WxmrTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Wxmr.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_Wxmr *WxmrSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.DecreaseApproval(&_Wxmr.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_Wxmr *WxmrTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.DecreaseApproval(&_Wxmr.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_Wxmr *WxmrTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Wxmr.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_Wxmr *WxmrSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.IncreaseApproval(&_Wxmr.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_Wxmr *WxmrTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.IncreaseApproval(&_Wxmr.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _recipient, uint256 _amount) returns(bool)
func (_Wxmr *WxmrTransactor) Mint(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wxmr.contract.Transact(opts, "mint", _recipient, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _recipient, uint256 _amount) returns(bool)
func (_Wxmr *WxmrSession) Mint(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.Mint(&_Wxmr.TransactOpts, _recipient, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _recipient, uint256 _amount) returns(bool)
func (_Wxmr *WxmrTransactorSession) Mint(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.Mint(&_Wxmr.TransactOpts, _recipient, _amount)
}

// PostReserveProof is a paid mutator transaction binding the contract method 0xd8706353.
//
// Solidity: function postReserveProof(bytes _proof) returns(bool)
func (_Wxmr *WxmrTransactor) PostReserveProof(opts *bind.TransactOpts, _proof []byte) (*types.Transaction, error) {
	return _Wxmr.contract.Transact(opts, "postReserveProof", _proof)
}

// PostReserveProof is a paid mutator transaction binding the contract method 0xd8706353.
//
// Solidity: function postReserveProof(bytes _proof) returns(bool)
func (_Wxmr *WxmrSession) PostReserveProof(_proof []byte) (*types.Transaction, error) {
	return _Wxmr.Contract.PostReserveProof(&_Wxmr.TransactOpts, _proof)
}

// PostReserveProof is a paid mutator transaction binding the contract method 0xd8706353.
//
// Solidity: function postReserveProof(bytes _proof) returns(bool)
func (_Wxmr *WxmrTransactorSession) PostReserveProof(_proof []byte) (*types.Transaction, error) {
	return _Wxmr.Contract.PostReserveProof(&_Wxmr.TransactOpts, _proof)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Wxmr *WxmrTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Wxmr.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Wxmr *WxmrSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Wxmr.Contract.SetAdmin(&_Wxmr.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Wxmr *WxmrTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Wxmr.Contract.SetAdmin(&_Wxmr.TransactOpts, _newAdmin)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 _rate) returns(bool)
func (_Wxmr *WxmrTransactor) SetExchangeRate(opts *bind.TransactOpts, _rate *big.Int) (*types.Transaction, error) {
	return _Wxmr.contract.Transact(opts, "setExchangeRate", _rate)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 _rate) returns(bool)
func (_Wxmr *WxmrSession) SetExchangeRate(_rate *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.SetExchangeRate(&_Wxmr.TransactOpts, _rate)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 _rate) returns(bool)
func (_Wxmr *WxmrTransactorSession) SetExchangeRate(_rate *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.SetExchangeRate(&_Wxmr.TransactOpts, _rate)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Wxmr *WxmrTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wxmr.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Wxmr *WxmrSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.Transfer(&_Wxmr.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Wxmr *WxmrTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.Transfer(&_Wxmr.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _recipient, uint256 _amount) returns(bool)
func (_Wxmr *WxmrTransactor) TransferFrom(opts *bind.TransactOpts, _owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wxmr.contract.Transact(opts, "transferFrom", _owner, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _recipient, uint256 _amount) returns(bool)
func (_Wxmr *WxmrSession) TransferFrom(_owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.TransferFrom(&_Wxmr.TransactOpts, _owner, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _recipient, uint256 _amount) returns(bool)
func (_Wxmr *WxmrTransactorSession) TransferFrom(_owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wxmr.Contract.TransferFrom(&_Wxmr.TransactOpts, _owner, _recipient, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Wxmr *WxmrTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Wxmr.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Wxmr *WxmrSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Wxmr.Contract.TransferOwnership(&_Wxmr.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Wxmr *WxmrTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Wxmr.Contract.TransferOwnership(&_Wxmr.TransactOpts, _newOwner)
}

// WxmrAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the Wxmr contract.
type WxmrAdminSetIterator struct {
	Event *WxmrAdminSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WxmrAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WxmrAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WxmrAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WxmrAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WxmrAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WxmrAdminSet represents a AdminSet event raised by the Wxmr contract.
type WxmrAdminSet struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(address _admin)
func (_Wxmr *WxmrFilterer) FilterAdminSet(opts *bind.FilterOpts) (*WxmrAdminSetIterator, error) {

	logs, sub, err := _Wxmr.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &WxmrAdminSetIterator{contract: _Wxmr.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(address _admin)
func (_Wxmr *WxmrFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *WxmrAdminSet) (event.Subscription, error) {

	logs, sub, err := _Wxmr.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WxmrAdminSet)
				if err := _Wxmr.contract.UnpackLog(event, "AdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminSet is a log parse operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(address _admin)
func (_Wxmr *WxmrFilterer) ParseAdminSet(log types.Log) (*WxmrAdminSet, error) {
	event := new(WxmrAdminSet)
	if err := _Wxmr.contract.UnpackLog(event, "AdminSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WxmrApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Wxmr contract.
type WxmrApprovalIterator struct {
	Event *WxmrApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WxmrApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WxmrApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WxmrApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WxmrApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WxmrApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WxmrApproval represents a Approval event raised by the Wxmr contract.
type WxmrApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _amount)
func (_Wxmr *WxmrFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*WxmrApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Wxmr.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &WxmrApprovalIterator{contract: _Wxmr.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _amount)
func (_Wxmr *WxmrFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WxmrApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Wxmr.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WxmrApproval)
				if err := _Wxmr.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _amount)
func (_Wxmr *WxmrFilterer) ParseApproval(log types.Log) (*WxmrApproval, error) {
	event := new(WxmrApproval)
	if err := _Wxmr.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WxmrCoinsBurnedIterator is returned from FilterCoinsBurned and is used to iterate over the raw logs and unpacked data for CoinsBurned events raised by the Wxmr contract.
type WxmrCoinsBurnedIterator struct {
	Event *WxmrCoinsBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WxmrCoinsBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WxmrCoinsBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WxmrCoinsBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WxmrCoinsBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WxmrCoinsBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WxmrCoinsBurned represents a CoinsBurned event raised by the Wxmr contract.
type WxmrCoinsBurned struct {
	Burner            common.Address
	BurnAmount        *big.Int
	MoneroAddressHash [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCoinsBurned is a free log retrieval operation binding the contract event 0xbfe5be076e85df6616ae2ba97cc602211875bbf1c4458fd10012104b899d8f90.
//
// Solidity: event CoinsBurned(address indexed _burner, uint256 _burnAmount, bytes32 _moneroAddressHash)
func (_Wxmr *WxmrFilterer) FilterCoinsBurned(opts *bind.FilterOpts, _burner []common.Address) (*WxmrCoinsBurnedIterator, error) {

	var _burnerRule []interface{}
	for _, _burnerItem := range _burner {
		_burnerRule = append(_burnerRule, _burnerItem)
	}

	logs, sub, err := _Wxmr.contract.FilterLogs(opts, "CoinsBurned", _burnerRule)
	if err != nil {
		return nil, err
	}
	return &WxmrCoinsBurnedIterator{contract: _Wxmr.contract, event: "CoinsBurned", logs: logs, sub: sub}, nil
}

// WatchCoinsBurned is a free log subscription operation binding the contract event 0xbfe5be076e85df6616ae2ba97cc602211875bbf1c4458fd10012104b899d8f90.
//
// Solidity: event CoinsBurned(address indexed _burner, uint256 _burnAmount, bytes32 _moneroAddressHash)
func (_Wxmr *WxmrFilterer) WatchCoinsBurned(opts *bind.WatchOpts, sink chan<- *WxmrCoinsBurned, _burner []common.Address) (event.Subscription, error) {

	var _burnerRule []interface{}
	for _, _burnerItem := range _burner {
		_burnerRule = append(_burnerRule, _burnerItem)
	}

	logs, sub, err := _Wxmr.contract.WatchLogs(opts, "CoinsBurned", _burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WxmrCoinsBurned)
				if err := _Wxmr.contract.UnpackLog(event, "CoinsBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCoinsBurned is a log parse operation binding the contract event 0xbfe5be076e85df6616ae2ba97cc602211875bbf1c4458fd10012104b899d8f90.
//
// Solidity: event CoinsBurned(address indexed _burner, uint256 _burnAmount, bytes32 _moneroAddressHash)
func (_Wxmr *WxmrFilterer) ParseCoinsBurned(log types.Log) (*WxmrCoinsBurned, error) {
	event := new(WxmrCoinsBurned)
	if err := _Wxmr.contract.UnpackLog(event, "CoinsBurned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WxmrCoinsMintedIterator is returned from FilterCoinsMinted and is used to iterate over the raw logs and unpacked data for CoinsMinted events raised by the Wxmr contract.
type WxmrCoinsMintedIterator struct {
	Event *WxmrCoinsMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WxmrCoinsMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WxmrCoinsMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WxmrCoinsMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WxmrCoinsMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WxmrCoinsMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WxmrCoinsMinted represents a CoinsMinted event raised by the Wxmr contract.
type WxmrCoinsMinted struct {
	Recipient  common.Address
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCoinsMinted is a free log retrieval operation binding the contract event 0xc9ac6f335722ba5012b311e37c325450961c6e224ea391135b78d95118e190c9.
//
// Solidity: event CoinsMinted(address indexed _recipient, uint256 _mintAmount)
func (_Wxmr *WxmrFilterer) FilterCoinsMinted(opts *bind.FilterOpts, _recipient []common.Address) (*WxmrCoinsMintedIterator, error) {

	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _Wxmr.contract.FilterLogs(opts, "CoinsMinted", _recipientRule)
	if err != nil {
		return nil, err
	}
	return &WxmrCoinsMintedIterator{contract: _Wxmr.contract, event: "CoinsMinted", logs: logs, sub: sub}, nil
}

// WatchCoinsMinted is a free log subscription operation binding the contract event 0xc9ac6f335722ba5012b311e37c325450961c6e224ea391135b78d95118e190c9.
//
// Solidity: event CoinsMinted(address indexed _recipient, uint256 _mintAmount)
func (_Wxmr *WxmrFilterer) WatchCoinsMinted(opts *bind.WatchOpts, sink chan<- *WxmrCoinsMinted, _recipient []common.Address) (event.Subscription, error) {

	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _Wxmr.contract.WatchLogs(opts, "CoinsMinted", _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WxmrCoinsMinted)
				if err := _Wxmr.contract.UnpackLog(event, "CoinsMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCoinsMinted is a log parse operation binding the contract event 0xc9ac6f335722ba5012b311e37c325450961c6e224ea391135b78d95118e190c9.
//
// Solidity: event CoinsMinted(address indexed _recipient, uint256 _mintAmount)
func (_Wxmr *WxmrFilterer) ParseCoinsMinted(log types.Log) (*WxmrCoinsMinted, error) {
	event := new(WxmrCoinsMinted)
	if err := _Wxmr.contract.UnpackLog(event, "CoinsMinted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WxmrOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Wxmr contract.
type WxmrOwnershipTransferredIterator struct {
	Event *WxmrOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WxmrOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WxmrOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WxmrOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WxmrOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WxmrOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WxmrOwnershipTransferred represents a OwnershipTransferred event raised by the Wxmr contract.
type WxmrOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address _previousOwner, address _newOwner)
func (_Wxmr *WxmrFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*WxmrOwnershipTransferredIterator, error) {

	logs, sub, err := _Wxmr.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &WxmrOwnershipTransferredIterator{contract: _Wxmr.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address _previousOwner, address _newOwner)
func (_Wxmr *WxmrFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WxmrOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _Wxmr.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WxmrOwnershipTransferred)
				if err := _Wxmr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address _previousOwner, address _newOwner)
func (_Wxmr *WxmrFilterer) ParseOwnershipTransferred(log types.Log) (*WxmrOwnershipTransferred, error) {
	event := new(WxmrOwnershipTransferred)
	if err := _Wxmr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WxmrReserveProofUpdateIterator is returned from FilterReserveProofUpdate and is used to iterate over the raw logs and unpacked data for ReserveProofUpdate events raised by the Wxmr contract.
type WxmrReserveProofUpdateIterator struct {
	Event *WxmrReserveProofUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WxmrReserveProofUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WxmrReserveProofUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WxmrReserveProofUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WxmrReserveProofUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WxmrReserveProofUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WxmrReserveProofUpdate represents a ReserveProofUpdate event raised by the Wxmr contract.
type WxmrReserveProofUpdate struct {
	New      []byte
	Previous []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReserveProofUpdate is a free log retrieval operation binding the contract event 0x566a252338bdd070ece239450e646cb7d583c610ec27f166eae07cf15ab83314.
//
// Solidity: event ReserveProofUpdate(bytes _new, bytes _previous)
func (_Wxmr *WxmrFilterer) FilterReserveProofUpdate(opts *bind.FilterOpts) (*WxmrReserveProofUpdateIterator, error) {

	logs, sub, err := _Wxmr.contract.FilterLogs(opts, "ReserveProofUpdate")
	if err != nil {
		return nil, err
	}
	return &WxmrReserveProofUpdateIterator{contract: _Wxmr.contract, event: "ReserveProofUpdate", logs: logs, sub: sub}, nil
}

// WatchReserveProofUpdate is a free log subscription operation binding the contract event 0x566a252338bdd070ece239450e646cb7d583c610ec27f166eae07cf15ab83314.
//
// Solidity: event ReserveProofUpdate(bytes _new, bytes _previous)
func (_Wxmr *WxmrFilterer) WatchReserveProofUpdate(opts *bind.WatchOpts, sink chan<- *WxmrReserveProofUpdate) (event.Subscription, error) {

	logs, sub, err := _Wxmr.contract.WatchLogs(opts, "ReserveProofUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WxmrReserveProofUpdate)
				if err := _Wxmr.contract.UnpackLog(event, "ReserveProofUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReserveProofUpdate is a log parse operation binding the contract event 0x566a252338bdd070ece239450e646cb7d583c610ec27f166eae07cf15ab83314.
//
// Solidity: event ReserveProofUpdate(bytes _new, bytes _previous)
func (_Wxmr *WxmrFilterer) ParseReserveProofUpdate(log types.Log) (*WxmrReserveProofUpdate, error) {
	event := new(WxmrReserveProofUpdate)
	if err := _Wxmr.contract.UnpackLog(event, "ReserveProofUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WxmrTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Wxmr contract.
type WxmrTransferIterator struct {
	Event *WxmrTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WxmrTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WxmrTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WxmrTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WxmrTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WxmrTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WxmrTransfer represents a Transfer event raised by the Wxmr contract.
type WxmrTransfer struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _sender, address indexed _recipient, uint256 _amount)
func (_Wxmr *WxmrFilterer) FilterTransfer(opts *bind.FilterOpts, _sender []common.Address, _recipient []common.Address) (*WxmrTransferIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _Wxmr.contract.FilterLogs(opts, "Transfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return &WxmrTransferIterator{contract: _Wxmr.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _sender, address indexed _recipient, uint256 _amount)
func (_Wxmr *WxmrFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WxmrTransfer, _sender []common.Address, _recipient []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _Wxmr.contract.WatchLogs(opts, "Transfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WxmrTransfer)
				if err := _Wxmr.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _sender, address indexed _recipient, uint256 _amount)
func (_Wxmr *WxmrFilterer) ParseTransfer(log types.Log) (*WxmrTransfer, error) {
	event := new(WxmrTransfer)
	if err := _Wxmr.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
