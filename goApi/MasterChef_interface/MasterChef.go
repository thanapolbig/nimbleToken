// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MasterChef_interface

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

// MasterChefEvent is an auto generated low-level Go binding around an user-defined struct.
type MasterChefEvent struct {
	CreateBy  common.Address
	NameEvent string
	Detial    string
	Reward    *big.Int
	Status    *big.Int
	TimeStart *big.Int
}

// ApiABI is the input ABI used to generate the binding from.
const ApiABI = "[{\"inputs\":[{\"internalType\":\"contractNimbleToken\",\"name\":\"_nimble\",\"type\":\"address\"},{\"internalType\":\"contractSyrupBar\",\"name\":\"_syrup\",\"type\":\"address\"},{\"internalType\":\"contractEventBar\",\"name\":\"_event\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"name\":\"LogVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_createBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_nameEvent\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_detial\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timeStart\",\"type\":\"uint256\"}],\"name\":\"_createEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eid\",\"type\":\"uint256\"}],\"name\":\"_join\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"listWallet\",\"type\":\"address[]\"}],\"name\":\"getScoreVote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eid\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"wallet\",\"type\":\"address[]\"}],\"name\":\"AcceptEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eid\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"wallet\",\"type\":\"address[]\"}],\"name\":\"AcceptEventAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"listWallet\",\"type\":\"address[]\"}],\"name\":\"addVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"addWorkday\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"account\",\"type\":\"address[]\"}],\"name\":\"autoClaimCheckin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"account\",\"type\":\"address[]\"}],\"name\":\"claimRewardScoreVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eid\",\"type\":\"uint256\"}],\"name\":\"closeEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"configMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nameEvent\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_detial\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeStart\",\"type\":\"uint256\"}],\"name\":\"createEvent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nameEvent\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_detial\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeStart\",\"type\":\"uint256\"}],\"name\":\"createEventAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daysCheckin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"eventInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"createBy\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nameEvent\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"detial\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStart\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eventbar\",\"outputs\":[{\"internalType\":\"contractEventBar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"getRightScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"getScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eid\",\"type\":\"uint256\"}],\"name\":\"joinEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nimble\",\"outputs\":[{\"internalType\":\"contractNimbleToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"searchEvent\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"createBy\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nameEvent\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"detial\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStart\",\"type\":\"uint256\"}],\"internalType\":\"structMasterChef.Event[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"searchEventAdmin\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"searchEventByAddress\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eid\",\"type\":\"uint256\"}],\"name\":\"searchParticipant\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eid\",\"type\":\"uint256\"}],\"name\":\"startEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syrup\",\"outputs\":[{\"internalType\":\"contractSyrupBar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"workday\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ApiBin is the compiled bytecode used for deploying new contracts.
var ApiBin = "0x608060405260006005556302faf080600855601e600955601e600a553480156200002857600080fd5b50604051620021d8380380620021d88339810160408190526200004b91620000eb565b62000056336200009b565b600280546001600160a01b03199081166001600160a01b0394851617909155600180548216948416949094179093556003805490931691161790554260045562000158565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000806000606084860312156200010157600080fd5b83516200010e816200013f565b602085015190935062000121816200013f565b604085015190925062000134816200013f565b809150509250925092565b6001600160a01b03811681146200015557600080fd5b50565b61207080620001686000396000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c8063715018a611610104578063a59a6953116100a2578063bce062b911610071578063bce062b9146103ed578063d47875d014610400578063f2fde38b1461042c578063f6aef40c1461043f57600080fd5b8063a59a69531461039e578063a637aed4146103b1578063adaee4ab146103d1578063ae0849b4146103da57600080fd5b80638da5cb5b116100de5780638da5cb5b1461034c57806399c944791461035d5780639d0c025b146103705780639f301cac1461039557600080fd5b8063715018a614610329578063781c07901461033157806386a952c41461033957600080fd5b806332ace39c1161017157806349eacceb1161014b57806349eacceb146102d05780634f7c47b0146102e35780635f74bbde146102f65780636dc22b241461030957600080fd5b806332ace39c1461027357806336d9d7461461029457806338d361e2146102bd57600080fd5b806327cacc19116101ad57806327cacc191461020f5780632e04349f146102225780632e9b017b146102355780632ee07c001461026057600080fd5b80631249c58b146101d457806316ea468f146101de57806318ec0147146101fc575b600080fd5b6101dc610452565b005b6101e661055e565b6040516101f39190611d8b565b60405180910390f35b6101dc61020a366004611c0a565b610715565b6101dc61021d366004611b05565b610853565b6101dc610230366004611b05565b610912565b600154610248906001600160a01b031681565b6040516001600160a01b0390911681526020016101f3565b6101dc61026e366004611bd8565b610a73565b610286610281366004611b64565b610b38565b6040519081526020016101f3565b6102866102a2366004611ab9565b6001600160a01b03166000908152600b602052604090205490565b6101dc6102cb366004611bd8565b610d7c565b600354610248906001600160a01b031681565b6101dc6102f1366004611c0a565b610eb3565b6101dc610304366004611adb565b610feb565b61031c610317366004611ab9565b6110e8565b6040516101f39190611e45565b6101dc611154565b61031c61118a565b600254610248906001600160a01b031681565b6000546001600160a01b0316610248565b61028661036b366004611bd8565b611229565b61038361037e366004611bd8565b611261565b6040516101f396959493929190611ce9565b61028660065481565b6101dc6103ac366004611bd8565b6113c3565b6103c46103bf366004611bd8565b6114d8565b6040516101f39190611d3e565b61028660075481565b6101dc6103e8366004611b05565b611543565b6102866103fb366004611b64565b61164c565b61028661040e366004611ab9565b6001600160a01b03166000908152600b602052604090206001015490565b6101dc61043a366004611ab9565b611718565b61028661044d366004611bd8565b6117b3565b60006005546301e133806104669190611f6c565b6004546104739190611f32565b9050804210156104b95760405162461bcd60e51b815260206004820152600c60248201526b1b9bdd081d1a5b59481e595d60a21b60448201526064015b60405180910390fd5b6001546002546008546040516340c10f1960e01b81526001600160a01b03928316600482015260248101919091529116906340c10f1990604401602060405180830381600087803b15801561050d57600080fd5b505af1158015610521573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105459190611b42565b506005805490600061055683611fdd565b919050555050565b6060600c805480602002602001604051908101604052809291908181526020016000905b8282101561070c5760008481526020908190206040805160c081019091526006850290910180546001600160a01b0316825260018101805492939192918401916105cb90611fa2565b80601f01602080910402602001604051908101604052809291908181526020018280546105f790611fa2565b80156106445780601f1061061957610100808354040283529160200191610644565b820191906000526020600020905b81548152906001019060200180831161062757829003601f168201915b5050505050815260200160028201805461065d90611fa2565b80601f016020809104026020016040519081016040528092919081815260200182805461068990611fa2565b80156106d65780601f106106ab576101008083540402835291602001916106d6565b820191906000526020600020905b8154815290600101906020018083116106b957829003601f168201915b50505050508152602001600382015481526020016004820154815260200160058201548152505081526020019060010190610582565b50505050905090565b336001600160a01b0316600c83815481106107325761073261200e565b60009182526020909120600690910201546001600160a01b0316146107695760405162461bcd60e51b81526004016104b090611eda565b600c828154811061077c5761077c61200e565b9060005260206000209060060201600401546001146107ad5760405162461bcd60e51b81526004016104b090611e7d565b6003600c83815481106107c2576107c261200e565b90600052602060002090600602016004018190555060008151600c84815481106107ee576107ee61200e565b90600052602060002090600602016003015461080a9190611f4a565b905060005b825181101561084d5761083b83828151811061082d5761082d61200e565b6020026020010151836117e7565b8061084581611fdd565b91505061080f565b50505050565b6000546001600160a01b0316331461087d5760405162461bcd60e51b81526004016104b090611ea5565b60005b81518110156108d7576004600b60008484815181106108a1576108a161200e565b6020908102919091018101516001600160a01b0316825281019190915260400160002055806108cf81611fdd565b915050610880565b507fb376534b1fb0889c24ac054e5b06cbc118d462433346a80ced64af664d889df6816040516109079190611d3e565b60405180910390a150565b6000546001600160a01b0316331461093c5760405162461bcd60e51b81526004016104b090611ea5565b6000805b82518110156109aa57600b600084838151811061095f5761095f61200e565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060010154826109969190611f32565b9150806109a281611fdd565b915050610940565b5060006064600a546008546109bf9190611f6c565b6109c99190611f4a565b905060006109d78383611f4a565b905060005b8451811015610a6c57600082600b60008885815181106109fe576109fe61200e565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060010154610a349190611f6c565b9050610a59868381518110610a4b57610a4b61200e565b60200260200101518261184a565b5080610a6481611fdd565b9150506109dc565b5050505050565b336001600160a01b0316600c8281548110610a9057610a9061200e565b60009182526020909120600690910201546001600160a01b031614610ac75760405162461bcd60e51b81526004016104b090611eda565b600c8181548110610ada57610ada61200e565b906000526020600020906006020160040154600114610b0b5760405162461bcd60e51b81526004016104b090611e7d565b6002600c8281548110610b2057610b2061200e565b90600052602060002090600602016004018190555050565b600080546001600160a01b03163314610b635760405162461bcd60e51b81526004016104b090611ea5565b6001546002546040516370a0823160e01b81526001600160a01b0391821660048201529116906370a082319060240160206040518083038186803b158015610baa57600080fd5b505afa158015610bbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610be29190611bf1565b831115610c285760405162461bcd60e51b81526020600482015260146024820152733cb7ba903237b713ba103430bb32903a37b5b2b760611b60448201526064016104b0565b6040805160c08101825233815260208082018881529282018790526060820186905260006080830181905260a08301869052600c8054600181018255915282517fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c7600690920291820180546001600160a01b0319166001600160a01b0390921691909117815593518051939493610ce8937fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c890930192919091019061190c565b5060408201518051610d0491600284019160209091019061190c565b50606082015181600301556080820151816004015560a0820151816005015550507f71f5db2c13452fa93276639f35267a3e4a4c979bf4f10bd87ea535b4d975cb6a3386868686604051610d5c959493929190611c9e565b60405180910390a1600c54610d7390600190611f8b565b95945050505050565b600c54610d8b90600190611f8b565b811115610dcb5760405162461bcd60e51b815260206004820152600e60248201526d115d995b9d081b9bdd081c99585b60921b60448201526064016104b0565b600c8181548110610dde57610dde61200e565b906000526020600020906006020160040154600114610e325760405162461bcd60e51b815260206004820152601060248201526f22bb32b73a1031b0b713ba103537b4b760811b60448201526064016104b0565b336000818152600e6020908152604080832080546001808201835591855283852001869055858452600d83528184208054918201815584529282902090920180546001600160a01b03191684179055815192835282018390527f2e55f4a4bcfc21d67be1e2e465ef0c4b6a9a3bce54eb1a9754dae16fc63828a29101610907565b336001600160a01b0316600c8381548110610ed057610ed061200e565b60009182526020909120600690910201546001600160a01b031614610f075760405162461bcd60e51b81526004016104b090611eda565b600c8281548110610f1a57610f1a61200e565b906000526020600020906006020160040154600114610f4b5760405162461bcd60e51b81526004016104b090611e7d565b6003600c8381548110610f6057610f6061200e565b90600052602060002090600602016004018190555060008151600c8481548110610f8c57610f8c61200e565b906000526020600020906006020160030154610fa89190611f4a565b905060005b825181101561084d57610fd9838281518110610fcb57610fcb61200e565b60200260200101518361184a565b80610fe381611fdd565b915050610fad565b336000908152600b602052604090205481111561103d5760405162461bcd60e51b815260206004820152601060248201526f6e6f7420656e6f7567682073636f726560801b60448201526064016104b0565b336000908152600b6020526040902054611058908290611f8b565b336000908152600b6020526040808220929092556001600160a01b03841681522060010154611088908290611f32565b6001600160a01b0383166000818152600b60205260409081902060010192909255905133907f49ce5cb7b86410ac7069ff893207f2804cf4614b4203eaf4e0e37bb41a2b0ef0906110dc9085815260200190565b60405180910390a35050565b6001600160a01b0381166000908152600e602090815260409182902080548351818402810184019094528084526060939283018282801561114857602002820191906000526020600020905b815481526020019060010190808311611134575b50505050509050919050565b6000546001600160a01b0316331461117e5760405162461bcd60e51b81526004016104b090611ea5565b6111886000611883565b565b60606000606060005b600c548111611222576000546001600160a01b03166001600160a01b0316600c82815481106111c4576111c461200e565b60009182526020909120600690910201546001600160a01b0316141561121057808284815181106111f7576111f761200e565b60209081029190910101528261120c81611fdd565b9350505b8061121a81611fdd565b915050611193565b5092915050565b600080546001600160a01b031633146112545760405162461bcd60e51b81526004016104b090611ea5565b506007819055805b919050565b600c818154811061127157600080fd5b6000918252602090912060069091020180546001820180546001600160a01b039092169350906112a090611fa2565b80601f01602080910402602001604051908101604052809291908181526020018280546112cc90611fa2565b80156113195780601f106112ee57610100808354040283529160200191611319565b820191906000526020600020905b8154815290600101906020018083116112fc57829003601f168201915b50505050509080600201805461132e90611fa2565b80601f016020809104026020016040519081016040528092919081815260200182805461135a90611fa2565b80156113a75780601f1061137c576101008083540402835291602001916113a7565b820191906000526020600020905b81548152906001019060200180831161138a57829003601f168201915b5050505050908060030154908060040154908060050154905086565b336001600160a01b0316600c82815481106113e0576113e061200e565b60009182526020909120600690910201546001600160a01b0316146114175760405162461bcd60e51b81526004016104b090611eda565b600c818154811061142a5761142a61200e565b90600052602060002090600602016004015460001461145b5760405162461bcd60e51b81526004016104b090611e7d565b42600c828154811061146f5761146f61200e565b90600052602060002090600602016005015411156114c35760405162461bcd60e51b8152602060048201526011602482015270125d09dcc81b9bdd081d1a5b59481e595d607a1b60448201526064016104b0565b6001600c8281548110610b2057610b2061200e565b6000818152600d602090815260409182902080548351818402810184019094528084526060939283018282801561114857602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161151a5750505050509050919050565b6000546001600160a01b0316331461156d5760405162461bcd60e51b81526004016104b090611ea5565b6000600654620151806115809190611f6c565b60045461158d9190611f32565b9050804210156115ce5760405162461bcd60e51b815260206004820152600c60248201526b1b9bdd081d1a5b59481e595d60a21b60448201526064016104b0565b600060646009546008546115e29190611f6c565b6115ec9190611f4a565b90506000600754826115fe9190611f4a565b9050600084518261160f9190611f4a565b905060005b855181101561164457611632868281518110610fcb57610fcb61200e565b8061163c81611fdd565b915050611614565b505050505050565b6001546040516370a0823160e01b81523360048201526000916001600160a01b0316906370a082319060240160206040518083038186803b15801561169057600080fd5b505afa1580156116a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116c89190611bf1565b83111561170e5760405162461bcd60e51b81526020600482015260146024820152733cb7ba903237b713ba103430bb32903a37b5b2b760611b60448201526064016104b0565b610c2833846118d3565b6000546001600160a01b031633146117425760405162461bcd60e51b81526004016104b090611ea5565b6001600160a01b0381166117a75760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016104b0565b6117b081611883565b50565b600080546001600160a01b031633146117de5760405162461bcd60e51b81526004016104b090611ea5565b50600881905590565b600354604051632683c1e360e11b81526001600160a01b0384811660048301526024820184905290911690634d0783c6906044015b600060405180830381600087803b15801561183657600080fd5b505af1158015611644573d6000803e3d6000fd5b600254604051632683c1e360e11b81526001600160a01b0384811660048301526024820184905290911690634d0783c69060440161181c565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60035460405163551026bb60e01b81526001600160a01b038481166004830152602482018490529091169063551026bb9060440161181c565b82805461191890611fa2565b90600052602060002090601f01602090048101928261193a5760008555611980565b82601f1061195357805160ff1916838001178555611980565b82800160010185558215611980579182015b82811115611980578251825591602001919060010190611965565b5061198c929150611990565b5090565b5b8082111561198c5760008155600101611991565b80356001600160a01b038116811461125c57600080fd5b600082601f8301126119cd57600080fd5b8135602067ffffffffffffffff8211156119e9576119e9612024565b8160051b6119f8828201611f01565b838152828101908684018388018501891015611a1357600080fd5b600093505b85841015611a3d57611a29816119a5565b835260019390930192918401918401611a18565b50979650505050505050565b600082601f830112611a5a57600080fd5b813567ffffffffffffffff811115611a7457611a74612024565b611a87601f8201601f1916602001611f01565b818152846020838601011115611a9c57600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215611acb57600080fd5b611ad4826119a5565b9392505050565b60008060408385031215611aee57600080fd5b611af7836119a5565b946020939093013593505050565b600060208284031215611b1757600080fd5b813567ffffffffffffffff811115611b2e57600080fd5b611b3a848285016119bc565b949350505050565b600060208284031215611b5457600080fd5b81518015158114611ad457600080fd5b60008060008060808587031215611b7a57600080fd5b843567ffffffffffffffff80821115611b9257600080fd5b611b9e88838901611a49565b95506020870135915080821115611bb457600080fd5b50611bc187828801611a49565b949794965050505060408301359260600135919050565b600060208284031215611bea57600080fd5b5035919050565b600060208284031215611c0357600080fd5b5051919050565b60008060408385031215611c1d57600080fd5b82359150602083013567ffffffffffffffff811115611c3b57600080fd5b611c47858286016119bc565b9150509250929050565b6000815180845260005b81811015611c7757602081850181015186830182015201611c5b565b81811115611c89576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b038616815260a060208201819052600090611cc290830187611c51565b8281036040840152611cd48187611c51565b60608401959095525050608001529392505050565b6001600160a01b038716815260c060208201819052600090611d0d90830188611c51565b8281036040840152611d1f8188611c51565b60608401969096525050608081019290925260a0909101529392505050565b6020808252825182820181905260009190848201906040850190845b81811015611d7f5783516001600160a01b031683529284019291840191600101611d5a565b50909695505050505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015611e3757888303603f19018552815180516001600160a01b031684528781015160c089860181905290611deb82870182611c51565b9150508782015185820389870152611e038282611c51565b606084810151908801526080808501519088015260a093840151939096019290925250509386019390860190600101611db2565b509098975050505050505050565b6020808252825182820181905260009190848201906040850190845b81811015611d7f57835183529284019291840191600101611e61565b6020808252600e908201526d1cdd185d1d5cc81a5b9d985b1a5960921b604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252600d908201526c1858d8d95cdcc819195b9a5959609a1b604082015260600190565b604051601f8201601f1916810167ffffffffffffffff81118282101715611f2a57611f2a612024565b604052919050565b60008219821115611f4557611f45611ff8565b500190565b600082611f6757634e487b7160e01b600052601260045260246000fd5b500490565b6000816000190483118215151615611f8657611f86611ff8565b500290565b600082821015611f9d57611f9d611ff8565b500390565b600181811c90821680611fb657607f821691505b60208210811415611fd757634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415611ff157611ff1611ff8565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212208745042ce2566e46f345cdbf15f5e28701a76e41b26b769023f8dd19b4bce9c964736f6c63430008060033"

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _nimble common.Address, _syrup common.Address, _event common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApiBin), backend, _nimble, _syrup, _event)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// DaysCheckin is a free data retrieval call binding the contract method 0x9f301cac.
//
// Solidity: function daysCheckin() view returns(uint256)
func (_Api *ApiCaller) DaysCheckin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "daysCheckin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DaysCheckin is a free data retrieval call binding the contract method 0x9f301cac.
//
// Solidity: function daysCheckin() view returns(uint256)
func (_Api *ApiSession) DaysCheckin() (*big.Int, error) {
	return _Api.Contract.DaysCheckin(&_Api.CallOpts)
}

