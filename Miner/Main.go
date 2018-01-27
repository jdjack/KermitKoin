package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	// Create a new router to route requests
	router := mux.NewRouter()

	// Route endpoints to function handlers
	router.HandleFunc("/getPeers", GetPeers).Methods("GET")
	router.HandleFunc("/getBlockchain", GetBlockchainReq).Methods("GET")
	router.HandleFunc("/authorizeBlock/{block}", AuthorizeBlockReq).Methods("POST")

	// Unknown request sent
	log.Fatal(http.ListenAndServe(":8081", router))

}
