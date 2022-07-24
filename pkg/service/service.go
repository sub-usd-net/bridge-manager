package service

import (
	"errors"
	"math/big"
	"time"

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
			if err := b.completeCChainToSubnetTransfer(big.NewInt(int64(id))); err != nil {
				if errors.Is(err, errRead) {
					time.Sleep(retryInterval)
				} else {
					return err
				}
			}
		}

		// subnet => c-chain
		for id := status.CChainCrossChainDepositId + 1; id < status.SChainDepositId; id++ {
			b.completeSubnetToCChainTransfer(big.NewInt(int64(id)))
		}

		time.Sleep(pollInterval)
	}
}

func (b *BridgeService) completeCChainToSubnetTransfer(depositId *big.Int) error {
	b.log.Infow("Bridging C-Chain => Subnet", "depositId", depositId)
	info, err := b.c.DepositInfo(depositId)
	if err != nil {
		b.log.Warnw("Error reading deposit info", "err", err)
		return errRead
	}
	b.log.Infow("Read Deposit Info",
		"depositId", depositId, "user", info.User, "amount", info.Amount)
	if info.User == zeroAddress {
		return errInvariant
	}
	tx, err := b.s.CompleteTransfer(depositId, info)
	if err != nil {
		return err
	}
	b.log.Infow("Submitted completeTransfer tx", "depositId", depositId, "hash", tx.Hash())
	return err
}

func (b *BridgeService) completeSubnetToCChainTransfer(depositId *big.Int) {
	b.log.Infow("Bridging Subnet => C-Chain", "depositId", depositId)
	// TODO
}