// DaysCheckin is a free data retrieval call binding the contract method 0x9f301cac.
//
// Solidity: function daysCheckin() view returns(uint256)
func (_Api *ApiCallerSession) DaysCheckin() (*big.Int, error) {
	return _Api.Contract.DaysCheckin(&_Api.CallOpts)
}

// EventInfo is a free data retrieval call binding the contract method 0x9d0c025b.
//
// Solidity: function eventInfo(uint256 ) view returns(address createBy, string nameEvent, string detial, uint256 reward, uint256 status, uint256 timeStart)
func (_Api *ApiCaller) EventInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CreateBy  common.Address
	NameEvent string
	Detial    string
	Reward    *big.Int
	Status    *big.Int
	TimeStart *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "eventInfo", arg0)

	outstruct := new(struct {
		CreateBy  common.Address
		NameEvent string
		Detial    string
		Reward    *big.Int
		Status    *big.Int
		TimeStart *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CreateBy = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NameEvent = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Detial = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Reward = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TimeStart = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EventInfo is a free data retrieval call binding the contract method 0x9d0c025b.
//
// Solidity: function eventInfo(uint256 ) view returns(address createBy, string nameEvent, string detial, uint256 reward, uint256 status, uint256 timeStart)
func (_Api *ApiSession) EventInfo(arg0 *big.Int) (struct {
	CreateBy  common.Address
	NameEvent string
	Detial    string
	Reward    *big.Int
	Status    *big.Int
	TimeStart *big.Int
}, error) {
	return _Api.Contract.EventInfo(&_Api.CallOpts, arg0)
}

// EventInfo is a free data retrieval call binding the contract method 0x9d0c025b.
//
// Solidity: function eventInfo(uint256 ) view returns(address createBy, string nameEvent, string detial, uint256 reward, uint256 status, uint256 timeStart)
func (_Api *ApiCallerSession) EventInfo(arg0 *big.Int) (struct {
	CreateBy  common.Address
	NameEvent string
	Detial    string
	Reward    *big.Int
	Status    *big.Int
	TimeStart *big.Int
}, error) {
	return _Api.Contract.EventInfo(&_Api.CallOpts, arg0)
}

// Eventbar is a free data retrieval call binding the contract method 0x49eacceb.
//
// Solidity: function eventbar() view returns(address)
func (_Api *ApiCaller) Eventbar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "eventbar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Eventbar is a free data retrieval call binding the contract method 0x49eacceb.
//
// Solidity: function eventbar() view returns(address)
func (_Api *ApiSession) Eventbar() (common.Address, error) {
	return _Api.Contract.Eventbar(&_Api.CallOpts)
}

// Eventbar is a free data retrieval call binding the contract method 0x49eacceb.
//
// Solidity: function eventbar() view returns(address)
func (_Api *ApiCallerSession) Eventbar() (common.Address, error) {
	return _Api.Contract.Eventbar(&_Api.CallOpts)
}

// GetRightScore is a free data retrieval call binding the contract method 0x36d9d746.
//
// Solidity: function getRightScore(address wallet) view returns(uint256)
func (_Api *ApiCaller) GetRightScore(opts *bind.CallOpts, wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getRightScore", wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRightScore is a free data retrieval call binding the contract method 0x36d9d746.
//
// Solidity: function getRightScore(address wallet) view returns(uint256)
func (_Api *ApiSession) GetRightScore(wallet common.Address) (*big.Int, error) {
	return _Api.Contract.GetRightScore(&_Api.CallOpts, wallet)
}

// GetRightScore is a free data retrieval call binding the contract method 0x36d9d746.
//
// Solidity: function getRightScore(address wallet) view returns(uint256)
func (_Api *ApiCallerSession) GetRightScore(wallet common.Address) (*big.Int, error) {
	return _Api.Contract.GetRightScore(&_Api.CallOpts, wallet)
}

// GetScore is a free data retrieval call binding the contract method 0xd47875d0.
//
// Solidity: function getScore(address wallet) view returns(uint256)
func (_Api *ApiCaller) GetScore(opts *bind.CallOpts, wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getScore", wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetScore is a free data retrieval call binding the contract method 0xd47875d0.
//
// Solidity: function getScore(address wallet) view returns(uint256)
func (_Api *ApiSession) GetScore(wallet common.Address) (*big.Int, error) {
	return _Api.Contract.GetScore(&_Api.CallOpts, wallet)
}

// GetScore is a free data retrieval call binding the contract method 0xd47875d0.
//
// Solidity: function getScore(address wallet) view returns(uint256)
func (_Api *ApiCallerSession) GetScore(wallet common.Address) (*big.Int, error) {
	return _Api.Contract.GetScore(&_Api.CallOpts, wallet)
}

// Nimble is a free data retrieval call binding the contract method 0x2e9b017b.
//
// Solidity: function nimble() view returns(address)
func (_Api *ApiCaller) Nimble(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "nimble")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nimble is a free data retrieval call binding the contract method 0x2e9b017b.
//
// Solidity: function nimble() view returns(address)
func (_Api *ApiSession) Nimble() (common.Address, error) {
	return _Api.Contract.Nimble(&_Api.CallOpts)
}

// Nimble is a free data retrieval call binding the contract method 0x2e9b017b.
//
// Solidity: function nimble() view returns(address)
func (_Api *ApiCallerSession) Nimble() (common.Address, error) {
	return _Api.Contract.Nimble(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCallerSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() view returns(address)
func (_Api *ApiCaller) Syrup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "syrup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() view returns(address)
func (_Api *ApiSession) Syrup() (common.Address, error) {
	return _Api.Contract.Syrup(&_Api.CallOpts)
}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() view returns(address)
func (_Api *ApiCallerSession) Syrup() (common.Address, error) {
	return _Api.Contract.Syrup(&_Api.CallOpts)
}

// Workday is a free data retrieval call binding the contract method 0xadaee4ab.
//
// Solidity: function workday() view returns(uint256)
func (_Api *ApiCaller) Workday(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "workday")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Workday is a free data retrieval call binding the contract method 0xadaee4ab.
//
// Solidity: function workday() view returns(uint256)
func (_Api *ApiSession) Workday() (*big.Int, error) {
	return _Api.Contract.Workday(&_Api.CallOpts)
}

// Workday is a free data retrieval call binding the contract method 0xadaee4ab.
//
// Solidity: function workday() view returns(uint256)
func (_Api *ApiCallerSession) Workday() (*big.Int, error) {
	return _Api.Contract.Workday(&_Api.CallOpts)
}

// AcceptEvent is a paid mutator transaction binding the contract method 0x18ec0147.
//
// Solidity: function AcceptEvent(uint256 eid, address[] wallet) returns()
func (_Api *ApiTransactor) AcceptEvent(opts *bind.TransactOpts, eid *big.Int, wallet []common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "AcceptEvent", eid, wallet)
}

// AcceptEvent is a paid mutator transaction binding the contract method 0x18ec0147.
//
// Solidity: function AcceptEvent(uint256 eid, address[] wallet) returns()
func (_Api *ApiSession) AcceptEvent(eid *big.Int, wallet []common.Address) (*types.Transaction, error) {
	return _Api.Contract.AcceptEvent(&_Api.TransactOpts, eid, wallet)
}

// AcceptEvent is a paid mutator transaction binding the contract method 0x18ec0147.
//
// Solidity: function AcceptEvent(uint256 eid, address[] wallet) returns()
func (_Api *ApiTransactorSession) AcceptEvent(eid *big.Int, wallet []common.Address) (*types.Transaction, error) {
	return _Api.Contract.AcceptEvent(&_Api.TransactOpts, eid, wallet)
}

// AcceptEventAdmin is a paid mutator transaction binding the contract method 0x4f7c47b0.
//
// Solidity: function AcceptEventAdmin(uint256 eid, address[] wallet) returns()
func (_Api *ApiTransactor) AcceptEventAdmin(opts *bind.TransactOpts, eid *big.Int, wallet []common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "AcceptEventAdmin", eid, wallet)
}

// AcceptEventAdmin is a paid mutator transaction binding the contract method 0x4f7c47b0.
//
// Solidity: function AcceptEventAdmin(uint256 eid, address[] wallet) returns()
func (_Api *ApiSession) AcceptEventAdmin(eid *big.Int, wallet []common.Address) (*types.Transaction, error) {
	return _Api.Contract.AcceptEventAdmin(&_Api.TransactOpts, eid, wallet)
}

// AcceptEventAdmin is a paid mutator transaction binding the contract method 0x4f7c47b0.
//
// Solidity: function AcceptEventAdmin(uint256 eid, address[] wallet) returns()
func (_Api *ApiTransactorSession) AcceptEventAdmin(eid *big.Int, wallet []common.Address) (*types.Transaction, error) {
	return _Api.Contract.AcceptEventAdmin(&_Api.TransactOpts, eid, wallet)
}

// AddVote is a paid mutator transaction binding the contract method 0x27cacc19.
//
// Solidity: function addVote(address[] listWallet) returns()
func (_Api *ApiTransactor) AddVote(opts *bind.TransactOpts, listWallet []common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addVote", listWallet)
}

// AddVote is a paid mutator transaction binding the contract method 0x27cacc19.
//
// Solidity: function addVote(address[] listWallet) returns()
func (_Api *ApiSession) AddVote(listWallet []common.Address) (*types.Transaction, error) {
	return _Api.Contract.AddVote(&_Api.TransactOpts, listWallet)
}

// AddVote is a paid mutator transaction binding the contract method 0x27cacc19.
//
// Solidity: function addVote(address[] listWallet) returns()
func (_Api *ApiTransactorSession) AddVote(listWallet []common.Address) (*types.Transaction, error) {
	return _Api.Contract.AddVote(&_Api.TransactOpts, listWallet)
}

// AddWorkday is a paid mutator transaction binding the contract method 0x99c94479.
//
// Solidity: function addWorkday(uint256 value) returns(uint256)
func (_Api *ApiTransactor) AddWorkday(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addWorkday", value)
}

// AddWorkday is a paid mutator transaction binding the contract method 0x99c94479.
//
// Solidity: function addWorkday(uint256 value) returns(uint256)
func (_Api *ApiSession) AddWorkday(value *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddWorkday(&_Api.TransactOpts, value)
}

// AddWorkday is a paid mutator transaction binding the contract method 0x99c94479.
//
// Solidity: function addWorkday(uint256 value) returns(uint256)
func (_Api *ApiTransactorSession) AddWorkday(value *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddWorkday(&_Api.TransactOpts, value)
}

// AutoClaimCheckin is a paid mutator transaction binding the contract method 0xae0849b4.
//
// Solidity: function autoClaimCheckin(address[] account) returns()
func (_Api *ApiTransactor) AutoClaimCheckin(opts *bind.TransactOpts, account []common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "autoClaimCheckin", account)
}

// AutoClaimCheckin is a paid mutator transaction binding the contract method 0xae0849b4.
//
// Solidity: function autoClaimCheckin(address[] account) returns()
func (_Api *ApiSession) AutoClaimCheckin(account []common.Address) (*types.Transaction, error) {
	return _Api.Contract.AutoClaimCheckin(&_Api.TransactOpts, account)
}

// AutoClaimCheckin is a paid mutator transaction binding the contract method 0xae0849b4.
//
// Solidity: function autoClaimCheckin(address[] account) returns()
func (_Api *ApiTransactorSession) AutoClaimCheckin(account []common.Address) (*types.Transaction, error) {
	return _Api.Contract.AutoClaimCheckin(&_Api.TransactOpts, account)
}

// ClaimRewardScoreVote is a paid mutator transaction binding the contract method 0x2e04349f.
//
// Solidity: function claimRewardScoreVote(address[] account) returns()
func (_Api *ApiTransactor) ClaimRewardScoreVote(opts *bind.TransactOpts, account []common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "claimRewardScoreVote", account)
}

// ClaimRewardScoreVote is a paid mutator transaction binding the contract method 0x2e04349f.
//
// Solidity: function claimRewardScoreVote(address[] account) returns()
func (_Api *ApiSession) ClaimRewardScoreVote(account []common.Address) (*types.Transaction, error) {
	return _Api.Contract.ClaimRewardScoreVote(&_Api.TransactOpts, account)
}

// ClaimRewardScoreVote is a paid mutator transaction binding the contract method 0x2e04349f.
//
// Solidity: function claimRewardScoreVote(address[] account) returns()
func (_Api *ApiTransactorSession) ClaimRewardScoreVote(account []common.Address) (*types.Transaction, error) {
	return _Api.Contract.ClaimRewardScoreVote(&_Api.TransactOpts, account)
}

// CloseEvent is a paid mutator transaction binding the contract method 0x2ee07c00.
//
// Solidity: function closeEvent(uint256 eid) returns()
func (_Api *ApiTransactor) CloseEvent(opts *bind.TransactOpts, eid *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "closeEvent", eid)
}

// CloseEvent is a paid mutator transaction binding the contract method 0x2ee07c00.
//
// Solidity: function closeEvent(uint256 eid) returns()
func (_Api *ApiSession) CloseEvent(eid *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CloseEvent(&_Api.TransactOpts, eid)
}

// CloseEvent is a paid mutator transaction binding the contract method 0x2ee07c00.
//
// Solidity: function closeEvent(uint256 eid) returns()
func (_Api *ApiTransactorSession) CloseEvent(eid *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CloseEvent(&_Api.TransactOpts, eid)
}

// ConfigMint is a paid mutator transaction binding the contract method 0xf6aef40c.
//
// Solidity: function configMint(uint256 value) returns(uint256)
func (_Api *ApiTransactor) ConfigMint(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "configMint", value)
}

