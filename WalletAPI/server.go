package main

import (
  "net/http"
  "log"
  "encoding/json"
  "strings"
	"io/ioutil"
  "fmt"
  "math/rand"
)

type balance struct {
  Balance float64 `json:"balance"`
}

type Ids struct {
  Key string `json:"key"`
  Address string `json:"address"`
}

type address struct {
  Address string `json:"Address"`
}

func StartHTTPServer() *http.Server {

  server := &http.Server{Addr: ":8080"}

  server.SetKeepAlivesEnabled(false)
  http.HandleFunc("/getBalance", GetBalanceReq)
  http.HandleFunc("/getAddress", GetWalletAddrReq)
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
  fmt.Println(b)
  j, err := json.Marshal(b)

  if err != nil {
    log.Printf("Bad json conversion: %s", err)
  }

  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Write(j)

}

func MakeTransactionReq(w http.ResponseWriter, req *http.Request) {

}

func GetWalletAddrReq(w http.ResponseWriter, req *http.Request) {


  keys, ok := req.URL.Query()["key"]

  if !ok {
    fmt.Printf("Bad Request")
    return
  }

  key := keys[0]

  db, _ := ioutil.ReadFile("keys.json")
  fmt.Printf(string(db))

  ksList := make([]Ids, 0)
  err := json.Unmarshal(db, ksList)

  if err != nil {
    fmt.Printf("Error: %s", err)
  }

  fmt.Println(ksList)
  for _, id := range(ksList) {
    if id.Key == key {
      output, _ := json.Marshal(&address{id.Address})
      w.Write(output)
      return
    }
  }

  addr := &Address{rand.Int()}
  id := &Ids{key, addr.toHex()}
  ksList = append(ksList, *id)

  j1, _ := json.Marshal(ksList)

  ioutil.WriteFile("keys.json", j1, 0644)

  w.Header().Set("Access-Control-Allow-Origin", "*")
  output, _ := json.Marshal(&address{addr.toHex()})
  w.Write(output)
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
