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
  "time"
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
  From   string     `json:"From"`
  Amount float64 `json:"Amount"`
  Hash   []byte  `json:"Hash"`
}

type output struct {
  To     string     `json:"To"`
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

func load_block(index int) *Block {
  filename := int_to_filename(index) + ".json"

  block := &Block{}
  if _, err := os.Stat("chain-data/" + filename); os.IsExist(err) {
    jsonBlock, _ := ioutil.ReadFile("chain-data/" + filename)
    _= json.Unmarshal(jsonBlock, block)
  }
  return block
}

func load_block_with_filename(filename string) *Block {
  block := &Block{}

  if _, err := os.Stat("chain-data/" + filename); os.IsNotExist(err) {
    return block
  }

  jsonBlock, _ := ioutil.ReadFile("chain-data/" + filename)
  _= json.Unmarshal(jsonBlock, block)

  return block

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
  if _, err := os.Stat("chain-data"); os.IsNotExist(err) {
    os.Mkdir("chain-data", 0700)
  }
  ioutil.WriteFile("chain-data/"+filename+".json", block.Block_to_json(), 0644)
}

func int_to_filename(i int) string {
  // Create string with at least 16 digits
  str := strconv.Itoa(i)
  zeroes := strings.Repeat("0", 16-len(str))
  return zeroes + str

}

func Validate(block *Block) bool {
  h := block.Generate_hash()
  if bytes.Compare(block.Hash, h) != 0 {
    return false
  }

  if block.User_transaction != nil {
    if !Verify_transaction(block.User_transaction) {
      return false
    }

  }

  if !CheckCommitExistanceForUser(block.UserName, string(block.Git_hash), oAuthToken) {
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
  if !Validate(block) {
    return false
  }


  latestBlock := CurrentChain.getLatestBlock()

  if (block.Index == latestBlock.Index) {
    RejectedChain := CurrentChain
    RejectedChain.RemoveLatestBlock()

    if bytes.Compare(block.Prev_hash, RejectedChain.getLatestBlock().Hash) == 0 {
      RejectedChain.addBlock(*block)
      return false
    }


  }

  if bytes.Compare(latestBlock.Hash, block.Prev_hash) == 0 {
    CurrentChain.addBlock(*block)
  } else if RejectedChain != nil {
    otherLatestBlock := RejectedChain.getLatestBlock()
    if bytes.Compare(otherLatestBlock.Hash, block.Prev_hash) == 0 {
      RejectedChain.addBlock(*otherLatestBlock)
    }
    if (RejectedChain.isLonger(CurrentChain)) {
      temp := RejectedChain
      RejectedChain = CurrentChain
      CurrentChain = temp
    }
  }

  block.Save_block()

  return true
}

func CreateBlock(git_hash []byte) bool {
  lastBlock := CurrentChain.getLatestBlock()

  miner_transaction := &transaction{
    Inputs: append(make([]input, 0), input{
      From:   "",
      Amount: 5.0,
      Hash:   nil,
    }),
    Outputs: append(make([]output, 0), output{
      To:     "12345",
      Amount: 5.0,
    }),
  }

  block := &Block{
    Index:             lastBlock.Index + 1,
    Prev_hash:         lastBlock.Hash,
    Git_hash:          git_hash,
    User_transaction:  nil,
    Timestamp:         time.Now().Unix(),
    UserName:          GetUserNameFromAuthToken(oAuthToken),
    Miner_transaction: miner_transaction,
  }

  block.Hash = block.Generate_hash()
  seenHashes = append(seenHashes, block.Hash)
  fmt.Println("Test")
  CurrentChain.addBlock(*block)
  block.Save_block()


  return true

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
