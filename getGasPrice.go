package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// getSuggestedGasPrice connects to an Ethereum node via RPC and retrieves the current suggested gas price.

func getSuggestedGasPrice(rpcUrl string) {

	// Connect to the Ethereum network using the provided RPC URL.
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Retrieve the currently suggested gas price for a new transaction.
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	// Print the suggested gas price to the terminal.
	fmt.Println("Suggested Gas Price:", gasPrice.String())
}
