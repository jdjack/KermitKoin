package main

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

