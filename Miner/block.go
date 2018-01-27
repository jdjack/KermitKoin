package main


type Block struct {
  index int
  prev_hash int
  git_hash string
  repo_id string
  timestamp string
  data string
}

func (block *Block) verify_transaction() bool {

  return false
}

func (block *Block) generate_hash() string {
  return nil
}

func (block *Block) validate() bool {
  return false;
}