// ConfigMint is a paid mutator transaction binding the contract method 0xf6aef40c.
//
// Solidity: function configMint(uint256 value) returns(uint256)
func (_Api *ApiSession) ConfigMint(value *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ConfigMint(&_Api.TransactOpts, value)
}

// ConfigMint is a paid mutator transaction binding the contract method 0xf6aef40c.
//
// Solidity: function configMint(uint256 value) returns(uint256)
func (_Api *ApiTransactorSession) ConfigMint(value *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ConfigMint(&_Api.TransactOpts, value)
}

// CreateEvent is a paid mutator transaction binding the contract method 0xbce062b9.
//
// Solidity: function createEvent(string _nameEvent, string _detial, uint256 _reward, uint256 _timeStart) returns(uint256)
func (_Api *ApiTransactor) CreateEvent(opts *bind.TransactOpts, _nameEvent string, _detial string, _reward *big.Int, _timeStart *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "createEvent", _nameEvent, _detial, _reward, _timeStart)
}

// CreateEvent is a paid mutator transaction binding the contract method 0xbce062b9.
//
// Solidity: function createEvent(string _nameEvent, string _detial, uint256 _reward, uint256 _timeStart) returns(uint256)
func (_Api *ApiSession) CreateEvent(_nameEvent string, _detial string, _reward *big.Int, _timeStart *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CreateEvent(&_Api.TransactOpts, _nameEvent, _detial, _reward, _timeStart)
}

