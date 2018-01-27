package main

import (
  "crypto/sha256"
  "encoding/json"
  "fmt"
  "encoding/base64"
  //"net/http"
  //"time"
  "io/ioutil"
  "strconv"
  "strings"
  "os"
  "net/http"
  "time"
)

type Block struct {
  Index int `json:"index"`
  Prev_hash []byte `json:"prev_hash"`
  Git_hash []byte `json:"git_hash"`
  Repo_id []byte `json:"repo_id"`
  Timestamp []byte `json:"timestamp"`
  Data []byte `json:"data"`
  Hash []byte `json:"hash"`
}

func (block *Block) Verify_transaction(chain Chain) bool {
  return false
}

func (block *Block) Generate_hash() []byte {
  h := sha256.New()
  combined := append(block.Prev_hash, []byte(string(block.Index))...)
  combined = append(combined, block.Prev_hash...)
  combined = append(combined, block.Timestamp...)
  combined = append(combined, block.Data...)
  combined = append(combined, block.Git_hash...)

  h.Write(combined)
  return h.Sum(nil)
}

func (block *Block) Block_to_json() []byte {
  data := base64.StdEncoding.EncodeToString(block.Data)
  b := block
  b.Data = []byte(data)
  j, err := json.Marshal(b)
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
  var decoded_data []byte
  base64.StdEncoding.Decode(block.Data, decoded_data)
  block.Data = decoded_data
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

func (block *Block) Validate() bool {
  //hash := block.Generate_hash()
  //
  //git_url := "https://api.github.com/graphql"
  //
  //client := http.Client{Timeout: time.Second * 10}
  //
  //req, err := http.NewRequest(http.MethodGet, git_url, nil)
  //
  //res, getErr := client.Do(req)
  //
  //// token : 915654d075db14e717a429e34f4fb3ce37cdc333
  return false
}


