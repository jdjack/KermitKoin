package main

import (
  "encoding/json"
  "fmt"
  "strings"
  "net/http"
  "log"
)

type Block struct {
  Index             int          `json:"index"`
  Prev_hash         []byte       `json:"prev_hash"`
  Git_hash          []byte       `json:"git_hash"`
  UserName          string       `json:"userName"`
  Timestamp         int64        `json:"timestamp"`
  Miner_transaction *transaction `json:"miner_transaction"`
  User_transaction  *transaction `json:"user_transaction"`
  Hash              []byte       `json:"hash"`
}

type Json_block struct {
  Index     int    `json:"index"`
  Prev_hash []byte `json:"prev_hash"`
  Git_hash  []byte `json:"git_hash"`
  Repo_id   []byte `json:"repo_id"`
  Timestamp []byte `json:"timestamp"`
  Data      []byte `json:"data"`
  Hash      []byte `json:"hash"`
}

type input struct {
  From   int     `json:"From"`
  Amount float64 `json:"Amount"`
  Hash   []byte  `json:"Hash"`
}

type output struct {
  To     int     `json:"To"`
  Amount float64 `json:"Amount"`
}

type transaction struct {
  Inputs  []input  `json:"Inputs"`
  Outputs []output `json:"Outputs"`
}

func Verify_transaction(t *transaction) bool {
  inputSum := 0.
  inputs := t.Inputs
  for _, i := range (inputs) {
    transactionBlock := CurrentChain.getBlockByHash(string(i.Hash))
    inputSum += i.Amount
    if !checkValue(i, transactionBlock.User_transaction.Outputs) {
      return false
    }
  }

  outputSum := 0.
  for _, o := range (t.Outputs) {
    outputSum += o.Amount
  }

  return outputSum == inputSum

}

func checkValue(i input, outputs []output) bool {
  for _, o := range (outputs) {
    if o.To == i.From && o.Amount == i.Amount {
      return true
    }
  }
  return false
}

func (block *Block) Block_to_json() []byte {
  //data_json, _ := json.Marshal(block.User_transaction)
  //minerJson, _ := json.Marshal(block.Miner_transaction)
  //data_json = append(data_json, minerJson...)
  //data := base64.StdEncoding.EncodeToString(data_json)
  //b := Json_block{block.Index, block.Prev_hash, block.Git_hash, block.Repo_id, []byte(string()), []byte(data), block.Hash}
  //b.Data = []byte(data)
  j, err := json.Marshal(block)
  if err != nil {
    fmt.Printf("Error %s", err)
    return make([]byte, 0)
  }
  return j
}

func Json_to_block(json_string []byte) *Block {
  block := &Block{}
  err := json.Unmarshal(json_string, block)
  if err != nil {
    fmt.Printf("Error: %s", err)
    return nil
  }
  //var decoded_data []byte
  //base64.StdEncoding.Decode(json_block.Data, decoded_data)
  //t, err := strconv.Atoi(string(json_block.Timestamp))
  //block := &Block{
  //  Index :     json_block.Index,
  //  Prev_hash : json_block.Prev_hash,
  //  Git_hash :  json_block.Git_hash,
  //  Repo_id :   json_block.Repo_id,
  //  Timestamp : t,
  //  Hash :      json_block.Hash,
  //  }

  return block
}


func SendBlock(block *Block) {
  jsonBlock := block.Block_to_json()

  for _, peer := range(livePeers) {
    url := "http://" + peer.IP + ":8081/authorizeBlock"

    req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonBlock)))

    if err != nil {
      log.Fatal(err)
    }

    fmt.Print("Sending Block\n")

    netClient.Do(req)
  }
}
