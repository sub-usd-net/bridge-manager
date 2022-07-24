set -e

mkdir -p pkg/bindings/bridge
abigen --abi abi/bridge.abi --pkg bridge --out pkg/bindings/bridge/bridge.go

mkdir -p pkg/bindings/erc20
abigen --abi abi/erc20.abi --pkg erc20 --out pkg/bindings/erc20/token.go
