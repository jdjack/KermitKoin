package main

import "bytes"

type Chain struct {
  chain []Block
}

func (chain *Chain) isLonger(chain2 *Chain) bool {
  return len(chain.chain) > len(chain2.chain)
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

func (chain *Chain) validate() bool {
  for index, block := range chain.chain {
    // Check blocks are valid.
    if !block.Validate() {
      return false
    }
    // Check blocks are in order.
    if block.Index != index {
      return false
    }
    // Check all prev_hash values are correct.
    if index > 0 {
      if bytes.Compare(block.Prev_hash, (chain.chain[index - 1]).Hash) != 0 {
        return false
      }
    }
    // Check genesis block exists
    if len(chain.chain) == 0 {
      return false
    }
  }
  return true
}

func (chain *Chain) addBlock(block Block) {
  chain.chain = append(chain.chain, block)
}

func (chain *Chain) saveChain() {
  for _, block := range chain.chain {
    block.Save_block()
  }
}
