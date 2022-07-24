package types

import (
	"io"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v2"
)

type BridgeConfig struct {
	CChainRpcUri string `yaml:"cChainRpcUri"`
	SubnetRpcUri string `yaml:"subnetRpcUri"`

	CChainBridgeAdminKeyPath string `yaml:"cChainBridgeAdminKeyPath"`
	SubnetBridgeAdminKeyPath string `yaml:"subnetBridgeAdminKeyPath"`

	CChainBridgeContractAddress common.Address `yaml:"cChainBridgeContractAddress"`
	SubnetBridgeContractAddress common.Address `yaml:"subnetBridgeContractAddress"`
}

func NewBridgeConfigFromPath(path string) (*BridgeConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var cfg BridgeConfig
	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *BridgeConfig) validate() error {
	// TODO
	return nil
}
