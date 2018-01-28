package main

import (
  "os"
  "log"
  "bufio"
  "bytes"
  "encoding/json"
  "time"
  "math/rand"
)

func FetchChain() *Chain {
  success := false
	chain := &Chain{}
	for !success {

		// Choose a random peer to get the chain from
		randPeerIndex := rand.Intn(len(livePeers))
		randPeerIP := livePeers[randPeerIndex].IP
		r, err := netClient.Get("http://" + randPeerIP + ":8081/getBlockchain")

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
		buf := bytes.NewBuffer(make([]byte, 0))
		_, err = buf.ReadFrom(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		blocks := make([]Block,0)
		err = json.Unmarshal(buf.Bytes(), &blocks)
		if err != nil {
			log.Fatal(err)
		}
    chain.chain = blocks

		success = true

	}

	return chain

}

func FetchPeers() []Peer {
  if len(alwaysOnPeers) == 0 {
		return make([]Peer, 0)
	}

	rand.Seed(time.Now().UnixNano())

	// Choose a random always-on peer to connect to
	randAlwaysOnPeer := rand.Intn(len(alwaysOnPeers))
	randKnownIP := alwaysOnPeers[randAlwaysOnPeer].IP
	r, err := netClient.Get("http://" + randKnownIP + ":8081/getPeers")
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
	livePeers = append(livePeers, alwaysOnPeers[randAlwaysOnPeer])

	return livePeers
}

func LoadAlwaysOnPeers() []Peer {
  path := "/peers.txt"

	// If the file does not exist, use the backup IP
	if _, err := os.Stat(path); os.IsNotExist(err) {

		singletonResult := make([]Peer, 0)

		singletonResult = append(singletonResult, Peer{BackupIP})
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
