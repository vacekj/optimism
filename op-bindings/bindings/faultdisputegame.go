// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
)

// FaultDisputeGameMetaData contains all meta data concerning the FaultDisputeGame contract.
var FaultDisputeGameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"Claim\",\"name\":\"_absolutePrestate\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxGameDepth\",\"type\":\"uint256\"},{\"internalType\":\"contractIBigStepper\",\"name\":\"_vm\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotDefendRootClaim\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClockNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClockTimeExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GameDepthExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GameNotInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrestate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidStep\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"parentIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"Claim\",\"name\":\"claim\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"}],\"name\":\"Move\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumGameStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"Resolved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ABSOLUTE_PRESTATE\",\"outputs\":[{\"internalType\":\"Claim\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_GAME_DEPTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VM\",\"outputs\":[{\"internalType\":\"contractIBigStepper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_parentIndex\",\"type\":\"uint256\"},{\"internalType\":\"Claim\",\"name\":\"_claim\",\"type\":\"bytes32\"}],\"name\":\"attack\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondManager\",\"outputs\":[{\"internalType\":\"contractIBondManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimData\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"parentIndex\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"countered\",\"type\":\"bool\"},{\"internalType\":\"Claim\",\"name\":\"claim\",\"type\":\"bytes32\"},{\"internalType\":\"Position\",\"name\":\"position\",\"type\":\"uint128\"},{\"internalType\":\"Clock\",\"name\":\"clock\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimDataLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"len_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createdAt\",\"outputs\":[{\"internalType\":\"Timestamp\",\"name\":\"createdAt_\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_parentIndex\",\"type\":\"uint256\"},{\"internalType\":\"Claim\",\"name\":\"_claim\",\"type\":\"bytes32\"}],\"name\":\"defend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"extraData_\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraDataLen\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"extraDataLen_\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameData\",\"outputs\":[{\"internalType\":\"GameType\",\"name\":\"gameType_\",\"type\":\"uint8\"},{\"internalType\":\"Claim\",\"name\":\"rootClaim_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraData_\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameStart\",\"outputs\":[{\"internalType\":\"Timestamp\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameType\",\"outputs\":[{\"internalType\":\"GameType\",\"name\":\"gameType_\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"l2BlockNumber_\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_challengeIndex\",\"type\":\"uint256\"},{\"internalType\":\"Claim\",\"name\":\"_claim\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_isAttack\",\"type\":\"bool\"}],\"name\":\"move\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"enumGameStatus\",\"name\":\"status_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootClaim\",\"outputs\":[{\"internalType\":\"Claim\",\"name\":\"rootClaim_\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"enumGameStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_claimIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isAttack\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_stateData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"step\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61014060405234801561001157600080fd5b506040516120c03803806120c08339810160408190526100309161005b565b6000608081905260a052600260c05260e092909252610100526001600160a01b0316610120526100a1565b60008060006060848603121561007057600080fd5b83516020850151604086015191945092506001600160a01b038116811461009657600080fd5b809150509250925092565b60805160a05160c05160e0516101005161012051611fad610113600039600081816103ed015261147c0152600081816103050152818161065f01528181610be5015261124a0152600081816102190152611384015260006109a50152600061097c015260006109530152611fad6000f3fe6080604052600436106101755760003560e01c80638129fc1c116100cb578063bcef3b551161007f578063cf09e0d011610059578063cf09e0d0146104df578063d8cc1a3c146104fe578063fa24f7431461051e57600080fd5b8063bcef3b551461042b578063c55cd0c714610468578063c6f0308c1461047b57600080fd5b80638b85902b116100b05780638b85902b1461039b57806392931298146103db578063bbdc02db1461040f57600080fd5b80638129fc1c146103715780638980e0cc1461038657600080fd5b806335fef5671161012d57806354fd4d501161010757806354fd4d5014610327578063609d333414610349578063632247ea1461035e57600080fd5b806335fef5671461027f578063363cc427146102945780634778efe8146102f357600080fd5b8063266198f91161015e578063266198f9146102075780632810e1d6146102495780633218b99d1461025e57600080fd5b80631f62d0c61461017a578063200d2ed2146101d4575b600080fd5b34801561018657600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003602001355b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101e057600080fd5b506000546101fa9068010000000000000000900460ff1681565b6040516101cb9190611a57565b34801561021357600080fd5b5061023b7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016101cb565b34801561025557600080fd5b506101fa610547565b34801561026a57600080fd5b506000546101b69067ffffffffffffffff1681565b61029261028d366004611a98565b61093c565b005b3480156102a057600080fd5b506000546102ce906901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101cb565b3480156102ff57600080fd5b5061023b7f000000000000000000000000000000000000000000000000000000000000000081565b34801561033357600080fd5b5061033c61094c565b6040516101cb9190611b34565b34801561035557600080fd5b5061033c6109ef565b61029261036c366004611b63565b610a29565b34801561037d57600080fd5b50610292610fe3565b34801561039257600080fd5b5060015461023b565b3480156103a757600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036020013561023b565b3480156103e757600080fd5b506102ce7f000000000000000000000000000000000000000000000000000000000000000081565b34801561041b57600080fd5b50604051600081526020016101cb565b34801561043757600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033561023b565b610292610476366004611a98565b611124565b34801561048757600080fd5b5061049b610496366004611b98565b611130565b6040805163ffffffff90961686529315156020860152928401919091526fffffffffffffffffffffffffffffffff908116606084015216608082015260a0016101cb565b3480156104eb57600080fd5b5060005467ffffffffffffffff166101b6565b34801561050a57600080fd5b50610292610519366004611bfa565b6111a1565b34801561052a57600080fd5b5061053361156b565b6040516101cb93929190611c84565b905090565b60008060005468010000000000000000900460ff16600281111561056d5761056d611a28565b146105a4576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180546000916105b491611cde565b90506fffffffffffffffffffffffffffffffff815b67ffffffffffffffff81101561069e576000600182815481106105ee576105ee611cf5565b6000918252602090912060039091020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9093019290915060ff640100000000909104161561063f57506105c9565b6002810154600090610683906fffffffffffffffffffffffffffffffff167f00000000000000000000000000000000000000000000000000000000000000006115a9565b905083811015610697578093508260010194505b50506105c9565b506000600183815481106106b4576106b4611cf5565b600091825260208220600390910201805490925063ffffffff9081169190821461071e57600182815481106106eb576106eb611cf5565b906000526020600020906003020160020160109054906101000a90046fffffffffffffffffffffffffffffffff1661074a565b600283015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff165b905062049d4061076467ffffffffffffffff831642611cde565b610780836fffffffffffffffffffffffffffffffff1660401c90565b67ffffffffffffffff166107949190611d24565b116107cb576040517ff2440b5300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028381015461086d906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b6108779190611d6b565b67ffffffffffffffff1615801561089e57506fffffffffffffffffffffffffffffffff8414155b156108ac57600295506108b1565b600195505b600080548791907fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000008360028111156108f6576108f6611a28565b02179055600281111561090b5761090b611a28565b6040517f5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da6090600090a2505050505090565b61094882826000610a29565b5050565b60606109777f000000000000000000000000000000000000000000000000000000000000000061165e565b6109a07f000000000000000000000000000000000000000000000000000000000000000061165e565b6109c97f000000000000000000000000000000000000000000000000000000000000000061165e565b6040516020016109db93929190611d92565b604051602081830303815290604052905090565b60606105426040367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036020013561179b565b6000805468010000000000000000900460ff166002811115610a4d57610a4d611a28565b14610a84576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82158015610a90575080155b15610ac7576040517fa42637bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060018481548110610adc57610adc611cf5565b60009182526020918290206040805160a0810182526003909302909101805463ffffffff8116845260ff64010000000090910416151593830193909352600180840154918301919091526002909201546fffffffffffffffffffffffffffffffff80821660608401527001000000000000000000000000000000009091041660808201528154909250819086908110610b7757610b77611cf5565b6000918252602082206003909102018054921515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff909316929092179091556060820151610be1906fffffffffffffffffffffffffffffffff1684151760011b90565b90507f0000000000000000000000000000000000000000000000000000000000000000610ca0826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff161115610ce2576040517f56f57b2b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160009063ffffffff90811614610d42576001836000015163ffffffff1681548110610d1157610d11611cf5565b906000526020600020906003020160020160109054906101000a90046fffffffffffffffffffffffffffffffff1690505b608083015160009067ffffffffffffffff1667ffffffffffffffff1642610d7b846fffffffffffffffffffffffffffffffff1660401c90565b67ffffffffffffffff16610d8f9190611d24565b610d999190611cde565b905062049d4067ffffffffffffffff82161115610de2576040517f3381d11400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000604082901b421790506000610e03888660009182526020526040902090565b60008181526002602052604090205490915060ff1615610e4f576040517f80497e3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260026020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155815160a08101835263ffffffff808f1682529381018581528184018e81526fffffffffffffffffffffffffffffffff808d16606085019081528a82166080860190815286548088018855968a52945160039096027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf68101805495511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000909616979099169690961793909317909655517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf78401555190518416700100000000000000000000000000000000029316929092177fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf8909201919091555133918a918c917f9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be91a4505050505050505050565b600080547fffffffffffffffffffffffffffffffffffffffffffffff000000000000000000164267ffffffffffffffff161781556040805160a08101825263ffffffff815260208101929092526001919081016110687ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe369081013560f01c90033590565b815260016020820152604001426fffffffffffffffffffffffffffffffff908116909152825460018181018555600094855260209485902084516003909302018054958501511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000090961663ffffffff909316929092179490941781556040830151938101939093556060820151608090920151811670010000000000000000000000000000000002911617600290910155565b61094882826001610a29565b6001818154811061114057600080fd5b600091825260209091206003909102018054600182015460029092015463ffffffff8216935064010000000090910460ff1691906fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041685565b6000805468010000000000000000900460ff1660028111156111c5576111c5611a28565b146111fc576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006001878154811061121157611211611cf5565b6000918252602082206003919091020160028101549092506fffffffffffffffffffffffffffffffff16908715821760011b90506112707f00000000000000000000000000000000000000000000000000000000000000006001611d24565b61130c826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff161461134d576040517f5f53dd9800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008089156113d357611371836fffffffffffffffffffffffffffffffff16611832565b67ffffffffffffffff166000036113aa577f000000000000000000000000000000000000000000000000000000000000000091506113c8565b6113c56113b8600186611e08565b865463ffffffff166118d8565b91505b5060018401546113ed565b846001015491506113ea8460016113b89190611e39565b90505b8189896040516113fe929190611e6d565b60405180910390201461143d576040517f696550ff00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517ff8e0cb96000000000000000000000000000000000000000000000000000000008152819073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063f8e0cb96906114b7908d908d908d908d90600401611ec6565b6020604051808303816000875af11580156114d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114fa9190611ef8565b03611531576040517ffb4e40dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505082547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff16640100000000179092555050505050505050565b6000367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033560606115a26109ef565b9050909192565b600080611636847e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1690508083036001841b600180831b0386831b17039250505092915050565b6060816000036116a157505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156116cb57806116b581611f11565b91506116c49050600a83611f49565b91506116a5565b60008167ffffffffffffffff8111156116e6576116e6611f5d565b6040519080825280601f01601f191660200182016040528015611710576020820181803683370190505b5090505b841561179357611725600183611cde565b9150611732600a86611f8c565b61173d906030611d24565b60f81b81838151811061175257611752611cf5565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061178c600a86611f49565b9450611714565b949350505050565b606060006117d284367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003611d24565b90508267ffffffffffffffff1667ffffffffffffffff8111156117f7576117f7611f5d565b6040519080825280601f01601f191660200182016040528015611821576020820181803683370190505b509150828160208401375092915050565b6000806118bf837e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b600167ffffffffffffffff919091161b90920392915050565b6000806118f6846fffffffffffffffffffffffffffffffff1661197c565b905060006001848154811061190d5761190d611cf5565b906000526020600020906003020190505b60028101546fffffffffffffffffffffffffffffffff83811691161461197057805460018054909163ffffffff1690811061195b5761195b611cf5565b9060005260206000209060030201905061191e565b60010154949350505050565b60008119600183011681611a10827e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169390931c8015179392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160038310611a92577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60008060408385031215611aab57600080fd5b50508035926020909101359150565b60005b83811015611ad5578181015183820152602001611abd565b83811115611ae4576000848401525b50505050565b60008151808452611b02816020860160208601611aba565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611b476020830184611aea565b9392505050565b80358015158114611b5e57600080fd5b919050565b600080600060608486031215611b7857600080fd5b8335925060208401359150611b8f60408501611b4e565b90509250925092565b600060208284031215611baa57600080fd5b5035919050565b60008083601f840112611bc357600080fd5b50813567ffffffffffffffff811115611bdb57600080fd5b602083019150836020828501011115611bf357600080fd5b9250929050565b60008060008060008060808789031215611c1357600080fd5b86359550611c2360208801611b4e565b9450604087013567ffffffffffffffff80821115611c4057600080fd5b611c4c8a838b01611bb1565b90965094506060890135915080821115611c6557600080fd5b50611c7289828a01611bb1565b979a9699509497509295939492505050565b60ff84168152826020820152606060408201526000611ca66060830184611aea565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611cf057611cf0611caf565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008219821115611d3757611d37611caf565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff80841680611d8657611d86611d3c565b92169190910692915050565b60008451611da4818460208901611aba565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611de0816001850160208a01611aba565b60019201918201528351611dfb816002840160208801611aba565b0160020195945050505050565b60006fffffffffffffffffffffffffffffffff83811690831681811015611e3157611e31611caf565b039392505050565b60006fffffffffffffffffffffffffffffffff808316818516808303821115611e6457611e64611caf565b01949350505050565b8183823760009101908152919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000611eda604083018688611e7d565b8281036020840152611eed818587611e7d565b979650505050505050565b600060208284031215611f0a57600080fd5b5051919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611f4257611f42611caf565b5060010190565b600082611f5857611f58611d3c565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082611f9b57611f9b611d3c565b50069056fea164736f6c634300080f000a",
}

