package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var oAuthToken string

type Peer struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var peers []Peer

func main() {

	// Put some dummy people in the struct
	peers = append(peers, Peer{ID: "1", Name: "Tom"})

	// Create a new router to route requests
	router := mux.NewRouter()

	// Route endpoints to function handlers
	router.HandleFunc("/getPeers", GetPeers).Methods("GET")
	router.HandleFunc("/getBlockchain/", GetBlockchain).Methods("GET")
	router.HandleFunc("/authorizeBlock/{block}", AuthorizeBlock).Methods("POST")

	// Unknown request sent
	log.Fatal(http.ListenAndServe(":8081", router))

}

// Each handler function has a 'w' (write) and 'r' (read) parameter
func GetPeers(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(peers)

}
func GetBlockchain(w http.ResponseWriter, r *http.Request) {

}
func AuthorizeBlock(w http.ResponseWriter, r *http.Request) {

}
