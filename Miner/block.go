package main

import "crypto/sha256"

type Block struct {
  index int
  prev_hash string
  git_hash string
  repo_id string
  timestamp string
  data string
}

func (block *Block) verify_transaction() bool {

  return false
}

func (block *Block) generate_hash() string {
  h := sha256.New()
  h.Write([]byte(block.prev_hash + string(block.index) + block.timestamp + block.git_hash))
  return string(h.Sum(nil))
}

func (block *Block) validate() bool {
  return false;
}

