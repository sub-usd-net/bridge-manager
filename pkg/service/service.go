package service

import (
	"fmt"
	"math/big"
	"time"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sub-usd-net/bridge-manager/pkg/types"
	"go.uber.org/zap"
)

const (
	retryInterval = time.Second * 5
	pollInterval  = time.Second * 10
)

type BridgeService struct {
	c types.BridgeReaderWriter
	s types.BridgeReaderWriter

	cfg *types.BridgeConfig
	log *zap.SugaredLogger

	// cache at startup
	// used for translating cChain token amounts into sChain native token amounts
	// and vice versa
	stableTokenDecimals *big.Int
}

func NewBridgeService(cfg *types.BridgeConfig, log *zap.SugaredLogger) (*BridgeService, error) {
	cClient, err := ethclient.Dial(cfg.CChainRpcUri)
	if err != nil {
		return nil, err
	}
	sClient, err := ethclient.Dial(cfg.SubnetRpcUri)
	if err != nil {
		return nil, err
	}

	c, err := NewBridgeReaderWriter(cClient, cfg.CChainBridgeContractAddress, cfg.CChainBridgeAdminKeyPath)
	if err != nil {
		return nil, err
	}
	s, err := NewBridgeReaderWriter(sClient, cfg.SubnetBridgeContractAddress, cfg.SubnetBridgeAdminKeyPath)
	if err != nil {
		return nil, err
	}

	b := &BridgeService{
		c:   c,
		s:   s,
		cfg: cfg,
		log: log,
	}

	b.stableTokenDecimals, err = b.geStableTokenDecimals(cClient)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (b *BridgeService) Start() error {
	b.log.Infow("Starting Bridge-Manager")
	b.log.Infow("Bridge Details",
		"C-Chain Bridge", b.cfg.CChainBridgeContractAddress,
		"C-Chain Bridge Admin", b.c.(*readerWriter).txOpts.From,
		"Subnet Bridget", b.cfg.SubnetBridgeContractAddress,
		"Subnet Bridge Admin", b.s.(*readerWriter).txOpts.From,
	)

	for {
		status, err := b.getCurrentStatus()
		if err != nil {
			b.log.Warn("Error getting status. Will Retry")
			time.Sleep(retryInterval)
			continue
		}

		if err := status.ValidateInvariant(); err != nil {
			b.log.Errorw("Invariant failure. Requires operator action", "err", err)
			return err
		}

		if !status.NeedsAction() {
			time.Sleep(retryInterval)
			continue
		}

		// c-chain => subnet
		for id := status.SChainCrossChainDepositId + 1; id <= status.CChainDepositId; id++ {
			shouldAbort, err := b.doCompleteTransfer("C->Subnet", big.NewInt(int64(id)), b.s.Client(), b.c.DepositInfo, b.s.CompleteTransfer)
			if shouldAbort {
				return err
			} else if err != nil {
				time.Sleep(retryInterval)
			}
		}

		// subnet => c-chain
		for id := status.CChainCrossChainDepositId + 1; id < status.SChainDepositId; id++ {
			shouldAbort, err := b.doCompleteTransfer("Subnet->>C", big.NewInt(int64(id)), b.c.Client(), b.s.DepositInfo, b.c.CompleteTransfer)
			if shouldAbort {
				return err
			} else if err != nil {
				time.Sleep(retryInterval)
			}
		}

		time.Sleep(pollInterval)
	}
}

// doCompleteTransfer completes transfer C->Subnet or Subnet->C
// returns false along with the error in cases where we should abort all further execution
func (b *BridgeService) doCompleteTransfer(
	prefix string,
	depositId *big.Int,
	transferSideClient *ethclient.Client,
	getDepositInfo func(*big.Int) (*types.DepositInfo, error),
	completeTransfer func(*big.Int, *types.DepositInfo) (*ethtypes.Transaction, error)) (bool, error) {

	b.log.Infow(prefix, "depositId", depositId)
	info, err := getDepositInfo(depositId)
	if err != nil {
		b.log.Warnw(fmt.Sprintf("%s: Error reading deposit info", prefix), "err", err)
		return false, errRead
	}

	b.log.Infow(fmt.Sprintf("%s: Successfuly read deposit Info", prefix),
		"depositId", depositId, "user", info.User, "amount", info.Amount)
	if info.User == zeroAddress {
		return true, errInvariant
	}

	tx, err := completeTransfer(depositId, info)
	if err != nil {
		return false, err
	}
	b.log.Infow(fmt.Sprintf("%s: Submitted completeTransfer tx", prefix), "depositId", depositId, "hash", tx.Hash())

	ok, err := waitForTx(transferSideClient, tx)
	if err != nil {
		return false, err
	} else if !ok {
		b.log.Warnw("Transfer tx failure")
		return true, errTransfer
	}

	return false, nil
}
