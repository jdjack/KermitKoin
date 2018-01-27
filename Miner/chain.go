package main

type Chain struct {
  chain []Block
}

func (chain1 *Chain) isLonger(chain2 *Chain) bool {
  return len(chain1.chain) > len(chain2.chain)
}
