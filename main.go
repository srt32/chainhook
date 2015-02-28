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

type hook struct {
	Id              string `json:"id"`
	CreatedAt       string `json:"created_at"`
	DeliveryAttempt int    `json:"delivery_attempt"`
	Payload         payload
}

type payload struct {
	Type            string
	BlockChain      string
	Block           block
	ChainReceivedAt string
}

type block struct {
	//          "hash": "0000000000000000329fcc0e6cd53b7ae32b6c00e3f8bf356ea898f5b23a73e7",
	//          "previous_block_hash": "0000000000000000185e56fbad963648dc1916e20e9ccd07e4308b3013812659",
	//          "height": 314645,
	//          "confirmations": 1,
	//          "merkle_root": "5994b79a9bd567a9c4199bea91529fd0d61b75c0b7fff2e7e7d60112c4987001",
	//          "time": "2014-08-09T03:33:10Z",
	//          "nonce": 2324252789,
	//          "difficulty": 19729645940.577133,
	//          "bits": "1837ba62",
	//          "transaction_hashes": [
	//            "62278007bbcced2a2e9d3cc118f447ebd13d9e8119d50bfb56ad98050aafadfd",
	//            "602a2eb5918b3d75819aec57c26f8bf211ab743e97c4afdd30b6918339d7cc84"
	//          ],
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

	log.Printf("Hook is: %v", h)

	w.WriteHeader(201)
}
