package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
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
func GetPeersReq(w http.ResponseWriter, r *http.Request) {

	// Add the requester to the list of alive IP's
	client := strings.Split(r.RemoteAddr, ":")[0]
	livePeers = append(livePeers, Peer{client})

	// Encode the response and send it
	encoded, err := json.Marshal(livePeers)
	if err != nil {
		w.Write(encoded)
	} else {
		log.Fatal(err)
	}

}

// Loads the always on peers from a file
func LoadAlwaysOnPeers() []Peer {

	path := "/peers.txt"
	backupIP := "129.31.196.107"

	// If the file does not exist, use the backup IP
	if _, err := os.Stat(path); os.IsNotExist(err) {

		singletonResult := make([]Peer, 0)
		if getMyIP() == backupIP {
			// Do nothing, no peers
			return singletonResult
		}

		singletonResult = append(singletonResult, Peer{backupIP})
		return singletonResult
	}

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

	if len(alwaysOnPeers) == 0 {
		return make([]Peer, 0)
	}

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
	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := buf.Bytes()
	livePeers := make([]Peer, 0)
	json.Unmarshal(body, &livePeers)

	return livePeers
}

// Fetch the current blockchain from one of the live peers
func FetchCurrentBlockchain() Chain {

	success := false
	chain := &Chain{}
	for !success {

		// Choose a random peer to get the chain from
		randPeerIndex := rand.Intn(len(livePeers) - 1)
		randPeerIP := livePeers[randPeerIndex].IP
		r, err := netClient.Get("http://" + randPeerIP + "/getBlockchain")

		if err != nil {
			// Remove this peer from the list
			livePeers = append(
				livePeers[:randPeerIndex],
				livePeers[randPeerIndex+1:]...)

			log.Print(err)
			continue
		}
		defer r.Body.Close()

		// Decode the response
		buf := bytes.NewBuffer(make([]byte, 0, r.ContentLength))
		_, err = buf.ReadFrom(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		// body := buf.Bytes()
		// TODO: pass to liam
		success = true

	}

	return *chain

}

// Handle a /getBlockchain request
func GetBlockchainReq(w http.ResponseWriter, r *http.Request) {

	// Give the user the current blockchain
	encoded, err := json.Marshal(CurrentChain)
	if err != nil {
		_, err := w.Write(encoded)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}

}

// Handles a /authorizeBlock request
func AuthorizeBlockReq(w http.ResponseWriter, r *http.Request) {

	// Fetch the request paramater

	decoder := json.NewDecoder(r.Body)

	block := &Block{}

	err := decoder.Decode(block)

	if err != nil {
	  panic(err)
  }

  r.Body.Close()

  if AuthoriseBlock(block) {
    SendBlock(block)
  }

	// Call json_to_block

	// Pass the result to authorize block
	// Return true or false if successful authorized

}

func getMyIP() string {

	ifaces, _ := net.Interfaces()
	var result string
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// Determine if we are a 4-byte or 16-byte IP address
			stringIP := ip.String()
			if stringIP != "" {
				result = stringIP
			}

		}
	}
	return result

}
