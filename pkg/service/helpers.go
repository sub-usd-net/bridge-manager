package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sub-usd-net/bridge-manager/pkg/bindings/erc20"
	"github.com/sub-usd-net/bridge-manager/pkg/types"
	"math/big"
)

var (
	nativeTokenDecimals = big.NewInt(18)
	zeroAddress         = common.BigToAddress(big.NewInt(0))
)

func (b *BridgeService) geStableTokenDecimals(client *ethclient.Client) (*big.Int, error) {
	addr, err := b.c.(*readerWriter).bridge.StableToken(nil)
	if err != nil {
		return nil, err
	}
	token, err := erc20.NewErc20(addr, client)
	if err != nil {
		return nil, err
	}
	dec, err := token.Decimals(nil)
	if err != nil {
		return nil, err
	}

	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(dec)), nil), err
}

func (b *BridgeService) getCurrentStatus() (*types.CombinedBridgeStatus, error) {
	cStatus, err := b.c.Status()
	if err != nil {
		return nil, err
	}
	sStatus, err := b.s.Status()
	if err != nil {
		return nil, err
	}
	return types.CombineBridgeStatus(cStatus, sStatus), nil
}

func (b *BridgeService) nativeTokensForStableTokens(st *big.Int) *big.Int {
	return mulDiv(st, nativeTokenDecimals, b.stableTokenDecimals)
}

func (b *BridgeService) stableTokensForNativeTokens(nt *big.Int) *big.Int {
	return mulDiv(nt, b.stableTokenDecimals, nativeTokenDecimals)
}

func mulDiv(a, b, c *big.Int) *big.Int {
	return new(big.Int).Div(new(big.Int).Mul(a, b), c)
}