// CreateEvent is a paid mutator transaction binding the contract method 0xbce062b9.
//
// Solidity: function createEvent(string _nameEvent, string _detial, uint256 _reward, uint256 _timeStart) returns(uint256)
func (_Api *ApiTransactorSession) CreateEvent(_nameEvent string, _detial string, _reward *big.Int, _timeStart *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CreateEvent(&_Api.TransactOpts, _nameEvent, _detial, _reward, _timeStart)
}

// CreateEventAdmin is a paid mutator transaction binding the contract method 0x32ace39c.
//
// Solidity: function createEventAdmin(string _nameEvent, string _detial, uint256 _reward, uint256 _timeStart) returns(uint256)
func (_Api *ApiTransactor) CreateEventAdmin(opts *bind.TransactOpts, _nameEvent string, _detial string, _reward *big.Int, _timeStart *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "createEventAdmin", _nameEvent, _detial, _reward, _timeStart)
}

// CreateEventAdmin is a paid mutator transaction binding the contract method 0x32ace39c.
//
// Solidity: function createEventAdmin(string _nameEvent, string _detial, uint256 _reward, uint256 _timeStart) returns(uint256)
func (_Api *ApiSession) CreateEventAdmin(_nameEvent string, _detial string, _reward *big.Int, _timeStart *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CreateEventAdmin(&_Api.TransactOpts, _nameEvent, _detial, _reward, _timeStart)
}