// FaultDisputeGameABI is the input ABI used to generate the binding from.
// Deprecated: Use FaultDisputeGameMetaData.ABI instead.
var FaultDisputeGameABI = FaultDisputeGameMetaData.ABI

// FaultDisputeGameBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FaultDisputeGameMetaData.Bin instead.
var FaultDisputeGameBin = FaultDisputeGameMetaData.Bin

// DeployFaultDisputeGame deploys a new Ethereum contract, binding an instance of FaultDisputeGame to it.
func DeployFaultDisputeGame(auth *bind.TransactOpts, backend bind.ContractBackend, _absolutePrestate [32]byte, _maxGameDepth *big.Int, _vm common.Address) (common.Address, *types.Transaction, *FaultDisputeGame, error) {
	parsed, err := FaultDisputeGameMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FaultDisputeGameBin), backend, _absolutePrestate, _maxGameDepth, _vm)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FaultDisputeGame{FaultDisputeGameCaller: FaultDisputeGameCaller{contract: contract}, FaultDisputeGameTransactor: FaultDisputeGameTransactor{contract: contract}, FaultDisputeGameFilterer: FaultDisputeGameFilterer{contract: contract}}, nil
}

// FaultDisputeGame is an auto generated Go binding around an Ethereum contract.
type FaultDisputeGame struct {
	FaultDisputeGameCaller     // Read-only binding to the contract
	FaultDisputeGameTransactor // Write-only binding to the contract
	FaultDisputeGameFilterer   // Log filterer for contract events
}

// FaultDisputeGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type FaultDisputeGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaultDisputeGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FaultDisputeGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaultDisputeGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FaultDisputeGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaultDisputeGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FaultDisputeGameSession struct {
	Contract     *FaultDisputeGame // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaultDisputeGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FaultDisputeGameCallerSession struct {
	Contract *FaultDisputeGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FaultDisputeGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FaultDisputeGameTransactorSession struct {
	Contract     *FaultDisputeGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FaultDisputeGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type FaultDisputeGameRaw struct {
	Contract *FaultDisputeGame // Generic contract binding to access the raw methods on
}

// FaultDisputeGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FaultDisputeGameCallerRaw struct {
	Contract *FaultDisputeGameCaller // Generic read-only contract binding to access the raw methods on
}

// FaultDisputeGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FaultDisputeGameTransactorRaw struct {
	Contract *FaultDisputeGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaultDisputeGame creates a new instance of FaultDisputeGame, bound to a specific deployed contract.
func NewFaultDisputeGame(address common.Address, backend bind.ContractBackend) (*FaultDisputeGame, error) {
	contract, err := bindFaultDisputeGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FaultDisputeGame{FaultDisputeGameCaller: FaultDisputeGameCaller{contract: contract}, FaultDisputeGameTransactor: FaultDisputeGameTransactor{contract: contract}, FaultDisputeGameFilterer: FaultDisputeGameFilterer{contract: contract}}, nil
}

