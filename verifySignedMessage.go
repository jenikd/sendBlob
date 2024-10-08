package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// handleVerifySig verifies the signature against the provided public key and hash.
func verifySig(signature, address, message string) bool {
	// Decode the signature into bytes
	sig, err := hexutil.Decode(signature)
	if err != nil {
		log.Fatalln(err)
	}

	// Adjust signature to standard format (remove Ethereum's recovery ID)
	sig[64] = sig[64] - 27

	// Construct the message prefix
	prefix := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message)))
	data := []byte(message)

	// Hash the prefix and data using Keccak-256
	hash := crypto.Keccak256Hash(prefix, data)

	// Recover the public key bytes from the signature
	sigPublicKeyBytes, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		log.Fatalln(err)
	}
	ecdsaPublicKey, err := crypto.UnmarshalPubkey(sigPublicKeyBytes)
	if err != nil {
		log.Fatalln(err)
	}

	// Derive the address from the recovered public key
	rAddress := crypto.PubkeyToAddress(*ecdsaPublicKey)

	// Check if the recovered address matches the provided address
	isSigner := strings.EqualFold(rAddress.String(), address)

	return isSigner
}
