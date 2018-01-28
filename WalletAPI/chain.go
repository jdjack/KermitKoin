package main

import "fmt"

var CurrentChain *Chain

type Chain struct {
  chain []Block `json:"chain"`
}

func (chain *Chain) getBlockByIndex(index int) *Block {
  if index < 0 || index > len(chain.chain) {
    return nil
  } else {
    return &chain.chain[index]
  }
}

func (chain *Chain) getBlockByHash(hash string) *Block {
  var retBlock Block
  for _, block := range chain.chain {
    if string(block.Hash) == hash {
      retBlock = block
      return &retBlock
    }
  }
  return nil
}

func (chain *Chain) getLatestBlock() *Block {
  return &chain.chain[len(chain.chain) - 1]
}

func (chain *Chain) addBlock(block Block) {
  chain.chain = append(chain.chain, block)
  fmt.Println("Block Recieved")
  ValidInputs = make(map[string][]input, 0)
  ParseChain()
}

