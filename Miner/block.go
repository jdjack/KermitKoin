package main

import (
  "crypto/sha256"
  "encoding/json"
  "fmt"
  "encoding/base64"
)

type Block struct {
  index int
  prev_hash string
  git_hash string
  repo_id string
  timestamp string
  data []byte
  hash []byte
}

func (block *Block) Verify_transaction() bool {
  return false
}

func (block *Block) Generate_hash() string {
  h := sha256.New()
  h.Write([]byte(block.prev_hash + string(block.index) + block.timestamp + block.git_hash))
  block.hash = h.Sum(nil)
  return string(h.Sum(nil))
}

func (block *Block) Block_to_json() string {
  j, err := json.Marshal(block)
  if err != nil {
    fmt.Printf("Error %s", err)
    return ""
  }
  return string(j)
}

func Json_to_block(json_string []byte) *Block {
  block := &Block{}
  err := json.Unmarshal(json_string, block)
  if err != nil {
    fmt.Printf("Error: %s", err)
    return nil
  }
  return block
}

func (block *Block) Validate() bool {
  return false
}