// NewFaultDisputeGameCaller creates a new read-only instance of FaultDisputeGame, bound to a specific deployed contract.
func NewFaultDisputeGameCaller(address common.Address, caller bind.ContractCaller) (*FaultDisputeGameCaller, error) {
	contract, err := bindFaultDisputeGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaultDisputeGameCaller{contract: contract}, nil
}

// NewFaultDisputeGameTransactor creates a new write-only instance of FaultDisputeGame, bound to a specific deployed contract.
func NewFaultDisputeGameTransactor(address common.Address, transactor bind.ContractTransactor) (*FaultDisputeGameTransactor, error) {
	contract, err := bindFaultDisputeGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaultDisputeGameTransactor{contract: contract}, nil
}

// NewFaultDisputeGameFilterer creates a new log filterer instance of FaultDisputeGame, bound to a specific deployed contract.
func NewFaultDisputeGameFilterer(address common.Address, filterer bind.ContractFilterer) (*FaultDisputeGameFilterer, error) {
	contract, err := bindFaultDisputeGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaultDisputeGameFilterer{contract: contract}, nil
}

// bindFaultDisputeGame binds a generic wrapper to an already deployed contract.
func bindFaultDisputeGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FaultDisputeGameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaultDisputeGame *FaultDisputeGameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaultDisputeGame.Contract.FaultDisputeGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaultDisputeGame *FaultDisputeGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.FaultDisputeGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaultDisputeGame *FaultDisputeGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.FaultDisputeGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaultDisputeGame *FaultDisputeGameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaultDisputeGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaultDisputeGame *FaultDisputeGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaultDisputeGame *FaultDisputeGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.contract.Transact(opts, method, params...)
}

// ABSOLUTEPRESTATE is a free data retrieval call binding the contract method 0x266198f9.
//
// Solidity: function ABSOLUTE_PRESTATE() view returns(bytes32)
func (_FaultDisputeGame *FaultDisputeGameCaller) ABSOLUTEPRESTATE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "ABSOLUTE_PRESTATE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ABSOLUTEPRESTATE is a free data retrieval call binding the contract method 0x266198f9.
//
// Solidity: function ABSOLUTE_PRESTATE() view returns(bytes32)
func (_FaultDisputeGame *FaultDisputeGameSession) ABSOLUTEPRESTATE() ([32]byte, error) {
	return _FaultDisputeGame.Contract.ABSOLUTEPRESTATE(&_FaultDisputeGame.CallOpts)
}

// ABSOLUTEPRESTATE is a free data retrieval call binding the contract method 0x266198f9.
//
// Solidity: function ABSOLUTE_PRESTATE() view returns(bytes32)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) ABSOLUTEPRESTATE() ([32]byte, error) {
	return _FaultDisputeGame.Contract.ABSOLUTEPRESTATE(&_FaultDisputeGame.CallOpts)
}

// MAXGAMEDEPTH is a free data retrieval call binding the contract method 0x4778efe8.
//
// Solidity: function MAX_GAME_DEPTH() view returns(uint256)
func (_FaultDisputeGame *FaultDisputeGameCaller) MAXGAMEDEPTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "MAX_GAME_DEPTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXGAMEDEPTH is a free data retrieval call binding the contract method 0x4778efe8.
//
// Solidity: function MAX_GAME_DEPTH() view returns(uint256)
func (_FaultDisputeGame *FaultDisputeGameSession) MAXGAMEDEPTH() (*big.Int, error) {
	return _FaultDisputeGame.Contract.MAXGAMEDEPTH(&_FaultDisputeGame.CallOpts)
}

// MAXGAMEDEPTH is a free data retrieval call binding the contract method 0x4778efe8.
//
// Solidity: function MAX_GAME_DEPTH() view returns(uint256)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) MAXGAMEDEPTH() (*big.Int, error) {
	return _FaultDisputeGame.Contract.MAXGAMEDEPTH(&_FaultDisputeGame.CallOpts)
}

// VM is a free data retrieval call binding the contract method 0x92931298.
//
// Solidity: function VM() view returns(address)
func (_FaultDisputeGame *FaultDisputeGameCaller) VM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "VM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VM is a free data retrieval call binding the contract method 0x92931298.
//
// Solidity: function VM() view returns(address)
func (_FaultDisputeGame *FaultDisputeGameSession) VM() (common.Address, error) {
	return _FaultDisputeGame.Contract.VM(&_FaultDisputeGame.CallOpts)
}

// VM is a free data retrieval call binding the contract method 0x92931298.
//
// Solidity: function VM() view returns(address)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) VM() (common.Address, error) {
	return _FaultDisputeGame.Contract.VM(&_FaultDisputeGame.CallOpts)
}

// BondManager is a free data retrieval call binding the contract method 0x363cc427.
//
// Solidity: function bondManager() view returns(address)
func (_FaultDisputeGame *FaultDisputeGameCaller) BondManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "bondManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BondManager is a free data retrieval call binding the contract method 0x363cc427.
//
// Solidity: function bondManager() view returns(address)
func (_FaultDisputeGame *FaultDisputeGameSession) BondManager() (common.Address, error) {
	return _FaultDisputeGame.Contract.BondManager(&_FaultDisputeGame.CallOpts)
}

// BondManager is a free data retrieval call binding the contract method 0x363cc427.
//
// Solidity: function bondManager() view returns(address)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) BondManager() (common.Address, error) {
	return _FaultDisputeGame.Contract.BondManager(&_FaultDisputeGame.CallOpts)
}

