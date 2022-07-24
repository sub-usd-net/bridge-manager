package service

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sub-usd-net/bridge-manager/pkg/bindings/bridge"
	"github.com/sub-usd-net/bridge-manager/pkg/key"
	"github.com/sub-usd-net/bridge-manager/pkg/types"
)

type readerWriter struct {
	bridge *bridge.Bridge
	txOpts *bind.TransactOpts
}

func NewBridgeReaderWriter(client *ethclient.Client, address common.Address, adminKeyPath string) (types.BridgeReaderWriter, error) {
	txOpts, err := getTxOpts(client, adminKeyPath)
	if err != nil {
		return nil, err
	}
	bridge, err := bridge.NewBridge(address, client)
	if err != nil {
		return nil, err
	}

	return &readerWriter{
		bridge: bridge,
		txOpts: txOpts,
	}, nil
}

func (r *readerWriter) Status() (*types.BridgeStatus, error) {
	id1, id2, err := r.bridge.CurrentIds(nil)
	if err != nil {
		return nil, err
	}
	return &types.BridgeStatus{
		DepositId:           id1,
		CrossChainDepositId: id2,
	}, nil
}

func (r *readerWriter) DepositInfo(depositId *big.Int) (*types.DepositInfo, error) {
	di, err := r.bridge.GetDepositInfo(nil, depositId)
	if err != nil {
		return nil, err
	}
	return &types.DepositInfo{
		User:   di.User,
		Amount: di.Amount,
	}, nil
}

func (r *readerWriter) CompleteTransfer(depositId *big.Int, depositInfo *types.DepositInfo) (*ethtypes.Transaction, error) {
	return r.bridge.CompleteTransfer(r.txOpts, depositInfo.User, depositId, depositInfo.Amount)
}

func getTxOpts(client *ethclient.Client, keyPath string) (*bind.TransactOpts, error) {
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	pkey, err := key.DecodeFromFile(keyPath)
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactorWithChainID(pkey, chainId)
}
