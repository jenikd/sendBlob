package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

func traceTx(hash, rpcUrl string) string {
	var (
		client *rpc.Client // Define a variable to hold the RPC client.
		err    error       // Variable to catch errors.
	)

	// Connect to the Ethereum RPC endpoint using the provided URL.
	client, err = rpc.Dial(rpcUrl)
	if err != nil {
		log.Fatalln(err)
	}

	var result json.RawMessage // Variable to hold the raw JSON result of the call.

	// Make the RPC call to trace the transaction using its hash. `ots_traceTransaction` is the method name.
	err = client.CallContext(context.Background(), &result, "debug_traceTransaction", hash, map[string]any{"tracer": "callTracer", "tracerConfig": map[string]any{"withLog": true}}) // or use debug_traceTransaction with a supported RPC URL and params: hash, map[string]any{"tracer": "callTracer", "tracerConfig": map[string]any{"withLog": true}} for Geth tracing
	if err != nil {
		log.Fatalln(err)
	}

	// Marshal the result into a formatted JSON string
	resBytes, err := json.MarshalIndent(result, " ", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	return string(resBytes)
}
