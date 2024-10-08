package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// getNonce fetches and prints the current and next nonce for a given Ethereum address.
func getNonce(address, rpcUrl string) (uint64, uint64) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// Retrieve the next nonce for the address
	nextNonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		log.Fatalln(err)
	}

	var currentNonce uint64 // Variable to hold the current nonce.
	if nextNonce > 0 {
		currentNonce = nextNonce - 1
	}

	return currentNonce, nextNonce
}