// ClaimData is a free data retrieval call binding the contract method 0xc6f0308c.
//
// Solidity: function claimData(uint256 ) view returns(uint32 parentIndex, bool countered, bytes32 claim, uint128 position, uint128 clock)
func (_FaultDisputeGame *FaultDisputeGameCaller) ClaimData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ParentIndex uint32
	Countered   bool
	Claim       [32]byte
	Position    *big.Int
	Clock       *big.Int
}, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "claimData", arg0)

	outstruct := new(struct {
		ParentIndex uint32
		Countered   bool
		Claim       [32]byte
		Position    *big.Int
		Clock       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ParentIndex = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Countered = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Claim = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.Position = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Clock = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ClaimData is a free data retrieval call binding the contract method 0xc6f0308c.
//
// Solidity: function claimData(uint256 ) view returns(uint32 parentIndex, bool countered, bytes32 claim, uint128 position, uint128 clock)
func (_FaultDisputeGame *FaultDisputeGameSession) ClaimData(arg0 *big.Int) (struct {
	ParentIndex uint32
	Countered   bool
	Claim       [32]byte
	Position    *big.Int
	Clock       *big.Int
}, error) {
	return _FaultDisputeGame.Contract.ClaimData(&_FaultDisputeGame.CallOpts, arg0)
}

// ClaimData is a free data retrieval call binding the contract method 0xc6f0308c.
//
// Solidity: function claimData(uint256 ) view returns(uint32 parentIndex, bool countered, bytes32 claim, uint128 position, uint128 clock)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) ClaimData(arg0 *big.Int) (struct {
	ParentIndex uint32
	Countered   bool
	Claim       [32]byte
	Position    *big.Int
	Clock       *big.Int
}, error) {
	return _FaultDisputeGame.Contract.ClaimData(&_FaultDisputeGame.CallOpts, arg0)
}

// ClaimDataLen is a free data retrieval call binding the contract method 0x8980e0cc.
//
// Solidity: function claimDataLen() view returns(uint256 len_)
func (_FaultDisputeGame *FaultDisputeGameCaller) ClaimDataLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "claimDataLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimDataLen is a free data retrieval call binding the contract method 0x8980e0cc.
//
// Solidity: function claimDataLen() view returns(uint256 len_)
func (_FaultDisputeGame *FaultDisputeGameSession) ClaimDataLen() (*big.Int, error) {
	return _FaultDisputeGame.Contract.ClaimDataLen(&_FaultDisputeGame.CallOpts)
}

// ClaimDataLen is a free data retrieval call binding the contract method 0x8980e0cc.
//
// Solidity: function claimDataLen() view returns(uint256 len_)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) ClaimDataLen() (*big.Int, error) {
	return _FaultDisputeGame.Contract.ClaimDataLen(&_FaultDisputeGame.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint64 createdAt_)
func (_FaultDisputeGame *FaultDisputeGameCaller) CreatedAt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "createdAt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint64 createdAt_)
func (_FaultDisputeGame *FaultDisputeGameSession) CreatedAt() (uint64, error) {
	return _FaultDisputeGame.Contract.CreatedAt(&_FaultDisputeGame.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint64 createdAt_)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) CreatedAt() (uint64, error) {
	return _FaultDisputeGame.Contract.CreatedAt(&_FaultDisputeGame.CallOpts)
}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() pure returns(bytes extraData_)
func (_FaultDisputeGame *FaultDisputeGameCaller) ExtraData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "extraData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() pure returns(bytes extraData_)
func (_FaultDisputeGame *FaultDisputeGameSession) ExtraData() ([]byte, error) {
	return _FaultDisputeGame.Contract.ExtraData(&_FaultDisputeGame.CallOpts)
}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() pure returns(bytes extraData_)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) ExtraData() ([]byte, error) {
	return _FaultDisputeGame.Contract.ExtraData(&_FaultDisputeGame.CallOpts)
}

// ExtraDataLen is a free data retrieval call binding the contract method 0x1f62d0c6.
//
// Solidity: function extraDataLen() pure returns(uint64 extraDataLen_)
func (_FaultDisputeGame *FaultDisputeGameCaller) ExtraDataLen(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "extraDataLen")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ExtraDataLen is a free data retrieval call binding the contract method 0x1f62d0c6.
//
// Solidity: function extraDataLen() pure returns(uint64 extraDataLen_)
func (_FaultDisputeGame *FaultDisputeGameSession) ExtraDataLen() (uint64, error) {
	return _FaultDisputeGame.Contract.ExtraDataLen(&_FaultDisputeGame.CallOpts)
}

// ExtraDataLen is a free data retrieval call binding the contract method 0x1f62d0c6.
//
// Solidity: function extraDataLen() pure returns(uint64 extraDataLen_)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) ExtraDataLen() (uint64, error) {
	return _FaultDisputeGame.Contract.ExtraDataLen(&_FaultDisputeGame.CallOpts)
}

// GameData is a free data retrieval call binding the contract method 0xfa24f743.
//
// Solidity: function gameData() pure returns(uint8 gameType_, bytes32 rootClaim_, bytes extraData_)
func (_FaultDisputeGame *FaultDisputeGameCaller) GameData(opts *bind.CallOpts) (struct {
	GameType  uint8
	RootClaim [32]byte
	ExtraData []byte
}, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "gameData")

	outstruct := new(struct {
		GameType  uint8
		RootClaim [32]byte
		ExtraData []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GameType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.RootClaim = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ExtraData = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GameData is a free data retrieval call binding the contract method 0xfa24f743.
//
// Solidity: function gameData() pure returns(uint8 gameType_, bytes32 rootClaim_, bytes extraData_)
func (_FaultDisputeGame *FaultDisputeGameSession) GameData() (struct {
	GameType  uint8
	RootClaim [32]byte
	ExtraData []byte
}, error) {
	return _FaultDisputeGame.Contract.GameData(&_FaultDisputeGame.CallOpts)
}

// GameData is a free data retrieval call binding the contract method 0xfa24f743.
//
// Solidity: function gameData() pure returns(uint8 gameType_, bytes32 rootClaim_, bytes extraData_)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) GameData() (struct {
	GameType  uint8
	RootClaim [32]byte
	ExtraData []byte
}, error) {
	return _FaultDisputeGame.Contract.GameData(&_FaultDisputeGame.CallOpts)
}

