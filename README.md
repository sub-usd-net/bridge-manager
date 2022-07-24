
## Bridge Manager

Bridge Manager Service for our super-simple Fuji -> Subnet [bridge](https://github.com/sub-usd-net/bridge)

The service monitors the bridge contracts and completes the transfer on the other side. (Still a WIP)

### Instructions

Copy `config.sample.yaml` to `config.yaml` and edit with bridge addresses / keys and then:

```shell
$ go run main/main.go service
```
