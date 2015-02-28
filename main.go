package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hooks", hookHandler)

	port := os.Getenv("PORT")
	fmt.Printf("Running on port: %v", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Printf("Problem in ListenAndServe: %v", err)
	}
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error parsing body: %v", err)
	}

	var h hook
	if err := json.Unmarshal(body, &h); err != nil {
		log.Fatalf("Error unmarshalling hook: %v", err)
	}

	log.Printf("Hook is: %+v", h)

	w.WriteHeader(201)
}

type hook struct {
	Id              string `json:"id"`
	CreatedAt       string `json:"created_at"`
	DeliveryAttempt int    `json:"delivery_attempt"`
	Payload         payload
}

type payload struct {
	Type       string `json:"type"`
	BlockChain string `json:"block_chain"`
	Block      block  `json:"block"`
}

type block struct {
	Hash              string   `json:"hash"`
	PreviousBlockHash string   `json:"previous_block_hash"`
	Height            uint32   `json:"height"`
	Confirmations     int      `json:"confirmations"`
	MerkleRoot        string   `json:"merkle_root"`
	Time              string   `json:"time"`
	Nonce             uint32   `json:"nonce"`
	Difficulty        float64  `json:"difficulty"`
	Bits              string   `json:"bits"`
	TransactionHashes []string `json:"transaction_hashes"`
	ChainReceivedAt   string   `json:"chain_received_at"`
}
