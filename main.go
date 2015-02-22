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

	var h interface{}
	if err := json.Unmarshal(body, &h); err != nil {
		log.Printf("Error unmarshallig hook: %v", err)
	}

	log.Printf("Hook is: %v", h)

	w.WriteHeader(201)
}
