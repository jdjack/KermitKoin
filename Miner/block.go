package main

import (
  "crypto/sha256"
  "encoding/json"
  "fmt"
  //"encoding/base64"
  //"time"
  "io/ioutil"
  "strconv"
  "strings"
  "os"
  "bytes"
)

var userName string
var oAuthToken string

type Block struct {
  Index int `json:"index"`
  Prev_hash []byte `json:"prev_hash"`
  Git_hash []byte `json:"git_hash"`
  Repo_id []byte `json:"repo_id"`
  Timestamp int `json:"timestamp"`
  Miner_transaction *transaction `json:"miner_transaction"`
  User_transaction *transaction `json:"user_transaction"`
  Hash []byte `jason:"hash"`
}

type Json_block struct {
  Index int `json:"index"`
  Prev_hash []byte `json:"prev_hash"`
  Git_hash []byte `json:"git_hash"`
  Repo_id []byte `json:"repo_id"`
  Timestamp []byte `json:"timestamp"`
  Data []byte `json:"data"`
  Hash []byte `json:"hash"`
}

type input struct {
  From []byte `json:"From"`
  Amount float64 `json:"Amount"`
  Hash []byte `json:"Hash"`
}

type output struct {
  To []byte `json:"To"`
  Amount float64 `json:"Amount"`
}

type transaction struct {
  Inputs []input `json:"Inputs"`
  Outputs []output `json:"Outputs"`
}

func Verify_transaction(t *transaction) bool {
  inputSum := 0.
  inputs := t.Inputs
  for _, i := range(inputs) {
    transactionBlock := CurrentChain.getBlockByHash(string(i.Hash))
    inputSum += i.Amount
    if !checkValue(i, transactionBlock.User_transaction.Outputs) {
      return false
    }
  }

  outputSum := 0.
  for _, o := range(t.Outputs) {
    outputSum += o.Amount
  }

  return outputSum == inputSum


}

func checkValue(i input, outputs []output) bool {
  for _, o := range(outputs) {
    if bytes.Compare(o.To, i.From) == 0 && o.Amount == i.Amount {
      return true
    }
  }
  return false
}

func (block *Block) BlockToJsonBlock() *Json_block {

}

func (block *Block) Generate_hash() []byte {
  h := sha256.New()
  combined := append(block.Prev_hash, []byte(string(block.Index))...)
  combined = append(combined, block.Prev_hash...)
  combined = append(combined, []byte(string(block.Timestamp))...)
  userTransactionJson, _ := json.Marshal(block.User_transaction)
  minerTransactionJson, _ := json.Marshal(block.Miner_transaction)
  combined = append(combined, userTransactionJson...)
  combined = append(combined, minerTransactionJson...)
  combined = append(combined, block.Git_hash...)

  h.Write(combined)
  return h.Sum(nil)
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

func (block *Block) Save_block() {
  filename := int_to_filename(block.Index)
  if _, err := os.Stat("chain-data") ; os.IsNotExist(err) {
    os.Mkdir("chain-data", 0700)
  }
  ioutil.WriteFile("chain-data/" + filename + ".json", block.Block_to_json(), 0644)
}

func int_to_filename(i int) string {
  // Create string with at least 16 digits
  str := strconv.Itoa(i)
  zeroes := strings.Repeat("0", 16 - len(str))
  return zeroes + str

}

func Validate(block *Block) bool {
  h := block.Generate_hash()
  if bytes.Compare(block.Hash, h) != 0 {
    return false
  }

  if !Verify_transaction(block.User_transaction) {
    return false
  }

  if !CheckCommitExistanceForUser(userName, string(block.Git_hash), oAuthToken) {
    return false
  }

  return true
}

func (block *Block) Add_transaction(data []byte) bool {
  t := &transaction{}
  json.Unmarshal(data, t)
  if !Verify_transaction(t) {
    return false
  }

  block.User_transaction = t
  return true
}

func AuthoriseBlock(block *Block) bool {
  return false
}