// GameStart is a free data retrieval call binding the contract method 0x3218b99d.
//
// Solidity: function gameStart() view returns(uint64)
func (_FaultDisputeGame *FaultDisputeGameCaller) GameStart(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "gameStart")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GameStart is a free data retrieval call binding the contract method 0x3218b99d.
//
// Solidity: function gameStart() view returns(uint64)
func (_FaultDisputeGame *FaultDisputeGameSession) GameStart() (uint64, error) {
	return _FaultDisputeGame.Contract.GameStart(&_FaultDisputeGame.CallOpts)
}

// GameStart is a free data retrieval call binding the contract method 0x3218b99d.
//
// Solidity: function gameStart() view returns(uint64)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) GameStart() (uint64, error) {
	return _FaultDisputeGame.Contract.GameStart(&_FaultDisputeGame.CallOpts)
}

// GameType is a free data retrieval call binding the contract method 0xbbdc02db.
//
// Solidity: function gameType() pure returns(uint8 gameType_)
func (_FaultDisputeGame *FaultDisputeGameCaller) GameType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "gameType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GameType is a free data retrieval call binding the contract method 0xbbdc02db.
//
// Solidity: function gameType() pure returns(uint8 gameType_)
func (_FaultDisputeGame *FaultDisputeGameSession) GameType() (uint8, error) {
	return _FaultDisputeGame.Contract.GameType(&_FaultDisputeGame.CallOpts)
}

// GameType is a free data retrieval call binding the contract method 0xbbdc02db.
//
// Solidity: function gameType() pure returns(uint8 gameType_)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) GameType() (uint8, error) {
	return _FaultDisputeGame.Contract.GameType(&_FaultDisputeGame.CallOpts)
}

// L2BlockNumber is a free data retrieval call binding the contract method 0x8b85902b.
//
// Solidity: function l2BlockNumber() pure returns(uint256 l2BlockNumber_)
func (_FaultDisputeGame *FaultDisputeGameCaller) L2BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "l2BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BlockNumber is a free data retrieval call binding the contract method 0x8b85902b.
//
// Solidity: function l2BlockNumber() pure returns(uint256 l2BlockNumber_)
func (_FaultDisputeGame *FaultDisputeGameSession) L2BlockNumber() (*big.Int, error) {
	return _FaultDisputeGame.Contract.L2BlockNumber(&_FaultDisputeGame.CallOpts)
}

// L2BlockNumber is a free data retrieval call binding the contract method 0x8b85902b.
//
// Solidity: function l2BlockNumber() pure returns(uint256 l2BlockNumber_)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) L2BlockNumber() (*big.Int, error) {
	return _FaultDisputeGame.Contract.L2BlockNumber(&_FaultDisputeGame.CallOpts)
}

// RootClaim is a free data retrieval call binding the contract method 0xbcef3b55.
//
// Solidity: function rootClaim() pure returns(bytes32 rootClaim_)
func (_FaultDisputeGame *FaultDisputeGameCaller) RootClaim(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "rootClaim")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RootClaim is a free data retrieval call binding the contract method 0xbcef3b55.
//
// Solidity: function rootClaim() pure returns(bytes32 rootClaim_)
func (_FaultDisputeGame *FaultDisputeGameSession) RootClaim() ([32]byte, error) {
	return _FaultDisputeGame.Contract.RootClaim(&_FaultDisputeGame.CallOpts)
}

