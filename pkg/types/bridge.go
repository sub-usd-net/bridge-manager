package types

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// CombinedBridgeStatus uses uint64 instead of *big.Int for easier manipulation
type CombinedBridgeStatus struct {
	CChainDepositId           uint64
	CChainCrossChainDepositId uint64

	SChainDepositId           uint64
	SChainCrossChainDepositId uint64
}

func (bs *CombinedBridgeStatus) ValidateInvariant() error {
	if bs.CChainCrossChainDepositId > bs.SChainDepositId {
		return errors.New("c-chain crossChainDepositId > s-chain depositId. did you redeploy contracts and only change 1 config")
	} else if bs.SChainCrossChainDepositId > bs.CChainDepositId {
		return errors.New("s-chain crossChainDepositId > c-chain depositId. did you redeploy contracts and only change 1 config")
	}
	return nil
}

func (bs *CombinedBridgeStatus) NeedsAction() bool {
	return bs.CChainDepositId != bs.SChainCrossChainDepositId || bs.SChainDepositId != bs.CChainCrossChainDepositId
}

type BridgeStatus struct {
	DepositId           *big.Int
	CrossChainDepositId *big.Int
}

type DepositInfo struct {
	User   common.Address
	Amount *big.Int
}

type BridgeReader interface {
	Status() (*BridgeStatus, error)
	DepositInfo(depositId *big.Int) (*DepositInfo, error)
}

type BridgeWriter interface {
	CompleteTransfer(depositId *big.Int, depositInfo *DepositInfo) (*types.Transaction, error)
}

type BridgeReaderWriter interface {
	BridgeReader
	BridgeWriter
	Client() *ethclient.Client
}

func CombineBridgeStatus(c *BridgeStatus, s *BridgeStatus) *CombinedBridgeStatus {
	return &CombinedBridgeStatus{
		CChainDepositId:           c.DepositId.Uint64(),
		CChainCrossChainDepositId: c.CrossChainDepositId.Uint64(),
		SChainDepositId:           s.DepositId.Uint64(),
		SChainCrossChainDepositId: s.CrossChainDepositId.Uint64(),
	}
}
