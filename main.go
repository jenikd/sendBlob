package main

import (
	"fmt"
	"log"
)

const (
	sepoliaRpcUrl = "http://localhost:8545" // sepolia rpc url
	from          = "0x239fA7623354eC26520dE878B52f13Fe84b06971"
	to            = "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5"
	data          = "Hello Ethereum!"
	privKey       = "163f5f0f9a621d72fedd85ffca3d08d131ab4e812181e0d30ffd1c885d20aac7"
	gasLimit      = uint64(21500) // adjust this if necessary
	wei           = uint64(0)     // 0 Wei
)

func main() {

	fmt.Println("using ethclient...")

	getSuggestedGasPrice(sepoliaRpcUrl)
	// get gas price on sepolia testnet. This was just added.

	eGas := estimateGas(sepoliaRpcUrl, from, to, data, wei)     // This was just added.
	fmt.Println("\nestimate gas for the transaction is:", eGas) // This was just added.

	rawTxRLPHex := createRawTransaction(sepoliaRpcUrl, to, data, privKey, gasLimit, wei) // This was just added.
	fmt.Println("\nRaw TX:\n", rawTxRLPHex)                                              // This was just added.

	sendRawTransaction(rawTxRLPHex, sepoliaRpcUrl) // This was just added.

	// sig, sDetails := signMessage(data, privKey) // This was just added.
	// fmt.Println("\nsigned message:", sDetails)  // This was just added.

	// if isSigner := verifySig(sig, from, data); isSigner { // This was just added.
	// 	fmt.Printf("\n%s signed %s\n", from, data)
	// } else {
	// 	fmt.Printf("\n%s did not sign %s\n", from, data)
	// }

	cNonce, nNonce := getNonce(to, sepoliaRpcUrl)      // This was just added.
	fmt.Printf("\n%s current nonce: %v\n", to, cNonce) // This was just added.
	fmt.Printf("%s next nonce: %v\n", to, nNonce)      // This was just added.

	// res := traceTx("0x6f02e6cea53cbc6ecda93dab6faeb94887b3f9c76dc44e318b7c91587f60a926", sepoliaRpcUrl) // This was just added.
	// fmt.Println("\ntrace result:\n", res)                                                               // This was just added.

	blob, err := sendBlobTX(sepoliaRpcUrl, to, data, privKey) // This was just added.
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\nBlob transaction hash:", blob) // This was just added.
}
