// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lucidNft

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LucidNftMetaData contains all meta data concerning the LucidNft contract.
var LucidNftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getMints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600580546001600160a01b03191670d883694d2fcc3f8ed53dc3b14d18c134271790553480156200003457600080fd5b5060405180606001604052806034815260200162001d3d603491396200005a816200006c565b5062000066336200007e565b62000241565b60026200007a828262000175565b5050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620000fb57607f821691505b6020821081036200011c57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200017057600081815260208120601f850160051c810160208610156200014b5750805b601f850160051c820191505b818110156200016c5782815560010162000157565b5050505b505050565b81516001600160401b03811115620001915762000191620000d0565b620001a981620001a28454620000e6565b8462000122565b602080601f831160018114620001e15760008415620001c85750858301515b600019600386901b1c1916600185901b1785556200016c565b600085815260208120601f198616915b828110156200021257888601518255948401946001909101908401620001f1565b5085821015620002315787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b611aec80620002516000396000f3fe608060405234801561001057600080fd5b50600436106100f45760003560e01c806359e1abaa11610097578063b49c99b811610066578063b49c99b814610206578063e985e9c514610219578063f242432a14610255578063f2fde38b1461026857600080fd5b806359e1abaa146101bd578063715018a6146101d05780638da5cb5b146101d8578063a22cb465146101f357600080fd5b806321335101116100d357806321335101146101625780632eb2c2d6146101825780634e1273f414610197578063564b81ef146101b757600080fd5b8062fdd58e146100f957806301ffc9a71461011f5780630e89341c14610142575b600080fd5b61010c610107366004611263565b61027b565b6040519081526020015b60405180910390f35b61013261012d3660046112a3565b610315565b6040519015158152602001610116565b6101556101503660046112c7565b610365565b6040516101169190611326565b61010c6101703660046112c7565b60009081526004602052604090205490565b610195610190366004611485565b6103f9565b005b6101aa6101a536600461152f565b610490565b6040516101169190611635565b4661010c565b61010c6101cb366004611648565b6105ba565b61019561061c565b6003546040516001600160a01b039091168152602001610116565b610195610201366004611695565b610682565b6101956102143660046116d1565b610691565b61013261022736600461172f565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b610195610263366004611762565b610854565b6101956102763660046117bb565b6108db565b60006001600160a01b0383166102ec5760405162461bcd60e51b815260206004820152602b60248201527f455243313135353a2062616c616e636520717565727920666f7220746865207a60448201526a65726f206164647265737360a81b60648201526084015b60405180910390fd5b506000818152602081815260408083206001600160a01b03861684529091529020545b92915050565b60006001600160e01b03198216636cdb3d1360e11b148061034657506001600160e01b031982166303a24d0760e21b145b8061030f57506301ffc9a760e01b6001600160e01b031983161461030f565b606060028054610374906117d6565b80601f01602080910402602001604051908101604052809291908181526020018280546103a0906117d6565b80156103ed5780601f106103c2576101008083540402835291602001916103ed565b820191906000526020600020905b8154815290600101906020018083116103d057829003601f168201915b50505050509050919050565b6001600160a01b03851633148061041557506104158533610227565b61047c5760405162461bcd60e51b815260206004820152603260248201527f455243313135353a207472616e736665722063616c6c6572206973206e6f74206044820152711bdddb995c881b9bdc88185c1c1c9bdd995960721b60648201526084016102e3565b61048985858585856109a6565b5050505050565b606081518351146104f55760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b60648201526084016102e3565b6000835167ffffffffffffffff81111561051157610511611339565b60405190808252806020026020018201604052801561053a578160200160208202803683370190505b50905060005b84518110156105b25761058585828151811061055e5761055e611810565b602002602001015185838151811061057857610578611810565b602002602001015161027b565b82828151811061059757610597611810565b60209081029190910101526105ab8161183c565b9050610540565b509392505050565b6040516bffffffffffffffffffffffff19606087811b8216602084015286901b16603482015260488101849052606881018390526088810182905260009060a80160405160208183030381529060405280519060200120905095945050505050565b6003546001600160a01b031633146106765760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102e3565b6106806000610b83565b565b61068d338383610bd5565b5050565b848483836000306001600160a01b031663564b81ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f99190611855565b9050600061070a86308488886105ba565b90506000610765826040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b905060006107738286610cb5565b6005549091506001600160a01b038083169116146108075760405162461bcd60e51b815260206004820152604560248201527f4c6f6c20796f7572206e6f74206f6e207468652077686974656c69737420766960448201527f7374202877652063616e206861726420636f6465205369676e65722055524c20606482015264686572652960d81b608482015260a4016102e3565b610822338c8e60405180602001604052806000815250610d34565b60008b815260046020526040812080548e929061084090849061186e565b909155505050505050505050505050505050565b6001600160a01b03851633148061087057506108708533610227565b6108ce5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7260448201526808185c1c1c9bdd995960ba1b60648201526084016102e3565b6104898585858585610e48565b6003546001600160a01b031633146109355760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102e3565b6001600160a01b03811661099a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102e3565b6109a381610b83565b50565b8151835114610a085760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b60648201526084016102e3565b6001600160a01b038416610a2e5760405162461bcd60e51b81526004016102e390611881565b3360005b8451811015610b15576000858281518110610a4f57610a4f611810565b602002602001015190506000858381518110610a6d57610a6d611810565b602090810291909101810151600084815280835260408082206001600160a01b038e168352909352919091205490915081811015610abd5760405162461bcd60e51b81526004016102e3906118c6565b6000838152602081815260408083206001600160a01b038e8116855292528083208585039055908b16825281208054849290610afa90849061186e565b9250508190555050505080610b0e9061183c565b9050610a32565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610b65929190611910565b60405180910390a4610b7b818787878787610f72565b505050505050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b816001600160a01b0316836001600160a01b031603610c485760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b60648201526084016102e3565b6001600160a01b03838116600081815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b600080600080610cc4856110cd565b6040805160008152602081018083528b905260ff8316918101919091526060810184905260808101839052929550909350915060019060a0016020604051602081039080840390855afa158015610d1f573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b6001600160a01b038416610d945760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b60648201526084016102e3565b336000610da085611141565b90506000610dad85611141565b90506000868152602081815260408083206001600160a01b038b16845290915281208054879290610ddf90849061186e565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610e3f8360008989898961118c565b50505050505050565b6001600160a01b038416610e6e5760405162461bcd60e51b81526004016102e390611881565b336000610e7a85611141565b90506000610e8785611141565b90506000868152602081815260408083206001600160a01b038c16845290915290205485811015610eca5760405162461bcd60e51b81526004016102e3906118c6565b6000878152602081815260408083206001600160a01b038d8116855292528083208985039055908a16825281208054889290610f0790849061186e565b909155505060408051888152602081018890526001600160a01b03808b16928c821692918816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610f67848a8a8a8a8a61118c565b505050505050505050565b6001600160a01b0384163b15610b7b5760405163bc197c8160e01b81526001600160a01b0385169063bc197c8190610fb6908990899088908890889060040161193e565b6020604051808303816000875af1925050508015610ff1575060408051601f3d908101601f19168201909252610fee9181019061199c565b60015b61109d57610ffd6119b9565b806308c379a00361103657506110116119d5565b8061101c5750611038565b8060405162461bcd60e51b81526004016102e39190611326565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b60648201526084016102e3565b6001600160e01b0319811663bc197c8160e01b14610e3f5760405162461bcd60e51b81526004016102e390611a5f565b600080600083516041146111235760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207369676e6174757265206c656e677468000000000000000060448201526064016102e3565b50505060208101516040820151606090920151909260009190911a90565b6040805160018082528183019092526060916000919060208083019080368337019050509050828160008151811061117b5761117b611810565b602090810291909101015292915050565b6001600160a01b0384163b15610b7b5760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906111d09089908990889088908890600401611aa7565b6020604051808303816000875af192505050801561120b575060408051601f3d908101601f191682019092526112089181019061199c565b60015b61121757610ffd6119b9565b6001600160e01b0319811663f23a6e6160e01b14610e3f5760405162461bcd60e51b81526004016102e390611a5f565b80356001600160a01b038116811461125e57600080fd5b919050565b6000806040838503121561127657600080fd5b61127f83611247565b946020939093013593505050565b6001600160e01b0319811681146109a357600080fd5b6000602082840312156112b557600080fd5b81356112c08161128d565b9392505050565b6000602082840312156112d957600080fd5b5035919050565b6000815180845260005b81811015611306576020818501810151868301820152016112ea565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006112c060208301846112e0565b634e487b7160e01b600052604160045260246000fd5b601f8201601f1916810167ffffffffffffffff8111828210171561137557611375611339565b6040525050565b600067ffffffffffffffff82111561139657611396611339565b5060051b60200190565b600082601f8301126113b157600080fd5b813560206113be8261137c565b6040516113cb828261134f565b83815260059390931b85018201928281019150868411156113eb57600080fd5b8286015b8481101561140657803583529183019183016113ef565b509695505050505050565b600082601f83011261142257600080fd5b813567ffffffffffffffff81111561143c5761143c611339565b604051611453601f8301601f19166020018261134f565b81815284602083860101111561146857600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a0868803121561149d57600080fd5b6114a686611247565b94506114b460208701611247565b9350604086013567ffffffffffffffff808211156114d157600080fd5b6114dd89838a016113a0565b945060608801359150808211156114f357600080fd5b6114ff89838a016113a0565b9350608088013591508082111561151557600080fd5b5061152288828901611411565b9150509295509295909350565b6000806040838503121561154257600080fd5b823567ffffffffffffffff8082111561155a57600080fd5b818501915085601f83011261156e57600080fd5b8135602061157b8261137c565b604051611588828261134f565b83815260059390931b85018201928281019150898411156115a857600080fd5b948201945b838610156115cd576115be86611247565b825294820194908201906115ad565b965050860135925050808211156115e357600080fd5b506115f0858286016113a0565b9150509250929050565b600081518084526020808501945080840160005b8381101561162a5781518752958201959082019060010161160e565b509495945050505050565b6020815260006112c060208301846115fa565b600080600080600060a0868803121561166057600080fd5b61166986611247565b945061167760208701611247565b94979496505050506040830135926060810135926080909101359150565b600080604083850312156116a857600080fd5b6116b183611247565b9150602083013580151581146116c657600080fd5b809150509250929050565b600080600080600060a086880312156116e957600080fd5b6116f286611247565b9450602086013593506040860135925060608601359150608086013567ffffffffffffffff81111561172357600080fd5b61152288828901611411565b6000806040838503121561174257600080fd5b61174b83611247565b915061175960208401611247565b90509250929050565b600080600080600060a0868803121561177a57600080fd5b61178386611247565b945061179160208701611247565b93506040860135925060608601359150608086013567ffffffffffffffff81111561172357600080fd5b6000602082840312156117cd57600080fd5b6112c082611247565b600181811c908216806117ea57607f821691505b60208210810361180a57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161184e5761184e611826565b5060010190565b60006020828403121561186757600080fd5b5051919050565b8082018082111561030f5761030f611826565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b60408152600061192360408301856115fa565b828103602084015261193581856115fa565b95945050505050565b6001600160a01b0386811682528516602082015260a06040820181905260009061196a908301866115fa565b828103606084015261197c81866115fa565b9050828103608084015261199081856112e0565b98975050505050505050565b6000602082840312156119ae57600080fd5b81516112c08161128d565b600060033d11156119d25760046000803e5060005160e01c5b90565b600060443d10156119e35790565b6040516003193d81016004833e81513d67ffffffffffffffff8160248401118184111715611a1357505050505090565b8285019150815181811115611a2b5750505050505090565b843d8701016020828501011115611a455750505050505090565b611a546020828601018761134f565b509095945050505050565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b6001600160a01b03868116825285166020820152604081018490526060810183905260a060808201819052600090611ae1908301846112e0565b9796505050505050505668747470733a2f2f7261696e626f776b69742d6d696e742d6e66742d64656d6f2e76657263656c2e6170702f6e66742e6a736f6e",
}

