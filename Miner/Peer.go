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
  "io/ioutil"
  "fmt"
)

var BackupIP string = "129.31.236.46"

var seenHashes [][]byte = make([][]byte, 0)

var seenTrans [][]byte = make([][]byte, 0)

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


	peersWithoutClient := make([]Peer, 0)

	isInLivePeers := false
	for _, livePeer := range livePeers {
		if livePeer.IP == client {
			isInLivePeers = true
		} else {
			peersWithoutClient = append(peersWithoutClient, livePeer)
		}
	}

	if !isInLivePeers {
		livePeers = append(livePeers, Peer{client})
	}

	// Encode the response and send it
	encoded, err := json.Marshal(peersWithoutClient)

	if err == nil {

		w.Write(encoded)

	} else {
		log.Fatal(err)
	}

}

// Loads the always on peers from a file
func LoadAlwaysOnPeers() []Peer {

	path := "/peers.txt"

	// If the file does not exist, use the backup IP
	if _, err := os.Stat(path); os.IsNotExist(err) {

		singletonResult := make([]Peer, 0)
		if getMyIP() == BackupIP {

			// I am the genesis user - load the blockchain if needed
			if CurrentChain == nil {
				CurrentChain = &Chain{}
				files, err := ioutil.ReadDir("chain-data/")
				if err != nil {
					log.Fatal(err)
				}

				for _, f := range files {
					if f.Name() == ".DS_Store" {
						continue
					}

					b := load_block_with_filename(f.Name())
					CurrentChain.addBlock(*b)
				}
			}

			// Do nothing, no peers
			return singletonResult

		}

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

// Fetch the live peers from a known peer
func FetchLivePeers() []Peer {

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

// Fetch the current blockchain from one of the live peers
func FetchCurrentBlockchain() *Chain {

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

// Handle a /getBlockchain request
func GetBlockchainReq(w http.ResponseWriter, r *http.Request) {

	// Give the user the current blockchain
	fmt.Println(CurrentChain.chain)
	encoded, err := json.Marshal(CurrentChain.chain)
	if err == nil {
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

	body, _ := ioutil.ReadAll(r.Body)

	block := &Block{}

	err := json.Unmarshal(body, block)

	if err != nil {
	  panic(err)
  }

  r.Body.Close()

  if checkIfHashSeen(block.Hash) {
    seenHashes = append(seenHashes, block.Hash)
    SendBlock(block)
    AuthoriseBlock(block)

  }


	// Call json_to_block

	// Pass the result to authorize block
	// Return true or false if successful authorized

}

func checkIfHashSeen(hash []byte) bool {
  for _, h := range(seenHashes) {
    if bytes.Compare(h, hash) == 0 {
      return false
    }
  }
  return true
}

func getMyIP() string {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func AddTransactionReq(w http.ResponseWriter, r *http.Request) {
  body, _ := ioutil.ReadAll(r.Body)

  t := &transaction{}

  err := json.Unmarshal(body, t)

  if err != nil {
    panic(err)
  }

  r.Body.Close()

  if checkIfTransSeen(body) {
    seenTrans = append(seenTrans, body)
    fmt.Println(seenTrans)
    SendTransaction(t)
    TransactionQueue.PushBack(t)
  }
}

func checkIfTransSeen(trans []byte) bool {
  for _, t := range(seenTrans) {
    if  bytes.Compare(trans, t) == 0 {
      return false
    }
  }
  return true
}

func equalTrans(t1, t2 *transaction) bool {


  for _, i := range(t1.Inputs) {
    for _, i1 := range(t2.Inputs) {
      if i.From != i1.From || i.Amount != i1.Amount || bytes.Compare(i.Hash, i1.Hash) != 0 {
        return false
      }
    }
  }

  for _, o := range(t1.Outputs) {
    for _, o1 := range(t2.Outputs) {
      if o.Amount != o1.Amount || o.To != o1.To {
        return false
      }
    }
  }


  return true
}