// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reserve

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

// ReserveABI is the input ABI used to generate the binding from.
const ReserveABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_reserveAddress\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_burnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_moneroAddressHash\",\"type\":\"bytes32\"}],\"name\":\"CoinsBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"}],\"name\":\"CoinsMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_signatureKeccakHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"ProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_moneroAddressHash\",\"type\":\"bytes32\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastProofHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_reserveAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_signatureKeccakHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_moneroBlockHeight\",\"type\":\"uint256\"}],\"name\":\"postReserveProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proofs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"signatureKeccakHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reserveAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"moneroBlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveAddress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setExchangeRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ReserveBin is the compiled bytecode used for deploying new contracts.
var ReserveBin = "0x60806040526000600355600060065560006007553480156200002057600080fd5b50604051620015c1380380620015c1833981810160405260208110156200004657600080fd5b81019080805160405193929190846401000000008211156200006757600080fd5b9083019060208201858111156200007d57600080fd5b82516401000000008111828201881017156200009857600080fd5b82525081516020918201929091019080838360005b83811015620000c7578181015183820152602001620000ad565b50505050905090810190601f168015620000f55780820380516001836020036101000a031916815260200191505b506040525050600080546001805460028054336001600160a01b03199586168117861681179096559184168517841685179092558216831790911690911790555080516200014b90600890602084019062000153565b5050620001ef565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200019657805160ff1916838001178555620001c6565b82800160010185558215620001c6579182015b82811115620001c6578251825591602001919060010190620001a9565b50620001d4929150620001d8565b5090565b5b80821115620001d45760008155600101620001d9565b6113c280620001ff6000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c80638522da07116100c3578063d73dd6231161007c578063d73dd623146105f9578063db068e0e14610625578063dd62ed3e14610642578063f2fde38b14610670578063f79ed94b14610696578063f851a4401461069e57610158565b80638522da07146103605780638da5cb5b1461049257806395d89b411461015d5780639ddaf5aa1461049a578063a9059cbb146105aa578063bcf64e05146105d657610158565b806340c10f191161011557806340c10f1914610290578063512bcb60146102bc57806366188463146102c4578063704b6c02146102f057806370a0823114610316578063741bef1a1461033c57610158565b806306fdde031461015d578063095ea7b3146101da57806318160ddd1461021a57806323b872dd14610234578063313ce5671461026a5780633385fdc314610288575b600080fd5b6101656106a6565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561019f578181015183820152602001610187565b50505050905090810190601f1680156101cc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610206600480360360408110156101f057600080fd5b506001600160a01b0381351690602001356106c6565b604080519115158252519081900360200190f35b61022261072c565b60408051918252519081900360200190f35b6102066004803603606081101561024a57600080fd5b506001600160a01b03813581169160208101359091169060400135610732565b610272610873565b6040805160ff9092168252519081900360200190f35b610222610878565b610206600480360360408110156102a657600080fd5b506001600160a01b03813516906020013561087e565b61022261096e565b610206600480360360408110156102da57600080fd5b506001600160a01b038135169060200135610974565b6102066004803603602081101561030657600080fd5b50356001600160a01b0316610a5d565b6102226004803603602081101561032c57600080fd5b50356001600160a01b0316610b41565b610344610b5c565b604080516001600160a01b039092168252519081900360200190f35b6102066004803603608081101561037657600080fd5b81019060208101813564010000000081111561039157600080fd5b8201836020820111156103a357600080fd5b803590602001918460018302840111640100000000831117156103c557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561041857600080fd5b82018360208201111561042a57600080fd5b8035906020019184600183028401116401000000008311171561044c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610b6b565b610344610c6a565b6104b7600480360360208110156104b057600080fd5b5035610c79565b604051808681526020018060200180602001858152602001848152602001838103835287818151815260200191508051906020019080838360005b8381101561050a5781810151838201526020016104f2565b50505050905090810190601f1680156105375780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b8381101561056a578181015183820152602001610552565b50505050905090810190601f1680156105975780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b610206600480360360408110156105c057600080fd5b506001600160a01b038135169060200135610dbe565b610206600480360360408110156105ec57600080fd5b5080359060200135610eaa565b6102066004803603604081101561060f57600080fd5b506001600160a01b038135169060200135610f50565b6102066004803603602081101561063b57600080fd5b5035610fe3565b6102226004803603604081101561065857600080fd5b506001600160a01b0381358116916020013516611006565b6102066004803603602081101561068657600080fd5b50356001600160a01b0316611031565b610165611102565b610344611190565b604051806040016040528060048152602001633bac26a960e11b81525081565b3360008181526005602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60035481565b6000826001600160a01b03811661078b576040805162461bcd60e51b81526020600482015260186024820152776d757374206265206e6f6e207a65726f206164647265737360401b604482015290519081900360640190fd5b6001600160a01b03851660009081526005602090815260408083203384529091529020546107b9908461119f565b6001600160a01b0386166000818152600560209081526040808320338452825280832094909455918152600490915220546107f4908461119f565b6001600160a01b03808716600090815260046020526040808220939093559086168152205461082390846111e8565b6001600160a01b03808616600081815260046020908152604091829020949094558051878152905191939289169260008051602061136d83398151915292918290030190a3506001949350505050565b600c81565b60075481565b600080546001600160a01b03163314806108a257506001546001600160a01b031633145b6108ab57600080fd5b6001600160a01b0383166000908152600460205260409020546108ce90836111e8565b6001600160a01b0384166000908152600460205260409020556003546108f490836111e8565b6003556040805183815290516001600160a01b0385169160009160008051602061136d8339815191529181900360200190a36040805183815290516001600160a01b038516917fc9ac6f335722ba5012b311e37c325450961c6e224ea391135b78d95118e190c9919081900360200190a250600192915050565b60065481565b3360009081526005602090815260408083206001600160a01b03861684529091528120548083106109c8573360009081526005602090815260408083206001600160a01b03881684529091528120556109f7565b6109d2818461119f565b3360009081526005602090815260408083206001600160a01b03891684529091529020555b3360008181526005602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600080546001600160a01b03163314610a7557600080fd5b816001600160a01b038116610acc576040805162461bcd60e51b81526020600482015260186024820152776d757374206265206e6f6e207a65726f206164647265737360401b604482015290519081900360640190fd5b6001546001600160a01b0384811691161415610ae757600080fd5b600180546001600160a01b0385166001600160a01b0319909116811790915560408051918252517f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c9181900360200190a150600192915050565b6001600160a01b031660009081526004602052604090205490565b6002546001600160a01b031681565b600080546001600160a01b0316331480610b8f57506001546001600160a01b031633145b610b9857600080fd5b6040805160a081018252848152602080820188815282840188905260608301869052600354608084015243600090815260098352939093208251815592518051929392610beb92600185019201906112d9565b5060408201518051610c079160028401916020909101906112d9565b506060820151600382015560809091015160049091015543600681905560408051858152602081019290925280517f3c0e85acf452e867ee09c9151d2cf45fe6d6a729271347e89a9a6b91b081f31d9281900390910190a1506001949350505050565b6000546001600160a01b031681565b6009602090815260009182526040918290208054600180830180548651600293821615610100026000190190911692909204601f810186900486028301860190965285825291949293909290830182828015610d165780601f10610ceb57610100808354040283529160200191610d16565b820191906000526020600020905b815481529060010190602001808311610cf957829003601f168201915b50505060028085018054604080516020601f6000196101006001871615020190941695909504928301859004850281018501909152818152959695945090925090830182828015610da85780601f10610d7d57610100808354040283529160200191610da8565b820191906000526020600020905b815481529060010190602001808311610d8b57829003601f168201915b5050505050908060030154908060040154905085565b6000826001600160a01b038116610e17576040805162461bcd60e51b81526020600482015260186024820152776d757374206265206e6f6e207a65726f206164647265737360401b604482015290519081900360640190fd5b33600090815260046020526040902054610e31908461119f565b33600090815260046020526040808220929092556001600160a01b03861681522054610e5d90846111e8565b6001600160a01b03851660008181526004602090815260409182902093909355805186815290519192339260008051602061136d8339815191529281900390910190a35060019392505050565b33600090815260046020526040812054610ec4908461119f565b33600090815260046020526040902055600354610ee1908461119f565b600355604080518481529051600091339160008051602061136d8339815191529181900360200190a3604080513381526020810185905280820184905290517fbfe5be076e85df6616ae2ba97cc602211875bbf1c4458fd10012104b899d8f909181900360600190a192915050565b3360009081526005602090815260408083206001600160a01b0386168452909152812054610f7e90836111e8565b3360008181526005602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6002546000906001600160a01b03163314610ffd57600080fd5b50600755600190565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205490565b600080546001600160a01b0316331461104957600080fd5b816001600160a01b0381166110a0576040805162461bcd60e51b81526020600482015260186024820152776d757374206265206e6f6e207a65726f206164647265737360401b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b03851690811790915560408051338152602081019290925280517f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09281900390910190a150600192915050565b6008805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156111885780601f1061115d57610100808354040283529160200191611188565b820191906000526020600020905b81548152906001019060200180831161116b57829003601f168201915b505050505081565b6001546001600160a01b031681565b60006111e183836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611242565b9392505050565b6000828201838110156111e1576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600081848411156112d15760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561129657818101518382015260200161127e565b50505050905090810190601f1680156112c35780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061131a57805160ff1916838001178555611347565b82800160010185558215611347579182015b8281111561134757825182559160200191906001019061132c565b50611353929150611357565b5090565b5b80821115611353576000815560010161135856feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa264697066735822122036083e844a191469505d5737cbd9c4b1d624a6c3e756d92118f5fa303e255b3a64736f6c63430007010033"

