package fault

import (
	"context"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type FaultCaller struct {
	*bindings.FaultDisputeGameCaller

	log log.Logger
}

func NewFaultCaller(fdgAddr common.Address, client *ethclient.Client, log log.Logger) (*FaultCaller, error) {
	caller, err := bindings.NewFaultDisputeGameCaller(fdgAddr, client)
	if err != nil {
		return nil, err
	}
	return &FaultCaller{
		caller,
		log,
	}, nil
}

// LogGameInfo logs the game info.
func (fc *FaultCaller) LogGameInfo() {
	fc.LogGameStatus()
	fc.LogClaimDataLength()
}

// GetGameStatus returns the current game status.
// 0: In Progress
// 1: Challenger Won
// 2: Defender Won
func (fc *FaultCaller) GetGameStatus(ctx context.Context) (uint8, error) {
	return fc.Status(&bind.CallOpts{Context: ctx})
}

func (fc *FaultCaller) LogGameStatus() {
	status, err := fc.GetGameStatus(context.Background())
	if err != nil {
		fc.log.Error("failed to get game status", "err", err)
	}
	fc.log.Info("Game status", "status", GameStatusString(status))
}

// GetClaimDataLength returns the number of claims in the game.
func (fc *FaultCaller) GetClaimDataLength(ctx context.Context) (*big.Int, error) {
	return fc.ClaimDataLen(&bind.CallOpts{Context: ctx})
}

func (fc *FaultCaller) LogClaimDataLength() {
	claimLen, err := fc.GetClaimDataLength(context.Background())
	if err != nil {
		fc.log.Error("failed to get claim count", "err", err)
	}
	fc.log.Info("Number of claims", "length", claimLen)
}

// GameStatusString returns the current game status as a string.
func GameStatusString(status uint8) string {
	switch status {
	case 0:
		return "In Progress"
	case 1:
		return "Challenger Won"
	case 2:
		return "Defender Won"
	default:
		return "Unknown"
	}
}
