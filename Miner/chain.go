package main

type Chain struct {
  chain []Block
}

func (chain *Chain) IsLonger(chain2 *Chain) bool {
  return len(chain.chain) > len(chain2.chain)
}

func (chain *Chain) GetBlockByIndex(index int) *Block {
  if index < 0 || index > len(chain.chain) {
    return nil
  } else {
    return &chain.chain[index]
  }
}

func (chain *Chain) GetBlockByHash(hash string) *Block {
  var retBlock Block
  for _, block := range chain.chain {
    if block.hash == hash {
      retBlock = block
      return &retBlock
    }
  }
  return nil
}

