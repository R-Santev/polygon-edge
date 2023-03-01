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

3. Compile contracts
   When cloning polygon-edge the core contracts are added as a submodule. Initialize and update submodules:

```
git submodule init
git submodule update
```

Copy the core-contracts folder and paste it in your setup directory (on the level of test-chain-1 and test-chain-2)

Then install the packages and compile the contracts

```
npm i
npx hardhat compile
```

4. Generate genesis file

```
polygon-edge genesis --consensus polybft --ibft-validators-prefix-path test-chain-
```

5. Run the chain

```
polygon-edge server --data-dir ./test-chain-1 --chain genesis.json --grpc-address :10000 --libp2p :10001 --jsonrpc :10002 --seal --log-level=DEBUG
```
