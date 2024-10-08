package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction represents the structure of the transaction JSON.
type Transaction struct {
	Type                 string   `json:"type"`
	ChainID              string   `json:"chainId"`
	Nonce                string   `json:"nonce"`
	To                   string   `json:"to"`
	Gas                  string   `json:"gas"`
	GasPrice             string   `json:"gasPrice,omitempty"`
	MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         string   `json:"maxFeePerGas"`
	Value                string   `json:"value"`
	Input                string   `json:"input"`
	AccessList           []string `json:"accessList"`
	V                    string   `json:"v"`
	R                    string   `json:"r"`
	S                    string   `json:"s"`
	YParity              string   `json:"yParity"`
	Hash                 string   `json:"hash"`
	TransactionTime      string   `json:"transactionTime,omitempty"`
	TransactionCost      string   `json:"transactionCost,omitempty"`
}

// sendRawTransaction sends a raw Ethereum transaction.
func sendRawTransaction(rawTx, rpcURL string) {
	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize an empty Transaction struct to hold the decoded data.
	tx := new(types.Transaction)

	// Decode the raw transaction bytes from hexadecimal to a Transaction struct.
	// This step converts the RLP (Recursive Length Prefix) encoded bytes back into
	// a structured Transaction format understood by the Ethereum client.
	err = rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		log.Fatalln(err)
	}

	// Establish an RPC connection to the specified RPC url
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalln(err)
	}

	// Propagate the transaction
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalln(err)
	}

	// Unmarshal the transaction JSON into a struct
	var txDetails Transaction
	txBytes, err := tx.MarshalJSON()
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(txBytes, &txDetails); err != nil {
		log.Fatalln(err)
	}

	// Add additional transaction details
	txDetails.TransactionTime = tx.Time().Format(time.RFC822)
	txDetails.TransactionCost = tx.Cost().String()

	// Format some hexadecimal string fields to decimal string
	convertFields := []string{"Nonce", "MaxPriorityFeePerGas", "MaxFeePerGas", "Value", "Type", "Gas"}
	for _, field := range convertFields {
		if err := convertHexField(&txDetails, field); err != nil {
			log.Fatalln(err)
		}
	}

	// Marshal the struct back to JSON
	txJSON, err := json.MarshalIndent(txDetails, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	// Print the entire JSON with the added fields
	fmt.Println("\nRaw TX Receipt:\n", string(txJSON))
}

func convertHexField(tx *Transaction, field string) error {

	// Get the type of the Transaction struct
	typeOfTx := reflect.TypeOf(*tx)

	// Get the value of the Transaction struct
	txValue := reflect.ValueOf(tx).Elem()

	// Parse the hexadecimal string as an integer
	hexStr := txValue.FieldByName(field).String()

	intValue, err := strconv.ParseUint(hexStr[2:], 16, 64)
	if err != nil {
		return err
	}

	// Convert the integer to a decimal string
	decimalStr := strconv.FormatUint(intValue, 10)

	// Check if the field exists
	_, ok := typeOfTx.FieldByName(field)
	if !ok {
		return fmt.Errorf("field %s does not exist in Transaction struct", field)
	}

	// Set the field value to the decimal string
	txValue.FieldByName(field).SetString(decimalStr)

	return nil
}
