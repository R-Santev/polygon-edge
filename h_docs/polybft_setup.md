## How to setup the polybft chain?

There is no official docs on polybft yet, so I will write my findings here.

# Local setup

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
