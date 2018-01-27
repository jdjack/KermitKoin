package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type IP string

type Peer struct {
	IPAddr IP   `json:"IP"`
	Alive  bool `json:"-"`
}

var knownPeers []Peer

// Send all the peers that are alive
func GetPeers(w http.ResponseWriter, r *http.Request) {

	// Filter peers that are alive
	var livePeers []Peer
	for _, peer := range peers {
		if peer.Alive {
			append(livePeers, peer)
		}
	}

	// Encode the response and send it
	w.write(json.Marshal(peers))

}

// Loads the current known peers from a file
func GetKnownPeers() []IP {

	// Peers are stored as a list of IP's on each line of a file
	file, err := os.Open("/peers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read each peer into the peers array
	var knownPeers []IP
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		append(knownPeers, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return knownPeers

}
