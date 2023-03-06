Polygon Edge uses two submodules one of which is the core contracts repo.

We are working on a fork of the core contracts, so the submodule is update the work with our custom version.

Commands used for the module update:

```
git submodule set-url core-contracts https://github.com/R-Santev/core-contracts.git
```

```
git submodule update --init  --remote ./core-contracts
```

## How to generate core-contracts binaries:

After successful init and update of the core-contracts submodule run the following:

```
go run ./consensus/polybft/contractsapi/artifacts-gen/main.go
go run ./consensus/polybft/contractsapi/bindings-gen/main.go
```
