# Devnet setup

The purpose of the devnet setup is to test the performance, availability and stability of the networks as well as to ensure it is bug - free.

## Installation

Ensure you have Docker installed and pull the image.

```
docker pull rsantev/hydrag-devnet:latest
```

## Run node

In case you already have the needed secrets, provide them the following way:

```
docker run -it -e KEY=<your_key> -e BLS_KEY=<your_bls_key> -e SIG=<your_sig> -e P2P_KEY=<your_p2p_key> rsantev/hydrag-devnet
```

Otherwise:

docker run -it rsantev/hydrag-devnet
