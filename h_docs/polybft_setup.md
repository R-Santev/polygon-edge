# Launch PolyBFT chain

There is no official docs on polybft yet, so I will write my findings here.
**EDIT:** there is already official docs but we segnificantly modify the way our consensus mechanism work because we don't have
root chain in opur setup. So I will continue describing the way to setup a chain in this file.

## Local setup

This setup is mainly used for development and testing purposes.

### Local Setup for version 0.9.x

The official docs can be found [here](https://wiki.polygon.technology/docs/category/launch-a-local-private-supernet).  
I am describing our custom process, because it is different.

1. Generate secrets

```
./polygon-edge polybft-secrets output --chain-id 187 --data-dir test-chain-1 --insecure /
./polygon-edge polybft-secrets output --chain-id 187 --data-dir test-chain-2 --insecure /
./polygon-edge polybft-secrets output --chain-id 187 --data-dir test-chain-3 --insecure /
./polygon-edge polybft-secrets output --chain-id 187 --data-dir test-chain-4 --insecure /
./polygon-edge polybft-secrets output --chain-id 187 --data-dir test-chain-5 --insecure

```

2. Generate genesis file

```
./polygon-edge genesis --block-gas-limit 10000000 --epoch-size 10 \ --validators-path ./ --validators-prefix test-chain- --consensus polybft --reward-wallet 0x61324166B0202DB1E7502924326262274Fa4358F:1000000 --chain-id 8844
```

4. Run the chain

```
./polygon-edge server --data-dir ./test-chain-1 --chain genesis.json --grpc-address :5001 --libp2p :30301 --jsonrpc :10001 --seal --log-level DEBUG --log-to ./log

./polygon-edge server --data-dir ./test-chain-2 --chain genesis.json --grpc-address :5002 --libp2p :30302 --jsonrpc :10002 --seal --log-level DEBUG --log-to ./log-2

./polygon-edge server --data-dir ./test-chain-3 --chain genesis.json --grpc-address :5003 --libp2p :30303 --jsonrpc :10003 --seal --log-level DEBUG --log-to ./log-3

./polygon-edge server --data-dir ./test-chain-4 --chain genesis.json --grpc-address :5004 --libp2p :30304 --jsonrpc :10004 --seal --log-level DEBUG --log-to ./log-4

./polygon-edge server --data-dir ./test-chain-5 --chain genesis.json --grpc-address :5005 --libp2p :30305 --jsonrpc :10005 --seal --log-level DEBUG --log-to ./log-5

```

### Local legacy setup

1. Generate secrets

```
polygon-edge polybft-secrets output --data-dir test-chain-1
polygon-edge polybft-secrets output --data-dir test-chain-2

```

2. Create manifest file
   This is the first version of edge that needs a manifest file. It contains information about the initial validators.

```
polygon-edge manifest --validators-prefix test-chain-
```

3. Generate genesis file

```
polygon-edge genesis --consensus polybft --ibft-validators-prefix-path test-chain-
```

4. Run the chain

```
polygon-edge server --data-dir ./test-chain-1 --chain genesis.json --grpc-address :10000 --libp2p :10001 --jsonrpc :10002 --seal --log-level=DEBUG
```
