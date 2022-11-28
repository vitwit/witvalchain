# witvalchain

**witvalchain** is built using the Cosmos SDK framework. `witvalchain` allows the creation and management of on-chain multisig accounts and enables voting for message execution based on configurable decision policies.

## Installing Go

Currently, witvalchain uses Go 1.18 to compile the code.

- [Go](https://go.dev/dl/) `>=1.18`
- Docker

Verify the installation by executing go version in your terminal:


```
$ go version
go version go1.18.1 darwin/arm64
```

## Build witvalchain

Build witvalchain from the source code:


```
make build
```

After building, you should see executable file  `build/witvalchain`.

## Generating protobuf 

```
make proto-gen
```