// CreateEventAdmin is a paid mutator transaction binding the contract method 0x32ace39c.
//
// Solidity: function createEventAdmin(string _nameEvent, string _detial, uint256 _reward, uint256 _timeStart) returns(uint256)
func (_Api *ApiTransactorSession) CreateEventAdmin(_nameEvent string, _detial string, _reward *big.Int, _timeStart *big.Int) (*types.Transaction, error) {
	return _Api.Contract.CreateEventAdmin(&_Api.TransactOpts, _nameEvent, _detial, _reward, _timeStart)
}

// JoinEvent is a paid mutator transaction binding the contract method 0x38d361e2.
//
// Solidity: function joinEvent(uint256 eid) returns()
func (_Api *ApiTransactor) JoinEvent(opts *bind.TransactOpts, eid *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "joinEvent", eid)
}

// JoinEvent is a paid mutator transaction binding the contract method 0x38d361e2.
//
// Solidity: function joinEvent(uint256 eid) returns()
func (_Api *ApiSession) JoinEvent(eid *big.Int) (*types.Transaction, error) {
	return _Api.Contract.JoinEvent(&_Api.TransactOpts, eid)
}

// JoinEvent is a paid mutator transaction binding the contract method 0x38d361e2.
//
// Solidity: function joinEvent(uint256 eid) returns()
func (_Api *ApiTransactorSession) JoinEvent(eid *big.Int) (*types.Transaction, error) {
	return _Api.Contract.JoinEvent(&_Api.TransactOpts, eid)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns()
func (_Api *ApiTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns()
func (_Api *ApiSession) Mint() (*types.Transaction, error) {
	return _Api.Contract.Mint(&_Api.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns()
func (_Api *ApiTransactorSession) Mint() (*types.Transaction, error) {
	return _Api.Contract.Mint(&_Api.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Api.Contract.RenounceOwnership(&_Api.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Api.Contract.RenounceOwnership(&_Api.TransactOpts)
}

// SearchEvent is a paid mutator transaction binding the contract method 0x16ea468f.
//
// Solidity: function searchEvent() returns((address,string,string,uint256,uint256,uint256)[])
func (_Api *ApiTransactor) SearchEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "searchEvent")
}

// SearchEvent is a paid mutator transaction binding the contract method 0x16ea468f.
//
// Solidity: function searchEvent() returns((address,string,string,uint256,uint256,uint256)[])
func (_Api *ApiSession) SearchEvent() (*types.Transaction, error) {
	return _Api.Contract.SearchEvent(&_Api.TransactOpts)
}

// SearchEvent is a paid mutator transaction binding the contract method 0x16ea468f.
//
// Solidity: function searchEvent() returns((address,string,string,uint256,uint256,uint256)[])
func (_Api *ApiTransactorSession) SearchEvent() (*types.Transaction, error) {
	return _Api.Contract.SearchEvent(&_Api.TransactOpts)
}

// SearchEventAdmin is a paid mutator transaction binding the contract method 0x781c0790.
//
// Solidity: function searchEventAdmin() returns(uint256[])
func (_Api *ApiTransactor) SearchEventAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "searchEventAdmin")
}

// SearchEventAdmin is a paid mutator transaction binding the contract method 0x781c0790.
//
// Solidity: function searchEventAdmin() returns(uint256[])
func (_Api *ApiSession) SearchEventAdmin() (*types.Transaction, error) {
	return _Api.Contract.SearchEventAdmin(&_Api.TransactOpts)
}

// SearchEventAdmin is a paid mutator transaction binding the contract method 0x781c0790.
//
// Solidity: function searchEventAdmin() returns(uint256[])
func (_Api *ApiTransactorSession) SearchEventAdmin() (*types.Transaction, error) {
	return _Api.Contract.SearchEventAdmin(&_Api.TransactOpts)
}

// SearchEventByAddress is a paid mutator transaction binding the contract method 0x6dc22b24.
//
// Solidity: function searchEventByAddress(address from) returns(uint256[])
func (_Api *ApiTransactor) SearchEventByAddress(opts *bind.TransactOpts, from common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "searchEventByAddress", from)
}

// SearchEventByAddress is a paid mutator transaction binding the contract method 0x6dc22b24.
//
// Solidity: function searchEventByAddress(address from) returns(uint256[])
func (_Api *ApiSession) SearchEventByAddress(from common.Address) (*types.Transaction, error) {
	return _Api.Contract.SearchEventByAddress(&_Api.TransactOpts, from)
}

// SearchEventByAddress is a paid mutator transaction binding the contract method 0x6dc22b24.
//
// Solidity: function searchEventByAddress(address from) returns(uint256[])
func (_Api *ApiTransactorSession) SearchEventByAddress(from common.Address) (*types.Transaction, error) {
	return _Api.Contract.SearchEventByAddress(&_Api.TransactOpts, from)
}

// SearchParticipant is a paid mutator transaction binding the contract method 0xa637aed4.
//
// Solidity: function searchParticipant(uint256 eid) returns(address[])
func (_Api *ApiTransactor) SearchParticipant(opts *bind.TransactOpts, eid *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "searchParticipant", eid)
}

