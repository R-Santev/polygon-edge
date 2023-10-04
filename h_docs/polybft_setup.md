# Launch PolyBFT chain

There is no official docs on polybft yet, so I will write my findings here.
**EDIT:** there is already official docs but we segnificantly modify the way our consensus mechanism work because we don't have
root chain in opur setup. So I will continue describing the way to setup a chain in this file.

## Local setup

This setup is mainly used for development and testing purposes.

### Local Setup for version 0.9.x

#### Initial chain setup

The official docs can be found [here](https://wiki.polygon.technology/docs/category/launch-a-local-private-supernet).  
I am describing our custom process, because it is different.

1. Generate secrets

```
./polygon-edge polybft-secrets output --chain-id 8844 --data-dir test-chain-1 --insecure /
./polygon-edge polybft-secrets output --chain-id 8844 --data-dir test-chain-2 --insecure /
./polygon-edge polybft-secrets output --chain-id 8844 --data-dir test-chain-3 --insecure /
./polygon-edge polybft-secrets output --chain-id 8844 --data-dir test-chain-4 --insecure /
./polygon-edge polybft-secrets output --chain-id 8844 --data-dir test-chain-5 --insecure

```

2. Generate genesis file

We need to set native token to be mintable, so we can premine balances to different addresses

```
./polygon-edge genesis --block-gas-limit 10000000 --epoch-size 10 \ --validators-path ./ --validators-prefix test-chain- --consensus polybft --native-token-config Hydra:HDR:18:true:0x211881Bb4893dd733825A2D97e48bFc38cc70a0c   --reward-wallet 0x61324166B0202DB1E7502924326262274Fa4358F:1000000 --premine 0x211881Bb4893dd733825A2D97e48bFc38cc70a0c:70000000000000000000000 --premine 0xdC3312E368A178e24850C6dAC169646c5fD14b93:700000000000000000000 --chain-id 8844
```

4. Run the chain

```
./polygon-edge server --data-dir ./test-chain-1 --chain genesis.json --grpc-address :5001 --libp2p :30301 --jsonrpc :10001 --seal --log-level DEBUG --log-to ./log

./polygon-edge server --data-dir ./test-chain-2 --chain genesis.json --grpc-address :5002 --libp2p :30302 --jsonrpc :10002 --seal --log-level DEBUG --log-to ./log-2

./polygon-edge server --data-dir ./test-chain-3 --chain genesis.json --grpc-address :5003 --libp2p :30303 --jsonrpc :10003 --seal --log-level DEBUG --log-to ./log-3

./polygon-edge server --data-dir ./test-chain-4 --chain genesis.json --grpc-address :5004 --libp2p :30304 --jsonrpc :10004 --seal --log-level DEBUG --log-to ./log-4

./polygon-edge server --data-dir ./test-chain-5 --chain genesis.json --grpc-address :5005 --libp2p :30305 --jsonrpc :10005 --seal --log-level DEBUG --log-to ./log-5

```

#### Add more validators

1. Generate new account (secrets):

```
./polygon-edge polybft-secrets output --chain-id 8844 --data-dir test-add-chain-1 --insecure

```

2. Use the governer (first validator by default) to whitelist the new account

```
./polygon-edge polybft whitelist-validator --data-dir ./test-chain-1 --address 0x7A94400e0d33B79B6C69979df3f7a46CF1963c69 --jsonrpc http://127.0.0.1:10001

```

3. Register account

--data-dir is the direcotry of the freshly created secrets  
Stake tx is made in this step as well

```

./polygon-edge polybft register-validator --data-dir ./test-add-chain-1 --stake 1000000000000000000 --chain-id 8844 --jsonrpc http://127.0.0.1:10001

```

4. Run new validator

```

./polygon-edge server --data-dir ./test-add-chain-1 --chain genesis.json --grpc-address :5006 --libp2p :30306 --jsonrpc :10006 --seal --log-level DEBUG --log-to ./log-6

```

### LEGACY local setup

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