// DeployReserve deploys a new Ethereum contract, binding an instance of Reserve to it.
func DeployReserve(auth *bind.TransactOpts, backend bind.ContractBackend, _reserveAddress []byte) (common.Address, *types.Transaction, *Reserve, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReserveBin), backend, _reserveAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reserve{ReserveCaller: ReserveCaller{contract: contract}, ReserveTransactor: ReserveTransactor{contract: contract}, ReserveFilterer: ReserveFilterer{contract: contract}}, nil
}

// Reserve is an auto generated Go binding around an Ethereum contract.
type Reserve struct {
	ReserveCaller     // Read-only binding to the contract
	ReserveTransactor // Write-only binding to the contract
	ReserveFilterer   // Log filterer for contract events
}

// ReserveCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReserveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReserveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReserveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReserveSession struct {
	Contract     *Reserve          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReserveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReserveCallerSession struct {
	Contract *ReserveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ReserveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReserveTransactorSession struct {
	Contract     *ReserveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReserveRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReserveRaw struct {
	Contract *Reserve // Generic contract binding to access the raw methods on
}

// ReserveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReserveCallerRaw struct {
	Contract *ReserveCaller // Generic read-only contract binding to access the raw methods on
}

// ReserveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReserveTransactorRaw struct {
	Contract *ReserveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReserve creates a new instance of Reserve, bound to a specific deployed contract.
func NewReserve(address common.Address, backend bind.ContractBackend) (*Reserve, error) {
	contract, err := bindReserve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reserve{ReserveCaller: ReserveCaller{contract: contract}, ReserveTransactor: ReserveTransactor{contract: contract}, ReserveFilterer: ReserveFilterer{contract: contract}}, nil
}

// NewReserveCaller creates a new read-only instance of Reserve, bound to a specific deployed contract.
func NewReserveCaller(address common.Address, caller bind.ContractCaller) (*ReserveCaller, error) {
	contract, err := bindReserve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveCaller{contract: contract}, nil
}

// NewReserveTransactor creates a new write-only instance of Reserve, bound to a specific deployed contract.
func NewReserveTransactor(address common.Address, transactor bind.ContractTransactor) (*ReserveTransactor, error) {
	contract, err := bindReserve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveTransactor{contract: contract}, nil
}

// NewReserveFilterer creates a new log filterer instance of Reserve, bound to a specific deployed contract.
func NewReserveFilterer(address common.Address, filterer bind.ContractFilterer) (*ReserveFilterer, error) {
	contract, err := bindReserve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReserveFilterer{contract: contract}, nil
}

// bindReserve binds a generic wrapper to an already deployed contract.
func bindReserve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reserve *ReserveRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reserve.Contract.ReserveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reserve *ReserveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.Contract.ReserveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reserve *ReserveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reserve.Contract.ReserveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reserve *ReserveCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reserve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reserve *ReserveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reserve *ReserveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reserve.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Reserve *ReserveCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Reserve *ReserveSession) Admin() (common.Address, error) {
	return _Reserve.Contract.Admin(&_Reserve.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Reserve *ReserveCallerSession) Admin() (common.Address, error) {
	return _Reserve.Contract.Admin(&_Reserve.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Reserve *ReserveCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Reserve *ReserveSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Reserve.Contract.Allowance(&_Reserve.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Reserve *ReserveCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Reserve.Contract.Allowance(&_Reserve.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _holder) view returns(uint256)
func (_Reserve *ReserveCaller) BalanceOf(opts *bind.CallOpts, _holder common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "balanceOf", _holder)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _holder) view returns(uint256)
func (_Reserve *ReserveSession) BalanceOf(_holder common.Address) (*big.Int, error) {
	return _Reserve.Contract.BalanceOf(&_Reserve.CallOpts, _holder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _holder) view returns(uint256)
func (_Reserve *ReserveCallerSession) BalanceOf(_holder common.Address) (*big.Int, error) {
	return _Reserve.Contract.BalanceOf(&_Reserve.CallOpts, _holder)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Reserve *ReserveCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Reserve *ReserveSession) Decimals() (uint8, error) {
	return _Reserve.Contract.Decimals(&_Reserve.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Reserve *ReserveCallerSession) Decimals() (uint8, error) {
	return _Reserve.Contract.Decimals(&_Reserve.CallOpts)
}

// ExchangeRateUSD is a free data retrieval call binding the contract method 0x3385fdc3.
//
// Solidity: function exchangeRateUSD() view returns(uint256)
func (_Reserve *ReserveCaller) ExchangeRateUSD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "exchangeRateUSD")
	return *ret0, err
}

// ExchangeRateUSD is a free data retrieval call binding the contract method 0x3385fdc3.
//
// Solidity: function exchangeRateUSD() view returns(uint256)
func (_Reserve *ReserveSession) ExchangeRateUSD() (*big.Int, error) {
	return _Reserve.Contract.ExchangeRateUSD(&_Reserve.CallOpts)
}

// ExchangeRateUSD is a free data retrieval call binding the contract method 0x3385fdc3.
//
// Solidity: function exchangeRateUSD() view returns(uint256)
func (_Reserve *ReserveCallerSession) ExchangeRateUSD() (*big.Int, error) {
	return _Reserve.Contract.ExchangeRateUSD(&_Reserve.CallOpts)
}

// LastProofHeight is a free data retrieval call binding the contract method 0x512bcb60.
//
// Solidity: function lastProofHeight() view returns(uint256)
func (_Reserve *ReserveCaller) LastProofHeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "lastProofHeight")
	return *ret0, err
}

// LastProofHeight is a free data retrieval call binding the contract method 0x512bcb60.
//
// Solidity: function lastProofHeight() view returns(uint256)
func (_Reserve *ReserveSession) LastProofHeight() (*big.Int, error) {
	return _Reserve.Contract.LastProofHeight(&_Reserve.CallOpts)
}

// LastProofHeight is a free data retrieval call binding the contract method 0x512bcb60.
//
// Solidity: function lastProofHeight() view returns(uint256)
func (_Reserve *ReserveCallerSession) LastProofHeight() (*big.Int, error) {
	return _Reserve.Contract.LastProofHeight(&_Reserve.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Reserve *ReserveCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Reserve *ReserveSession) Name() (string, error) {
	return _Reserve.Contract.Name(&_Reserve.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Reserve *ReserveCallerSession) Name() (string, error) {
	return _Reserve.Contract.Name(&_Reserve.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reserve *ReserveCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reserve *ReserveSession) Owner() (common.Address, error) {
	return _Reserve.Contract.Owner(&_Reserve.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reserve *ReserveCallerSession) Owner() (common.Address, error) {
	return _Reserve.Contract.Owner(&_Reserve.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Reserve *ReserveCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "priceFeed")
	return *ret0, err
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Reserve *ReserveSession) PriceFeed() (common.Address, error) {
	return _Reserve.Contract.PriceFeed(&_Reserve.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Reserve *ReserveCallerSession) PriceFeed() (common.Address, error) {
	return _Reserve.Contract.PriceFeed(&_Reserve.CallOpts)
}

// Proofs is a free data retrieval call binding the contract method 0x9ddaf5aa.
//
// Solidity: function proofs(uint256 ) view returns(bytes32 signatureKeccakHash, bytes message, bytes reserveAddress, uint256 moneroBlockHeight, uint256 totalSupply)
func (_Reserve *ReserveCaller) Proofs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SignatureKeccakHash [32]byte
	Message             []byte
	ReserveAddress      []byte
	MoneroBlockHeight   *big.Int
	TotalSupply         *big.Int
}, error) {
	ret := new(struct {
		SignatureKeccakHash [32]byte
		Message             []byte
		ReserveAddress      []byte
		MoneroBlockHeight   *big.Int
		TotalSupply         *big.Int
	})
	out := ret
	err := _Reserve.contract.Call(opts, out, "proofs", arg0)
	return *ret, err
}

// Proofs is a free data retrieval call binding the contract method 0x9ddaf5aa.
//
// Solidity: function proofs(uint256 ) view returns(bytes32 signatureKeccakHash, bytes message, bytes reserveAddress, uint256 moneroBlockHeight, uint256 totalSupply)
func (_Reserve *ReserveSession) Proofs(arg0 *big.Int) (struct {
	SignatureKeccakHash [32]byte
	Message             []byte
	ReserveAddress      []byte
	MoneroBlockHeight   *big.Int
	TotalSupply         *big.Int
}, error) {
	return _Reserve.Contract.Proofs(&_Reserve.CallOpts, arg0)
}

// Proofs is a free data retrieval call binding the contract method 0x9ddaf5aa.
//
// Solidity: function proofs(uint256 ) view returns(bytes32 signatureKeccakHash, bytes message, bytes reserveAddress, uint256 moneroBlockHeight, uint256 totalSupply)
func (_Reserve *ReserveCallerSession) Proofs(arg0 *big.Int) (struct {
	SignatureKeccakHash [32]byte
	Message             []byte
	ReserveAddress      []byte
	MoneroBlockHeight   *big.Int
	TotalSupply         *big.Int
}, error) {
	return _Reserve.Contract.Proofs(&_Reserve.CallOpts, arg0)
}

// ReserveAddress is a free data retrieval call binding the contract method 0xf79ed94b.
//
// Solidity: function reserveAddress() view returns(bytes)
func (_Reserve *ReserveCaller) ReserveAddress(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "reserveAddress")
	return *ret0, err
}

// ReserveAddress is a free data retrieval call binding the contract method 0xf79ed94b.
//
// Solidity: function reserveAddress() view returns(bytes)
func (_Reserve *ReserveSession) ReserveAddress() ([]byte, error) {
	return _Reserve.Contract.ReserveAddress(&_Reserve.CallOpts)
}

// ReserveAddress is a free data retrieval call binding the contract method 0xf79ed94b.
//
// Solidity: function reserveAddress() view returns(bytes)
func (_Reserve *ReserveCallerSession) ReserveAddress() ([]byte, error) {
	return _Reserve.Contract.ReserveAddress(&_Reserve.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Reserve *ReserveCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Reserve *ReserveSession) Symbol() (string, error) {
	return _Reserve.Contract.Symbol(&_Reserve.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Reserve *ReserveCallerSession) Symbol() (string, error) {
	return _Reserve.Contract.Symbol(&_Reserve.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Reserve *ReserveCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Reserve *ReserveSession) TotalSupply() (*big.Int, error) {
	return _Reserve.Contract.TotalSupply(&_Reserve.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Reserve *ReserveCallerSession) TotalSupply() (*big.Int, error) {
	return _Reserve.Contract.TotalSupply(&_Reserve.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Reserve *ReserveTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Reserve *ReserveSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Approve(&_Reserve.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Reserve *ReserveTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Approve(&_Reserve.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0xbcf64e05.
//
// Solidity: function burn(uint256 _amount, bytes32 _moneroAddressHash) returns(bool)
func (_Reserve *ReserveTransactor) Burn(opts *bind.TransactOpts, _amount *big.Int, _moneroAddressHash [32]byte) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "burn", _amount, _moneroAddressHash)
}

// Burn is a paid mutator transaction binding the contract method 0xbcf64e05.
//
// Solidity: function burn(uint256 _amount, bytes32 _moneroAddressHash) returns(bool)
func (_Reserve *ReserveSession) Burn(_amount *big.Int, _moneroAddressHash [32]byte) (*types.Transaction, error) {
	return _Reserve.Contract.Burn(&_Reserve.TransactOpts, _amount, _moneroAddressHash)
}

// Burn is a paid mutator transaction binding the contract method 0xbcf64e05.
//
// Solidity: function burn(uint256 _amount, bytes32 _moneroAddressHash) returns(bool)
func (_Reserve *ReserveTransactorSession) Burn(_amount *big.Int, _moneroAddressHash [32]byte) (*types.Transaction, error) {
	return _Reserve.Contract.Burn(&_Reserve.TransactOpts, _amount, _moneroAddressHash)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_Reserve *ReserveTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_Reserve *ReserveSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.DecreaseApproval(&_Reserve.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_Reserve *ReserveTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.DecreaseApproval(&_Reserve.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_Reserve *ReserveTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_Reserve *ReserveSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.IncreaseApproval(&_Reserve.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_Reserve *ReserveTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.IncreaseApproval(&_Reserve.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _recipient, uint256 _amount) returns(bool)
func (_Reserve *ReserveTransactor) Mint(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "mint", _recipient, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _recipient, uint256 _amount) returns(bool)
func (_Reserve *ReserveSession) Mint(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Mint(&_Reserve.TransactOpts, _recipient, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _recipient, uint256 _amount) returns(bool)
func (_Reserve *ReserveTransactorSession) Mint(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Mint(&_Reserve.TransactOpts, _recipient, _amount)
}

// PostReserveProof is a paid mutator transaction binding the contract method 0x8522da07.
//
// Solidity: function postReserveProof(bytes _message, bytes _reserveAddress, bytes32 _signatureKeccakHash, uint256 _moneroBlockHeight) returns(bool)
func (_Reserve *ReserveTransactor) PostReserveProof(opts *bind.TransactOpts, _message []byte, _reserveAddress []byte, _signatureKeccakHash [32]byte, _moneroBlockHeight *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "postReserveProof", _message, _reserveAddress, _signatureKeccakHash, _moneroBlockHeight)
}

// PostReserveProof is a paid mutator transaction binding the contract method 0x8522da07.
//
// Solidity: function postReserveProof(bytes _message, bytes _reserveAddress, bytes32 _signatureKeccakHash, uint256 _moneroBlockHeight) returns(bool)
func (_Reserve *ReserveSession) PostReserveProof(_message []byte, _reserveAddress []byte, _signatureKeccakHash [32]byte, _moneroBlockHeight *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.PostReserveProof(&_Reserve.TransactOpts, _message, _reserveAddress, _signatureKeccakHash, _moneroBlockHeight)
}

// PostReserveProof is a paid mutator transaction binding the contract method 0x8522da07.
//
// Solidity: function postReserveProof(bytes _message, bytes _reserveAddress, bytes32 _signatureKeccakHash, uint256 _moneroBlockHeight) returns(bool)
func (_Reserve *ReserveTransactorSession) PostReserveProof(_message []byte, _reserveAddress []byte, _signatureKeccakHash [32]byte, _moneroBlockHeight *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.PostReserveProof(&_Reserve.TransactOpts, _message, _reserveAddress, _signatureKeccakHash, _moneroBlockHeight)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Reserve *ReserveTransactor) SetAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "setAdmin", _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Reserve *ReserveSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.SetAdmin(&_Reserve.TransactOpts, _newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _newAdmin) returns(bool)
func (_Reserve *ReserveTransactorSession) SetAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.SetAdmin(&_Reserve.TransactOpts, _newAdmin)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 _rate) returns(bool)
func (_Reserve *ReserveTransactor) SetExchangeRate(opts *bind.TransactOpts, _rate *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "setExchangeRate", _rate)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 _rate) returns(bool)
func (_Reserve *ReserveSession) SetExchangeRate(_rate *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetExchangeRate(&_Reserve.TransactOpts, _rate)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 _rate) returns(bool)
func (_Reserve *ReserveTransactorSession) SetExchangeRate(_rate *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetExchangeRate(&_Reserve.TransactOpts, _rate)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Reserve *ReserveTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Reserve *ReserveSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Transfer(&_Reserve.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Reserve *ReserveTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Transfer(&_Reserve.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _recipient, uint256 _amount) returns(bool)
func (_Reserve *ReserveTransactor) TransferFrom(opts *bind.TransactOpts, _owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "transferFrom", _owner, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _recipient, uint256 _amount) returns(bool)
func (_Reserve *ReserveSession) TransferFrom(_owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.TransferFrom(&_Reserve.TransactOpts, _owner, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _recipient, uint256 _amount) returns(bool)
func (_Reserve *ReserveTransactorSession) TransferFrom(_owner common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.TransferFrom(&_Reserve.TransactOpts, _owner, _recipient, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Reserve *ReserveTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Reserve *ReserveSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.TransferOwnership(&_Reserve.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_Reserve *ReserveTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.TransferOwnership(&_Reserve.TransactOpts, _newOwner)
}

// ReserveAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the Reserve contract.
type ReserveAdminSetIterator struct {
	Event *ReserveAdminSet // Event containing the contract specifics and raw log

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
func (it *ReserveAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveAdminSet)
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
		it.Event = new(ReserveAdminSet)
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
func (it *ReserveAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveAdminSet represents a AdminSet event raised by the Reserve contract.
type ReserveAdminSet struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(address _admin)
func (_Reserve *ReserveFilterer) FilterAdminSet(opts *bind.FilterOpts) (*ReserveAdminSetIterator, error) {

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &ReserveAdminSetIterator{contract: _Reserve.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c.
//
// Solidity: event AdminSet(address _admin)
func (_Reserve *ReserveFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *ReserveAdminSet) (event.Subscription, error) {

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveAdminSet)
				if err := _Reserve.contract.UnpackLog(event, "AdminSet", log); err != nil {
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
func (_Reserve *ReserveFilterer) ParseAdminSet(log types.Log) (*ReserveAdminSet, error) {
	event := new(ReserveAdminSet)
	if err := _Reserve.contract.UnpackLog(event, "AdminSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReserveApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Reserve contract.
type ReserveApprovalIterator struct {
	Event *ReserveApproval // Event containing the contract specifics and raw log

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
func (it *ReserveApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveApproval)
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
		it.Event = new(ReserveApproval)
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
func (it *ReserveApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveApproval represents a Approval event raised by the Reserve contract.
type ReserveApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _amount)
func (_Reserve *ReserveFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ReserveApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ReserveApprovalIterator{contract: _Reserve.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _amount)
func (_Reserve *ReserveFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ReserveApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveApproval)
				if err := _Reserve.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Reserve *ReserveFilterer) ParseApproval(log types.Log) (*ReserveApproval, error) {
	event := new(ReserveApproval)
	if err := _Reserve.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReserveCoinsBurnedIterator is returned from FilterCoinsBurned and is used to iterate over the raw logs and unpacked data for CoinsBurned events raised by the Reserve contract.
type ReserveCoinsBurnedIterator struct {
	Event *ReserveCoinsBurned // Event containing the contract specifics and raw log

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
func (it *ReserveCoinsBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveCoinsBurned)
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
		it.Event = new(ReserveCoinsBurned)
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
func (it *ReserveCoinsBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveCoinsBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveCoinsBurned represents a CoinsBurned event raised by the Reserve contract.
type ReserveCoinsBurned struct {
	Burner            common.Address
	BurnAmount        *big.Int
	MoneroAddressHash [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCoinsBurned is a free log retrieval operation binding the contract event 0xbfe5be076e85df6616ae2ba97cc602211875bbf1c4458fd10012104b899d8f90.
//
// Solidity: event CoinsBurned(address _burner, uint256 _burnAmount, bytes32 _moneroAddressHash)
func (_Reserve *ReserveFilterer) FilterCoinsBurned(opts *bind.FilterOpts) (*ReserveCoinsBurnedIterator, error) {

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "CoinsBurned")
	if err != nil {
		return nil, err
	}
	return &ReserveCoinsBurnedIterator{contract: _Reserve.contract, event: "CoinsBurned", logs: logs, sub: sub}, nil
}

// WatchCoinsBurned is a free log subscription operation binding the contract event 0xbfe5be076e85df6616ae2ba97cc602211875bbf1c4458fd10012104b899d8f90.
//
// Solidity: event CoinsBurned(address _burner, uint256 _burnAmount, bytes32 _moneroAddressHash)
func (_Reserve *ReserveFilterer) WatchCoinsBurned(opts *bind.WatchOpts, sink chan<- *ReserveCoinsBurned) (event.Subscription, error) {

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "CoinsBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveCoinsBurned)
				if err := _Reserve.contract.UnpackLog(event, "CoinsBurned", log); err != nil {
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
// Solidity: event CoinsBurned(address _burner, uint256 _burnAmount, bytes32 _moneroAddressHash)
func (_Reserve *ReserveFilterer) ParseCoinsBurned(log types.Log) (*ReserveCoinsBurned, error) {
	event := new(ReserveCoinsBurned)
	if err := _Reserve.contract.UnpackLog(event, "CoinsBurned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReserveCoinsMintedIterator is returned from FilterCoinsMinted and is used to iterate over the raw logs and unpacked data for CoinsMinted events raised by the Reserve contract.
type ReserveCoinsMintedIterator struct {
	Event *ReserveCoinsMinted // Event containing the contract specifics and raw log

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
func (it *ReserveCoinsMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveCoinsMinted)
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
		it.Event = new(ReserveCoinsMinted)
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
func (it *ReserveCoinsMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveCoinsMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveCoinsMinted represents a CoinsMinted event raised by the Reserve contract.
type ReserveCoinsMinted struct {
	Recipient  common.Address
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCoinsMinted is a free log retrieval operation binding the contract event 0xc9ac6f335722ba5012b311e37c325450961c6e224ea391135b78d95118e190c9.
//
// Solidity: event CoinsMinted(address indexed _recipient, uint256 _mintAmount)
func (_Reserve *ReserveFilterer) FilterCoinsMinted(opts *bind.FilterOpts, _recipient []common.Address) (*ReserveCoinsMintedIterator, error) {

	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "CoinsMinted", _recipientRule)
	if err != nil {
		return nil, err
	}
	return &ReserveCoinsMintedIterator{contract: _Reserve.contract, event: "CoinsMinted", logs: logs, sub: sub}, nil
}

// WatchCoinsMinted is a free log subscription operation binding the contract event 0xc9ac6f335722ba5012b311e37c325450961c6e224ea391135b78d95118e190c9.
//
// Solidity: event CoinsMinted(address indexed _recipient, uint256 _mintAmount)
func (_Reserve *ReserveFilterer) WatchCoinsMinted(opts *bind.WatchOpts, sink chan<- *ReserveCoinsMinted, _recipient []common.Address) (event.Subscription, error) {

	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "CoinsMinted", _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveCoinsMinted)
				if err := _Reserve.contract.UnpackLog(event, "CoinsMinted", log); err != nil {
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
func (_Reserve *ReserveFilterer) ParseCoinsMinted(log types.Log) (*ReserveCoinsMinted, error) {
	event := new(ReserveCoinsMinted)
	if err := _Reserve.contract.UnpackLog(event, "CoinsMinted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReserveOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Reserve contract.
type ReserveOwnershipTransferredIterator struct {
	Event *ReserveOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReserveOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveOwnershipTransferred)
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
		it.Event = new(ReserveOwnershipTransferred)
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
func (it *ReserveOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveOwnershipTransferred represents a OwnershipTransferred event raised by the Reserve contract.
type ReserveOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address _previousOwner, address _newOwner)
func (_Reserve *ReserveFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*ReserveOwnershipTransferredIterator, error) {

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &ReserveOwnershipTransferredIterator{contract: _Reserve.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address _previousOwner, address _newOwner)
func (_Reserve *ReserveFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReserveOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveOwnershipTransferred)
				if err := _Reserve.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Reserve *ReserveFilterer) ParseOwnershipTransferred(log types.Log) (*ReserveOwnershipTransferred, error) {
	event := new(ReserveOwnershipTransferred)
	if err := _Reserve.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReserveProofSubmittedIterator is returned from FilterProofSubmitted and is used to iterate over the raw logs and unpacked data for ProofSubmitted events raised by the Reserve contract.
type ReserveProofSubmittedIterator struct {
	Event *ReserveProofSubmitted // Event containing the contract specifics and raw log

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
func (it *ReserveProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveProofSubmitted)
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
		it.Event = new(ReserveProofSubmitted)
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
func (it *ReserveProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveProofSubmitted represents a ProofSubmitted event raised by the Reserve contract.
type ReserveProofSubmitted struct {
	SignatureKeccakHash [32]byte
	BlockNumber         *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterProofSubmitted is a free log retrieval operation binding the contract event 0x3c0e85acf452e867ee09c9151d2cf45fe6d6a729271347e89a9a6b91b081f31d.
//
// Solidity: event ProofSubmitted(bytes32 _signatureKeccakHash, uint256 _blockNumber)
func (_Reserve *ReserveFilterer) FilterProofSubmitted(opts *bind.FilterOpts) (*ReserveProofSubmittedIterator, error) {

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "ProofSubmitted")
	if err != nil {
		return nil, err
	}
	return &ReserveProofSubmittedIterator{contract: _Reserve.contract, event: "ProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchProofSubmitted is a free log subscription operation binding the contract event 0x3c0e85acf452e867ee09c9151d2cf45fe6d6a729271347e89a9a6b91b081f31d.
//
// Solidity: event ProofSubmitted(bytes32 _signatureKeccakHash, uint256 _blockNumber)
func (_Reserve *ReserveFilterer) WatchProofSubmitted(opts *bind.WatchOpts, sink chan<- *ReserveProofSubmitted) (event.Subscription, error) {

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "ProofSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveProofSubmitted)
				if err := _Reserve.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
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

// ParseProofSubmitted is a log parse operation binding the contract event 0x3c0e85acf452e867ee09c9151d2cf45fe6d6a729271347e89a9a6b91b081f31d.
//
// Solidity: event ProofSubmitted(bytes32 _signatureKeccakHash, uint256 _blockNumber)
func (_Reserve *ReserveFilterer) ParseProofSubmitted(log types.Log) (*ReserveProofSubmitted, error) {
	event := new(ReserveProofSubmitted)
	if err := _Reserve.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReserveTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Reserve contract.
type ReserveTransferIterator struct {
	Event *ReserveTransfer // Event containing the contract specifics and raw log

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
func (it *ReserveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveTransfer)
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
		it.Event = new(ReserveTransfer)
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
func (it *ReserveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveTransfer represents a Transfer event raised by the Reserve contract.
type ReserveTransfer struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _sender, address indexed _recipient, uint256 _amount)
func (_Reserve *ReserveFilterer) FilterTransfer(opts *bind.FilterOpts, _sender []common.Address, _recipient []common.Address) (*ReserveTransferIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "Transfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return &ReserveTransferIterator{contract: _Reserve.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _sender, address indexed _recipient, uint256 _amount)
func (_Reserve *ReserveFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ReserveTransfer, _sender []common.Address, _recipient []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "Transfer", _senderRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveTransfer)
				if err := _Reserve.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Reserve *ReserveFilterer) ParseTransfer(log types.Log) (*ReserveTransfer, error) {
	event := new(ReserveTransfer)
	if err := _Reserve.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
