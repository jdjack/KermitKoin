package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Peer struct {
	IP string `json:"IP"`
}

// Peers which are always alive
var alwaysOnPeers []Peer

// Peers which are currently alive
var livePeers []Peer

// Send all the peers that are alive
func GetPeers(w http.ResponseWriter, r *http.Request) {

	// Encode the response and send it
	encoded, err := json.Marshal(livePeers)
	if err != nil {
		w.Write(encoded)
	} else {
		log.Fatal(err)
	}

}

// Loads the always on peers from a file
func GetAlwaysOnPeers() []Peer {

	// Peers are stored as a list of IP's on each line of a file
	file, err := os.Open("/peers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read each peer into the peers array
	var peersFromDisk []Peer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		peersFromDisk = append(peersFromDisk, Peer{scanner.Text()})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return peersFromDisk

}

// Fetch the live peers from a known peer
func FetchLivePeers() []Peer {

	rand.Seed(time.Now().UnixNano())

	// Choose a random always-on peer to connect to
	randAlwaysOnPeer := rand.Intn(len(alwaysOnPeers) - 1)
	randKnownIP := alwaysOnPeers[randAlwaysOnPeer].IP
	r, err := netClient.Get("http://" + randKnownIP + "/getPeers")
	if err != nil {

		// Remove this peer from the list
		alwaysOnPeers = append(
			alwaysOnPeers[:randAlwaysOnPeer],
			alwaysOnPeers[randAlwaysOnPeer+1:]...)

		log.Fatal(err)
	}
	defer r.Body.Close()

	// Read and parse the response
	buf := bytes.NewBuffer(make([]byte, 0, r.ContentLength))
	_, readErr := buf.ReadFrom(r.Body)
	body := buf.Bytes()
	livePeers := make([]Peer, 0)
	json.Unmarshal(body, &livePeers)

	return livePeers
}
