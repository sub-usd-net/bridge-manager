package key

import (
	"crypto/ecdsa"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func DecodeFromFile(path string) (*ecdsa.PrivateKey, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	s := strings.TrimSpace(string(content))
	return crypto.ToECDSA(common.FromHex(s))
}