// SearchParticipant is a paid mutator transaction binding the contract method 0xa637aed4.
//
// Solidity: function searchParticipant(uint256 eid) returns(address[])
func (_Api *ApiSession) SearchParticipant(eid *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SearchParticipant(&_Api.TransactOpts, eid)
}

// SearchParticipant is a paid mutator transaction binding the contract method 0xa637aed4.
//
// Solidity: function searchParticipant(uint256 eid) returns(address[])
func (_Api *ApiTransactorSession) SearchParticipant(eid *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SearchParticipant(&_Api.TransactOpts, eid)
}

// StartEvent is a paid mutator transaction binding the contract method 0xa59a6953.
//
// Solidity: function startEvent(uint256 eid) returns()
func (_Api *ApiTransactor) StartEvent(opts *bind.TransactOpts, eid *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "startEvent", eid)
}

// StartEvent is a paid mutator transaction binding the contract method 0xa59a6953.
//
// Solidity: function startEvent(uint256 eid) returns()
func (_Api *ApiSession) StartEvent(eid *big.Int) (*types.Transaction, error) {
	return _Api.Contract.StartEvent(&_Api.TransactOpts, eid)
}

// StartEvent is a paid mutator transaction binding the contract method 0xa59a6953.
//
// Solidity: function startEvent(uint256 eid) returns()
func (_Api *ApiTransactorSession) StartEvent(eid *big.Int) (*types.Transaction, error) {
	return _Api.Contract.StartEvent(&_Api.TransactOpts, eid)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Api.Contract.TransferOwnership(&_Api.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Api.Contract.TransferOwnership(&_Api.TransactOpts, newOwner)
}

// Vote is a paid mutator transaction binding the contract method 0x5f74bbde.
//
// Solidity: function vote(address _to, uint256 score) returns()
func (_Api *ApiTransactor) Vote(opts *bind.TransactOpts, _to common.Address, score *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "vote", _to, score)
}

