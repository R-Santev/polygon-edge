#!/bin/bash

# Genesis script
main() {
  GENESIS_URL="https://raw.githubusercontent.com/R-Santev/hydrag-files/genesis/genesis.json"

  # Fetch and save genesis.json in the current directory
  curl -o genesis.json $GENESIS_URL
}

main
