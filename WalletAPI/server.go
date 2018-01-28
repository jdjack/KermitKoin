package main

import (
  "net/http"
  "log"
  "encoding/json"
  "strings"
	"io/ioutil"
  "fmt"
)

type balance struct {
  balance float64 `json:"balance"`
}

func StartHTTPServer() *http.Server {

  server := &http.Server{Addr: ":8082"}

  http.HandleFunc("/getBalance/{id}", GetBalanceReq)
  http.HandleFunc("/getWalletID/{key}", GetWalletAddrReq)
  http.HandleFunc("/makeTransaction", MakeTransactionReq)
  http.HandleFunc("/authorizeBlock", AuthorizeBlockReq)
  http.HandleFunc("/getPeers", GetPeersReq)
	http.HandleFunc("/getBlockchain", GetBlockchainReq)

  go func() {
    if err := server.ListenAndServe(); err != nil {
      log.Printf("Error: %s", err)
    }
  }()

  return server
}

// Shutdown the HTTP server
func ShutdownHTTPServer(server *http.Server) {

  // Try and shutdown the server
  if err := server.Shutdown(nil); err != nil {
    log.Fatal(err)
  }

}

func GetBalanceReq(w http.ResponseWriter, req *http.Request) {

  walletID, ok := req.URL.Query()["id"]

  if !ok {
    fmt.Printf("Bad Request")
    return
  }

  inputs := ValidInputs[walletID[0]]

  var sum float64 = 0

  for _, i := range(inputs) {
    sum += i.Amount
  }

  b := &balance{sum}
  json, err := json.Marshal(b)

  if err != nil {
    log.Printf("Bad json conversion: %s", err)
  }

  w.Write(json)

}

func MakeTransactionReq(w http.ResponseWriter, req *http.Request) {

}

func GetWalletAddrReq(w http.ResponseWriter, req *http.Request) {

}

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

func AuthorizeBlockReq(w http.ResponseWriter, r *http.Request) {

	// Fetch the request paramater

	body, _ := ioutil.ReadAll(r.Body)

	block := &Block{}

	err := json.Unmarshal(body, block)

	if err != nil {
	  panic(err)
  }

  r.Body.Close()

  ParseBlock(block)


	// Call json_to_block

	// Pass the result to authorize block
	// Return true or false if successful authorized

}
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