// LucidNftABI is the input ABI used to generate the binding from.
// Deprecated: Use LucidNftMetaData.ABI instead.
var LucidNftABI = LucidNftMetaData.ABI

// LucidNftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LucidNftMetaData.Bin instead.
var LucidNftBin = LucidNftMetaData.Bin

// DeployLucidNft deploys a new Ethereum contract, binding an instance of LucidNft to it.
func DeployLucidNft(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LucidNft, error) {
	parsed, err := LucidNftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LucidNftBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LucidNft{LucidNftCaller: LucidNftCaller{contract: contract}, LucidNftTransactor: LucidNftTransactor{contract: contract}, LucidNftFilterer: LucidNftFilterer{contract: contract}}, nil
}

// LucidNft is an auto generated Go binding around an Ethereum contract.
type LucidNft struct {
	LucidNftCaller     // Read-only binding to the contract
	LucidNftTransactor // Write-only binding to the contract
	LucidNftFilterer   // Log filterer for contract events
}

// LucidNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type LucidNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LucidNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LucidNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LucidNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LucidNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LucidNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LucidNftSession struct {
	Contract     *LucidNft         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LucidNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LucidNftCallerSession struct {
	Contract *LucidNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LucidNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LucidNftTransactorSession struct {
	Contract     *LucidNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LucidNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type LucidNftRaw struct {
	Contract *LucidNft // Generic contract binding to access the raw methods on
}

// LucidNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LucidNftCallerRaw struct {
	Contract *LucidNftCaller // Generic read-only contract binding to access the raw methods on
}

// LucidNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LucidNftTransactorRaw struct {
	Contract *LucidNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLucidNft creates a new instance of LucidNft, bound to a specific deployed contract.
func NewLucidNft(address common.Address, backend bind.ContractBackend) (*LucidNft, error) {
	contract, err := bindLucidNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LucidNft{LucidNftCaller: LucidNftCaller{contract: contract}, LucidNftTransactor: LucidNftTransactor{contract: contract}, LucidNftFilterer: LucidNftFilterer{contract: contract}}, nil
}

// NewLucidNftCaller creates a new read-only instance of LucidNft, bound to a specific deployed contract.
func NewLucidNftCaller(address common.Address, caller bind.ContractCaller) (*LucidNftCaller, error) {
	contract, err := bindLucidNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LucidNftCaller{contract: contract}, nil
}

// NewLucidNftTransactor creates a new write-only instance of LucidNft, bound to a specific deployed contract.
func NewLucidNftTransactor(address common.Address, transactor bind.ContractTransactor) (*LucidNftTransactor, error) {
	contract, err := bindLucidNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LucidNftTransactor{contract: contract}, nil
}

// NewLucidNftFilterer creates a new log filterer instance of LucidNft, bound to a specific deployed contract.
func NewLucidNftFilterer(address common.Address, filterer bind.ContractFilterer) (*LucidNftFilterer, error) {
	contract, err := bindLucidNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LucidNftFilterer{contract: contract}, nil
}

// bindLucidNft binds a generic wrapper to an already deployed contract.
func bindLucidNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LucidNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LucidNft *LucidNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LucidNft.Contract.LucidNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LucidNft *LucidNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LucidNft.Contract.LucidNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LucidNft *LucidNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LucidNft.Contract.LucidNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LucidNft *LucidNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LucidNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LucidNft *LucidNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LucidNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LucidNft *LucidNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LucidNft.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_LucidNft *LucidNftCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LucidNft.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_LucidNft *LucidNftSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _LucidNft.Contract.BalanceOf(&_LucidNft.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_LucidNft *LucidNftCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _LucidNft.Contract.BalanceOf(&_LucidNft.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_LucidNft *LucidNftCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _LucidNft.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_LucidNft *LucidNftSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _LucidNft.Contract.BalanceOfBatch(&_LucidNft.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_LucidNft *LucidNftCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _LucidNft.Contract.BalanceOfBatch(&_LucidNft.CallOpts, accounts, ids)
}

// GetChainID is a free data retrieval call binding the contract method 0x564b81ef.
//
// Solidity: function getChainID() view returns(uint256)
func (_LucidNft *LucidNftCaller) GetChainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LucidNft.contract.Call(opts, &out, "getChainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainID is a free data retrieval call binding the contract method 0x564b81ef.
//
// Solidity: function getChainID() view returns(uint256)
func (_LucidNft *LucidNftSession) GetChainID() (*big.Int, error) {
	return _LucidNft.Contract.GetChainID(&_LucidNft.CallOpts)
}

// GetChainID is a free data retrieval call binding the contract method 0x564b81ef.
//
// Solidity: function getChainID() view returns(uint256)
func (_LucidNft *LucidNftCallerSession) GetChainID() (*big.Int, error) {
	return _LucidNft.Contract.GetChainID(&_LucidNft.CallOpts)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x59e1abaa.
//
// Solidity: function getMessageHash(address _to, address _contractAddress, uint256 _chain, uint256 _amount, uint256 _nonce) pure returns(bytes32)
func (_LucidNft *LucidNftCaller) GetMessageHash(opts *bind.CallOpts, _to common.Address, _contractAddress common.Address, _chain *big.Int, _amount *big.Int, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LucidNft.contract.Call(opts, &out, "getMessageHash", _to, _contractAddress, _chain, _amount, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x59e1abaa.
//
// Solidity: function getMessageHash(address _to, address _contractAddress, uint256 _chain, uint256 _amount, uint256 _nonce) pure returns(bytes32)
func (_LucidNft *LucidNftSession) GetMessageHash(_to common.Address, _contractAddress common.Address, _chain *big.Int, _amount *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _LucidNft.Contract.GetMessageHash(&_LucidNft.CallOpts, _to, _contractAddress, _chain, _amount, _nonce)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x59e1abaa.
//
// Solidity: function getMessageHash(address _to, address _contractAddress, uint256 _chain, uint256 _amount, uint256 _nonce) pure returns(bytes32)
func (_LucidNft *LucidNftCallerSession) GetMessageHash(_to common.Address, _contractAddress common.Address, _chain *big.Int, _amount *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _LucidNft.Contract.GetMessageHash(&_LucidNft.CallOpts, _to, _contractAddress, _chain, _amount, _nonce)
}

// GetMints is a free data retrieval call binding the contract method 0x21335101.
//
// Solidity: function getMints(uint256 _tokenId) view returns(uint256)
func (_LucidNft *LucidNftCaller) GetMints(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LucidNft.contract.Call(opts, &out, "getMints", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMints is a free data retrieval call binding the contract method 0x21335101.
//
// Solidity: function getMints(uint256 _tokenId) view returns(uint256)
func (_LucidNft *LucidNftSession) GetMints(_tokenId *big.Int) (*big.Int, error) {
	return _LucidNft.Contract.GetMints(&_LucidNft.CallOpts, _tokenId)
}

// GetMints is a free data retrieval call binding the contract method 0x21335101.
//
// Solidity: function getMints(uint256 _tokenId) view returns(uint256)
func (_LucidNft *LucidNftCallerSession) GetMints(_tokenId *big.Int) (*big.Int, error) {
	return _LucidNft.Contract.GetMints(&_LucidNft.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_LucidNft *LucidNftCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LucidNft.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_LucidNft *LucidNftSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _LucidNft.Contract.IsApprovedForAll(&_LucidNft.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_LucidNft *LucidNftCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _LucidNft.Contract.IsApprovedForAll(&_LucidNft.CallOpts, account, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LucidNft *LucidNftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LucidNft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LucidNft *LucidNftSession) Owner() (common.Address, error) {
	return _LucidNft.Contract.Owner(&_LucidNft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LucidNft *LucidNftCallerSession) Owner() (common.Address, error) {
	return _LucidNft.Contract.Owner(&_LucidNft.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LucidNft *LucidNftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LucidNft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LucidNft *LucidNftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LucidNft.Contract.SupportsInterface(&_LucidNft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LucidNft *LucidNftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LucidNft.Contract.SupportsInterface(&_LucidNft.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_LucidNft *LucidNftCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _LucidNft.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_LucidNft *LucidNftSession) Uri(arg0 *big.Int) (string, error) {
	return _LucidNft.Contract.Uri(&_LucidNft.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_LucidNft *LucidNftCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _LucidNft.Contract.Uri(&_LucidNft.CallOpts, arg0)
}

// Mint is a paid mutator transaction binding the contract method 0xb49c99b8.
//
// Solidity: function mint(address _to, uint256 _amount, uint256 _tokenId, uint256 _nonce, bytes signature) returns()
func (_LucidNft *LucidNftTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _tokenId *big.Int, _nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _LucidNft.contract.Transact(opts, "mint", _to, _amount, _tokenId, _nonce, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xb49c99b8.
//
// Solidity: function mint(address _to, uint256 _amount, uint256 _tokenId, uint256 _nonce, bytes signature) returns()
func (_LucidNft *LucidNftSession) Mint(_to common.Address, _amount *big.Int, _tokenId *big.Int, _nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _LucidNft.Contract.Mint(&_LucidNft.TransactOpts, _to, _amount, _tokenId, _nonce, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xb49c99b8.
//
// Solidity: function mint(address _to, uint256 _amount, uint256 _tokenId, uint256 _nonce, bytes signature) returns()
func (_LucidNft *LucidNftTransactorSession) Mint(_to common.Address, _amount *big.Int, _tokenId *big.Int, _nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _LucidNft.Contract.Mint(&_LucidNft.TransactOpts, _to, _amount, _tokenId, _nonce, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LucidNft *LucidNftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LucidNft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LucidNft *LucidNftSession) RenounceOwnership() (*types.Transaction, error) {
	return _LucidNft.Contract.RenounceOwnership(&_LucidNft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LucidNft *LucidNftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LucidNft.Contract.RenounceOwnership(&_LucidNft.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_LucidNft *LucidNftTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _LucidNft.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_LucidNft *LucidNftSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _LucidNft.Contract.SafeBatchTransferFrom(&_LucidNft.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_LucidNft *LucidNftTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _LucidNft.Contract.SafeBatchTransferFrom(&_LucidNft.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_LucidNft *LucidNftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _LucidNft.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_LucidNft *LucidNftSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _LucidNft.Contract.SafeTransferFrom(&_LucidNft.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_LucidNft *LucidNftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _LucidNft.Contract.SafeTransferFrom(&_LucidNft.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LucidNft *LucidNftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _LucidNft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LucidNft *LucidNftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LucidNft.Contract.SetApprovalForAll(&_LucidNft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LucidNft *LucidNftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LucidNft.Contract.SetApprovalForAll(&_LucidNft.TransactOpts, operator, approved)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LucidNft *LucidNftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LucidNft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LucidNft *LucidNftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LucidNft.Contract.TransferOwnership(&_LucidNft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LucidNft *LucidNftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LucidNft.Contract.TransferOwnership(&_LucidNft.TransactOpts, newOwner)
}

// LucidNftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the LucidNft contract.
type LucidNftApprovalForAllIterator struct {
	Event *LucidNftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LucidNftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LucidNftApprovalForAll)
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
		it.Event = new(LucidNftApprovalForAll)
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
func (it *LucidNftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LucidNftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LucidNftApprovalForAll represents a ApprovalForAll event raised by the LucidNft contract.
type LucidNftApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_LucidNft *LucidNftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*LucidNftApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LucidNft.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LucidNftApprovalForAllIterator{contract: _LucidNft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_LucidNft *LucidNftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LucidNftApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LucidNft.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LucidNftApprovalForAll)
				if err := _LucidNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_LucidNft *LucidNftFilterer) ParseApprovalForAll(log types.Log) (*LucidNftApprovalForAll, error) {
	event := new(LucidNftApprovalForAll)
	if err := _LucidNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LucidNftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LucidNft contract.
type LucidNftOwnershipTransferredIterator struct {
	Event *LucidNftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LucidNftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LucidNftOwnershipTransferred)
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
		it.Event = new(LucidNftOwnershipTransferred)
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
func (it *LucidNftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LucidNftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LucidNftOwnershipTransferred represents a OwnershipTransferred event raised by the LucidNft contract.
type LucidNftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LucidNft *LucidNftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LucidNftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LucidNft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LucidNftOwnershipTransferredIterator{contract: _LucidNft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LucidNft *LucidNftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LucidNftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LucidNft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LucidNftOwnershipTransferred)
				if err := _LucidNft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LucidNft *LucidNftFilterer) ParseOwnershipTransferred(log types.Log) (*LucidNftOwnershipTransferred, error) {
	event := new(LucidNftOwnershipTransferred)
	if err := _LucidNft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LucidNftTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the LucidNft contract.
type LucidNftTransferBatchIterator struct {
	Event *LucidNftTransferBatch // Event containing the contract specifics and raw log

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
func (it *LucidNftTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LucidNftTransferBatch)
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
		it.Event = new(LucidNftTransferBatch)
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
func (it *LucidNftTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LucidNftTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LucidNftTransferBatch represents a TransferBatch event raised by the LucidNft contract.
type LucidNftTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_LucidNft *LucidNftFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*LucidNftTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LucidNft.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LucidNftTransferBatchIterator{contract: _LucidNft.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_LucidNft *LucidNftFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *LucidNftTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LucidNft.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LucidNftTransferBatch)
				if err := _LucidNft.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_LucidNft *LucidNftFilterer) ParseTransferBatch(log types.Log) (*LucidNftTransferBatch, error) {
	event := new(LucidNftTransferBatch)
	if err := _LucidNft.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LucidNftTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the LucidNft contract.
type LucidNftTransferSingleIterator struct {
	Event *LucidNftTransferSingle // Event containing the contract specifics and raw log

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
func (it *LucidNftTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LucidNftTransferSingle)
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
		it.Event = new(LucidNftTransferSingle)
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
func (it *LucidNftTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LucidNftTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LucidNftTransferSingle represents a TransferSingle event raised by the LucidNft contract.
type LucidNftTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_LucidNft *LucidNftFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*LucidNftTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LucidNft.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LucidNftTransferSingleIterator{contract: _LucidNft.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_LucidNft *LucidNftFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *LucidNftTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LucidNft.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LucidNftTransferSingle)
				if err := _LucidNft.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_LucidNft *LucidNftFilterer) ParseTransferSingle(log types.Log) (*LucidNftTransferSingle, error) {
	event := new(LucidNftTransferSingle)
	if err := _LucidNft.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LucidNftURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the LucidNft contract.
type LucidNftURIIterator struct {
	Event *LucidNftURI // Event containing the contract specifics and raw log

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
func (it *LucidNftURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LucidNftURI)
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
		it.Event = new(LucidNftURI)
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
func (it *LucidNftURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LucidNftURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LucidNftURI represents a URI event raised by the LucidNft contract.
type LucidNftURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_LucidNft *LucidNftFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*LucidNftURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LucidNft.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &LucidNftURIIterator{contract: _LucidNft.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_LucidNft *LucidNftFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *LucidNftURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LucidNft.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LucidNftURI)
				if err := _LucidNft.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_LucidNft *LucidNftFilterer) ParseURI(log types.Log) (*LucidNftURI, error) {
	event := new(LucidNftURI)
	if err := _LucidNft.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
