package main

import (
  "net/http"
  "log"
  "encoding/json"
  "strings"
	"io/ioutil"
  "fmt"
  "math/rand"
  "strconv"
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
  lowerWalletID := strings.ToLower(walletID[0])

  if !ok {
    fmt.Printf("Bad Request")
    return
  }

  inputs := ValidInputs[lowerWalletID]

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
  ownIDs, ok := req.URL.Query()["ownID"]
  destIDs, ok1 := req.URL.Query()["destID"]
  amounts, ok2 := req.URL.Query()["amount"]

  if !(ok && ok1 && ok2) {
    fmt.Printf("Bad Request")
    return
  }

  ownID := ownIDs[0]
  destID := destIDs[0]
  amountStr := amounts[0]
  amount, _ := strconv.ParseFloat(amountStr, 64)

  fmt.Println("============")
  fmt.Println(ownID)
  fmt.Println(destID)
  fmt.Println(amount)

  ownID = strings.ToLower(ownID)
  ourValidInputs := ValidInputs[ownID]

  fmt.Println(ownID)

  var counter float64 = 0
  index := 0
  inputsUsed := make([]input, 0)

  for counter < amount {
    inp := ourValidInputs[index]
    counter += inp.Amount
    inputsUsed = append(inputsUsed, inp)
    index++
  }

  change := counter - amount

  fmt.Println(change)

  if change < 0 {
    fmt.Printf("Insufficient Funds")
    return
  }

  outputs := make([]output, 0)
  recipientOutput := output{destID, amount}
  outputs = append(outputs, recipientOutput)
  if change != 0 {
    selfOutput := output{ownID, change}
    outputs = append(outputs, selfOutput)
  }

  transac := transaction{inputsUsed, outputs}

  fmt.Println(transac)

  // Encode the response and send it
  encoded, err := json.Marshal(transac)

  fmt.Println(encoded)

  fmt.Println("============")

  if err == nil {
    for _, peer := range(livePeers) {
      url := "http://" + peer.IP + ":8081/addTransaction"

      req, err := http.NewRequest("POST", url, strings.NewReader(string(encoded)))

      if err != nil {
        log.Fatal(err)
      }

      fmt.Print("Sending Block\n")

      netClient.Do(req)
    }


    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write([]byte("Success"))

  } else {
    log.Fatal(err)
  }

}

func GetWalletAddrReq(w http.ResponseWriter, req *http.Request) {


  keys, ok := req.URL.Query()["key"]

  if !ok {
    fmt.Printf("Bad Request")
    return
  }

  key := keys[0]

  db, _ := ioutil.ReadFile("keys.json")

  ksList := make([]Ids, 0)
  err := json.Unmarshal(db, &ksList)

  if err != nil {
    fmt.Printf("Error: %s", err)
  }


  for _, id := range(ksList) {
    //fmt.Println(id)
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

  fmt.Print(block)

  CurrentChain.addBlock(*block)


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