// RootClaim is a free data retrieval call binding the contract method 0xbcef3b55.
//
// Solidity: function rootClaim() pure returns(bytes32 rootClaim_)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) RootClaim() ([32]byte, error) {
	return _FaultDisputeGame.Contract.RootClaim(&_FaultDisputeGame.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_FaultDisputeGame *FaultDisputeGameCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_FaultDisputeGame *FaultDisputeGameSession) Status() (uint8, error) {
	return _FaultDisputeGame.Contract.Status(&_FaultDisputeGame.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) Status() (uint8, error) {
	return _FaultDisputeGame.Contract.Status(&_FaultDisputeGame.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_FaultDisputeGame *FaultDisputeGameCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FaultDisputeGame.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_FaultDisputeGame *FaultDisputeGameSession) Version() (string, error) {
	return _FaultDisputeGame.Contract.Version(&_FaultDisputeGame.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_FaultDisputeGame *FaultDisputeGameCallerSession) Version() (string, error) {
	return _FaultDisputeGame.Contract.Version(&_FaultDisputeGame.CallOpts)
}

// Attack is a paid mutator transaction binding the contract method 0xc55cd0c7.
//
// Solidity: function attack(uint256 _parentIndex, bytes32 _claim) payable returns()
func (_FaultDisputeGame *FaultDisputeGameTransactor) Attack(opts *bind.TransactOpts, _parentIndex *big.Int, _claim [32]byte) (*types.Transaction, error) {
	return _FaultDisputeGame.contract.Transact(opts, "attack", _parentIndex, _claim)
}

// Attack is a paid mutator transaction binding the contract method 0xc55cd0c7.
//
// Solidity: function attack(uint256 _parentIndex, bytes32 _claim) payable returns()
func (_FaultDisputeGame *FaultDisputeGameSession) Attack(_parentIndex *big.Int, _claim [32]byte) (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.Attack(&_FaultDisputeGame.TransactOpts, _parentIndex, _claim)
}

// Attack is a paid mutator transaction binding the contract method 0xc55cd0c7.
//
// Solidity: function attack(uint256 _parentIndex, bytes32 _claim) payable returns()
func (_FaultDisputeGame *FaultDisputeGameTransactorSession) Attack(_parentIndex *big.Int, _claim [32]byte) (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.Attack(&_FaultDisputeGame.TransactOpts, _parentIndex, _claim)
}

// Defend is a paid mutator transaction binding the contract method 0x35fef567.
//
// Solidity: function defend(uint256 _parentIndex, bytes32 _claim) payable returns()
func (_FaultDisputeGame *FaultDisputeGameTransactor) Defend(opts *bind.TransactOpts, _parentIndex *big.Int, _claim [32]byte) (*types.Transaction, error) {
	return _FaultDisputeGame.contract.Transact(opts, "defend", _parentIndex, _claim)
}

// Defend is a paid mutator transaction binding the contract method 0x35fef567.
//
// Solidity: function defend(uint256 _parentIndex, bytes32 _claim) payable returns()
func (_FaultDisputeGame *FaultDisputeGameSession) Defend(_parentIndex *big.Int, _claim [32]byte) (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.Defend(&_FaultDisputeGame.TransactOpts, _parentIndex, _claim)
}

// Defend is a paid mutator transaction binding the contract method 0x35fef567.
//
// Solidity: function defend(uint256 _parentIndex, bytes32 _claim) payable returns()
func (_FaultDisputeGame *FaultDisputeGameTransactorSession) Defend(_parentIndex *big.Int, _claim [32]byte) (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.Defend(&_FaultDisputeGame.TransactOpts, _parentIndex, _claim)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FaultDisputeGame *FaultDisputeGameTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaultDisputeGame.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FaultDisputeGame *FaultDisputeGameSession) Initialize() (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.Initialize(&_FaultDisputeGame.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FaultDisputeGame *FaultDisputeGameTransactorSession) Initialize() (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.Initialize(&_FaultDisputeGame.TransactOpts)
}

// Move is a paid mutator transaction binding the contract method 0x632247ea.
//
// Solidity: function move(uint256 _challengeIndex, bytes32 _claim, bool _isAttack) payable returns()
func (_FaultDisputeGame *FaultDisputeGameTransactor) Move(opts *bind.TransactOpts, _challengeIndex *big.Int, _claim [32]byte, _isAttack bool) (*types.Transaction, error) {
	return _FaultDisputeGame.contract.Transact(opts, "move", _challengeIndex, _claim, _isAttack)
}

// Move is a paid mutator transaction binding the contract method 0x632247ea.
//
// Solidity: function move(uint256 _challengeIndex, bytes32 _claim, bool _isAttack) payable returns()
func (_FaultDisputeGame *FaultDisputeGameSession) Move(_challengeIndex *big.Int, _claim [32]byte, _isAttack bool) (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.Move(&_FaultDisputeGame.TransactOpts, _challengeIndex, _claim, _isAttack)
}

// Move is a paid mutator transaction binding the contract method 0x632247ea.
//
// Solidity: function move(uint256 _challengeIndex, bytes32 _claim, bool _isAttack) payable returns()
func (_FaultDisputeGame *FaultDisputeGameTransactorSession) Move(_challengeIndex *big.Int, _claim [32]byte, _isAttack bool) (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.Move(&_FaultDisputeGame.TransactOpts, _challengeIndex, _claim, _isAttack)
}

// Resolve is a paid mutator transaction binding the contract method 0x2810e1d6.
//
// Solidity: function resolve() returns(uint8 status_)
func (_FaultDisputeGame *FaultDisputeGameTransactor) Resolve(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaultDisputeGame.contract.Transact(opts, "resolve")
}

// Resolve is a paid mutator transaction binding the contract method 0x2810e1d6.
//
// Solidity: function resolve() returns(uint8 status_)
func (_FaultDisputeGame *FaultDisputeGameSession) Resolve() (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.Resolve(&_FaultDisputeGame.TransactOpts)
}

// Resolve is a paid mutator transaction binding the contract method 0x2810e1d6.
//
// Solidity: function resolve() returns(uint8 status_)
func (_FaultDisputeGame *FaultDisputeGameTransactorSession) Resolve() (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.Resolve(&_FaultDisputeGame.TransactOpts)
}

// Step is a paid mutator transaction binding the contract method 0xd8cc1a3c.
//
// Solidity: function step(uint256 _claimIndex, bool _isAttack, bytes _stateData, bytes _proof) returns()
func (_FaultDisputeGame *FaultDisputeGameTransactor) Step(opts *bind.TransactOpts, _claimIndex *big.Int, _isAttack bool, _stateData []byte, _proof []byte) (*types.Transaction, error) {
	return _FaultDisputeGame.contract.Transact(opts, "step", _claimIndex, _isAttack, _stateData, _proof)
}

// Step is a paid mutator transaction binding the contract method 0xd8cc1a3c.
//
// Solidity: function step(uint256 _claimIndex, bool _isAttack, bytes _stateData, bytes _proof) returns()
func (_FaultDisputeGame *FaultDisputeGameSession) Step(_claimIndex *big.Int, _isAttack bool, _stateData []byte, _proof []byte) (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.Step(&_FaultDisputeGame.TransactOpts, _claimIndex, _isAttack, _stateData, _proof)
}

// Step is a paid mutator transaction binding the contract method 0xd8cc1a3c.
//
// Solidity: function step(uint256 _claimIndex, bool _isAttack, bytes _stateData, bytes _proof) returns()
func (_FaultDisputeGame *FaultDisputeGameTransactorSession) Step(_claimIndex *big.Int, _isAttack bool, _stateData []byte, _proof []byte) (*types.Transaction, error) {
	return _FaultDisputeGame.Contract.Step(&_FaultDisputeGame.TransactOpts, _claimIndex, _isAttack, _stateData, _proof)
}

// FaultDisputeGameMoveIterator is returned from FilterMove and is used to iterate over the raw logs and unpacked data for Move events raised by the FaultDisputeGame contract.
type FaultDisputeGameMoveIterator struct {
	Event *FaultDisputeGameMove // Event containing the contract specifics and raw log

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
func (it *FaultDisputeGameMoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaultDisputeGameMove)
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
		it.Event = new(FaultDisputeGameMove)
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
func (it *FaultDisputeGameMoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaultDisputeGameMoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaultDisputeGameMove represents a Move event raised by the FaultDisputeGame contract.
type FaultDisputeGameMove struct {
	ParentIndex *big.Int
	Claim       [32]byte
	Claimant    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMove is a free log retrieval operation binding the contract event 0x9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be.
//
// Solidity: event Move(uint256 indexed parentIndex, bytes32 indexed claim, address indexed claimant)
func (_FaultDisputeGame *FaultDisputeGameFilterer) FilterMove(opts *bind.FilterOpts, parentIndex []*big.Int, claim [][32]byte, claimant []common.Address) (*FaultDisputeGameMoveIterator, error) {

	var parentIndexRule []interface{}
	for _, parentIndexItem := range parentIndex {
		parentIndexRule = append(parentIndexRule, parentIndexItem)
	}
	var claimRule []interface{}
	for _, claimItem := range claim {
		claimRule = append(claimRule, claimItem)
	}
	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _FaultDisputeGame.contract.FilterLogs(opts, "Move", parentIndexRule, claimRule, claimantRule)
	if err != nil {
		return nil, err
	}
	return &FaultDisputeGameMoveIterator{contract: _FaultDisputeGame.contract, event: "Move", logs: logs, sub: sub}, nil
}

// WatchMove is a free log subscription operation binding the contract event 0x9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be.
//
// Solidity: event Move(uint256 indexed parentIndex, bytes32 indexed claim, address indexed claimant)
func (_FaultDisputeGame *FaultDisputeGameFilterer) WatchMove(opts *bind.WatchOpts, sink chan<- *FaultDisputeGameMove, parentIndex []*big.Int, claim [][32]byte, claimant []common.Address) (event.Subscription, error) {

	var parentIndexRule []interface{}
	for _, parentIndexItem := range parentIndex {
		parentIndexRule = append(parentIndexRule, parentIndexItem)
	}
	var claimRule []interface{}
	for _, claimItem := range claim {
		claimRule = append(claimRule, claimItem)
	}
	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _FaultDisputeGame.contract.WatchLogs(opts, "Move", parentIndexRule, claimRule, claimantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaultDisputeGameMove)
				if err := _FaultDisputeGame.contract.UnpackLog(event, "Move", log); err != nil {
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

// ParseMove is a log parse operation binding the contract event 0x9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be.
//
// Solidity: event Move(uint256 indexed parentIndex, bytes32 indexed claim, address indexed claimant)
func (_FaultDisputeGame *FaultDisputeGameFilterer) ParseMove(log types.Log) (*FaultDisputeGameMove, error) {
	event := new(FaultDisputeGameMove)
	if err := _FaultDisputeGame.contract.UnpackLog(event, "Move", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FaultDisputeGameResolvedIterator is returned from FilterResolved and is used to iterate over the raw logs and unpacked data for Resolved events raised by the FaultDisputeGame contract.
type FaultDisputeGameResolvedIterator struct {
	Event *FaultDisputeGameResolved // Event containing the contract specifics and raw log

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
func (it *FaultDisputeGameResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaultDisputeGameResolved)
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
		it.Event = new(FaultDisputeGameResolved)
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
func (it *FaultDisputeGameResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaultDisputeGameResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaultDisputeGameResolved represents a Resolved event raised by the FaultDisputeGame contract.
type FaultDisputeGameResolved struct {
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResolved is a free log retrieval operation binding the contract event 0x5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da60.
//
// Solidity: event Resolved(uint8 indexed status)
func (_FaultDisputeGame *FaultDisputeGameFilterer) FilterResolved(opts *bind.FilterOpts, status []uint8) (*FaultDisputeGameResolvedIterator, error) {

	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _FaultDisputeGame.contract.FilterLogs(opts, "Resolved", statusRule)
	if err != nil {
		return nil, err
	}
	return &FaultDisputeGameResolvedIterator{contract: _FaultDisputeGame.contract, event: "Resolved", logs: logs, sub: sub}, nil
}

// WatchResolved is a free log subscription operation binding the contract event 0x5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da60.
//
// Solidity: event Resolved(uint8 indexed status)
func (_FaultDisputeGame *FaultDisputeGameFilterer) WatchResolved(opts *bind.WatchOpts, sink chan<- *FaultDisputeGameResolved, status []uint8) (event.Subscription, error) {

	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _FaultDisputeGame.contract.WatchLogs(opts, "Resolved", statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaultDisputeGameResolved)
				if err := _FaultDisputeGame.contract.UnpackLog(event, "Resolved", log); err != nil {
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

// ParseResolved is a log parse operation binding the contract event 0x5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da60.
//
// Solidity: event Resolved(uint8 indexed status)
func (_FaultDisputeGame *FaultDisputeGameFilterer) ParseResolved(log types.Log) (*FaultDisputeGameResolved, error) {
	event := new(FaultDisputeGameResolved)
	if err := _FaultDisputeGame.contract.UnpackLog(event, "Resolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