// Vote is a paid mutator transaction binding the contract method 0x5f74bbde.
//
// Solidity: function vote(address _to, uint256 score) returns()
func (_Api *ApiSession) Vote(_to common.Address, score *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Vote(&_Api.TransactOpts, _to, score)
}

// Vote is a paid mutator transaction binding the contract method 0x5f74bbde.
//
// Solidity: function vote(address _to, uint256 score) returns()
func (_Api *ApiTransactorSession) Vote(_to common.Address, score *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Vote(&_Api.TransactOpts, _to, score)
}

// ApiLogVoteIterator is returned from FilterLogVote and is used to iterate over the raw logs and unpacked data for LogVote events raised by the Api contract.
type ApiLogVoteIterator struct {
	Event *ApiLogVote // Event containing the contract specifics and raw log

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
func (it *ApiLogVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiLogVote)
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
		it.Event = new(ApiLogVote)
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
func (it *ApiLogVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiLogVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiLogVote represents a LogVote event raised by the Api contract.
type ApiLogVote struct {
	From  common.Address
	To    common.Address
	Score *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogVote is a free log retrieval operation binding the contract event 0x49ce5cb7b86410ac7069ff893207f2804cf4614b4203eaf4e0e37bb41a2b0ef0.
//
// Solidity: event LogVote(address indexed from, address indexed _to, uint256 score)
func (_Api *ApiFilterer) FilterLogVote(opts *bind.FilterOpts, from []common.Address, _to []common.Address) (*ApiLogVoteIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "LogVote", fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ApiLogVoteIterator{contract: _Api.contract, event: "LogVote", logs: logs, sub: sub}, nil
}

// WatchLogVote is a free log subscription operation binding the contract event 0x49ce5cb7b86410ac7069ff893207f2804cf4614b4203eaf4e0e37bb41a2b0ef0.
//
// Solidity: event LogVote(address indexed from, address indexed _to, uint256 score)
func (_Api *ApiFilterer) WatchLogVote(opts *bind.WatchOpts, sink chan<- *ApiLogVote, from []common.Address, _to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "LogVote", fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiLogVote)
				if err := _Api.contract.UnpackLog(event, "LogVote", log); err != nil {
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

// ParseLogVote is a log parse operation binding the contract event 0x49ce5cb7b86410ac7069ff893207f2804cf4614b4203eaf4e0e37bb41a2b0ef0.
//
// Solidity: event LogVote(address indexed from, address indexed _to, uint256 score)
func (_Api *ApiFilterer) ParseLogVote(log types.Log) (*ApiLogVote, error) {
	event := new(ApiLogVote)
	if err := _Api.contract.UnpackLog(event, "LogVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Api contract.
type ApiOwnershipTransferredIterator struct {
	Event *ApiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ApiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiOwnershipTransferred)
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
		it.Event = new(ApiOwnershipTransferred)
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
func (it *ApiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiOwnershipTransferred represents a OwnershipTransferred event raised by the Api contract.
type ApiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ApiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ApiOwnershipTransferredIterator{contract: _Api.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ApiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiOwnershipTransferred)
				if err := _Api.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) ParseOwnershipTransferred(log types.Log) (*ApiOwnershipTransferred, error) {
	event := new(ApiOwnershipTransferred)
	if err := _Api.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiCreateEventIterator is returned from FilterCreateEvent and is used to iterate over the raw logs and unpacked data for CreateEvent events raised by the Api contract.
type ApiCreateEventIterator struct {
	Event *ApiCreateEvent // Event containing the contract specifics and raw log

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
func (it *ApiCreateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiCreateEvent)
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
		it.Event = new(ApiCreateEvent)
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
func (it *ApiCreateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiCreateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiCreateEvent represents a CreateEvent event raised by the Api contract.
type ApiCreateEvent struct {
	CreateBy  common.Address
	NameEvent string
	Detial    string
	Reward    *big.Int
	TimeStart *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateEvent is a free log retrieval operation binding the contract event 0x71f5db2c13452fa93276639f35267a3e4a4c979bf4f10bd87ea535b4d975cb6a.
//
// Solidity: event _createEvent(address _createBy, string _nameEvent, string _detial, uint256 _reward, uint256 _timeStart)
func (_Api *ApiFilterer) FilterCreateEvent(opts *bind.FilterOpts) (*ApiCreateEventIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "_createEvent")
	if err != nil {
		return nil, err
	}
	return &ApiCreateEventIterator{contract: _Api.contract, event: "_createEvent", logs: logs, sub: sub}, nil
}

// WatchCreateEvent is a free log subscription operation binding the contract event 0x71f5db2c13452fa93276639f35267a3e4a4c979bf4f10bd87ea535b4d975cb6a.
//
// Solidity: event _createEvent(address _createBy, string _nameEvent, string _detial, uint256 _reward, uint256 _timeStart)
func (_Api *ApiFilterer) WatchCreateEvent(opts *bind.WatchOpts, sink chan<- *ApiCreateEvent) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "_createEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiCreateEvent)
				if err := _Api.contract.UnpackLog(event, "_createEvent", log); err != nil {
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

// ParseCreateEvent is a log parse operation binding the contract event 0x71f5db2c13452fa93276639f35267a3e4a4c979bf4f10bd87ea535b4d975cb6a.
//
// Solidity: event _createEvent(address _createBy, string _nameEvent, string _detial, uint256 _reward, uint256 _timeStart)
func (_Api *ApiFilterer) ParseCreateEvent(log types.Log) (*ApiCreateEvent, error) {
	event := new(ApiCreateEvent)
	if err := _Api.contract.UnpackLog(event, "_createEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiJoinIterator is returned from FilterJoin and is used to iterate over the raw logs and unpacked data for Join events raised by the Api contract.
type ApiJoinIterator struct {
	Event *ApiJoin // Event containing the contract specifics and raw log

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
func (it *ApiJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiJoin)
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
		it.Event = new(ApiJoin)
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
func (it *ApiJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiJoin represents a Join event raised by the Api contract.
type ApiJoin struct {
	From common.Address
	Eid  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterJoin is a free log retrieval operation binding the contract event 0x2e55f4a4bcfc21d67be1e2e465ef0c4b6a9a3bce54eb1a9754dae16fc63828a2.
//
// Solidity: event _join(address from, uint256 eid)
func (_Api *ApiFilterer) FilterJoin(opts *bind.FilterOpts) (*ApiJoinIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "_join")
	if err != nil {
		return nil, err
	}
	return &ApiJoinIterator{contract: _Api.contract, event: "_join", logs: logs, sub: sub}, nil
}

// WatchJoin is a free log subscription operation binding the contract event 0x2e55f4a4bcfc21d67be1e2e465ef0c4b6a9a3bce54eb1a9754dae16fc63828a2.
//
// Solidity: event _join(address from, uint256 eid)
func (_Api *ApiFilterer) WatchJoin(opts *bind.WatchOpts, sink chan<- *ApiJoin) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "_join")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiJoin)
				if err := _Api.contract.UnpackLog(event, "_join", log); err != nil {
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

// ParseJoin is a log parse operation binding the contract event 0x2e55f4a4bcfc21d67be1e2e465ef0c4b6a9a3bce54eb1a9754dae16fc63828a2.
//
// Solidity: event _join(address from, uint256 eid)
func (_Api *ApiFilterer) ParseJoin(log types.Log) (*ApiJoin, error) {
	event := new(ApiJoin)
	if err := _Api.contract.UnpackLog(event, "_join", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiGetScoreVoteIterator is returned from FilterGetScoreVote and is used to iterate over the raw logs and unpacked data for GetScoreVote events raised by the Api contract.
type ApiGetScoreVoteIterator struct {
	Event *ApiGetScoreVote // Event containing the contract specifics and raw log

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
func (it *ApiGetScoreVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiGetScoreVote)
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
		it.Event = new(ApiGetScoreVote)
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
func (it *ApiGetScoreVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiGetScoreVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiGetScoreVote represents a GetScoreVote event raised by the Api contract.
type ApiGetScoreVote struct {
	ListWallet []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGetScoreVote is a free log retrieval operation binding the contract event 0xb376534b1fb0889c24ac054e5b06cbc118d462433346a80ced64af664d889df6.
//
// Solidity: event getScoreVote(address[] listWallet)
func (_Api *ApiFilterer) FilterGetScoreVote(opts *bind.FilterOpts) (*ApiGetScoreVoteIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "getScoreVote")
	if err != nil {
		return nil, err
	}
	return &ApiGetScoreVoteIterator{contract: _Api.contract, event: "getScoreVote", logs: logs, sub: sub}, nil
}

// WatchGetScoreVote is a free log subscription operation binding the contract event 0xb376534b1fb0889c24ac054e5b06cbc118d462433346a80ced64af664d889df6.
//
// Solidity: event getScoreVote(address[] listWallet)
func (_Api *ApiFilterer) WatchGetScoreVote(opts *bind.WatchOpts, sink chan<- *ApiGetScoreVote) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "getScoreVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiGetScoreVote)
				if err := _Api.contract.UnpackLog(event, "getScoreVote", log); err != nil {
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

// ParseGetScoreVote is a log parse operation binding the contract event 0xb376534b1fb0889c24ac054e5b06cbc118d462433346a80ced64af664d889df6.
//
// Solidity: event getScoreVote(address[] listWallet)
func (_Api *ApiFilterer) ParseGetScoreVote(log types.Log) (*ApiGetScoreVote, error) {
	event := new(ApiGetScoreVote)
	if err := _Api.contract.UnpackLog(event, "getScoreVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
