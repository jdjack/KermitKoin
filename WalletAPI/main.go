package main

import (
  "bufio"
  "os"
  "net/http"
  "time"
  "fmt"
)

type Peer struct {
  IP string `json:"IP"`
}

var netClient = http.Client{
  Timeout: time.Second * 10,
}

var ValidInputs map[string][]input = make(map[string][]input, 0)

var BackupIP string = "129.31.236.46"

var livePeers []Peer = make([]Peer, 0)
var alwaysOnPeers []Peer = make([]Peer, 0)

func main() {

  alwaysOnPeers = LoadAlwaysOnPeers()
  livePeers = FetchPeers()
  fmt.Println(livePeers)
  CurrentChain = FetchChain()

  ParseChain()

  server := StartHTTPServer()
  defer ShutdownHTTPServer(server)

  var text string
  for text != "shutdown\n" {
    reader := bufio.NewReader(os.Stdin)
    text, _ = reader.ReadString('\n')
  }
}

func ParseChain() {
  for _, block := range (CurrentChain.chain) {

    ParseBlock(&block)

  }

}

func CheckInputs(inputs []input, in input) {

  for index, i := range (inputs) {
    if i.Amount == in.Amount {
      x := inputs[:index]
      y := inputs[index+1:]
      inputs = append(x, y...)
      return
    }
  }

}

func AddInput(inputs []input, out output, block Block) []input {
  return append(inputs, input{out.To, out.Amount, block.Hash})
}

func ParseBlock(block *Block) {
  if block.User_transaction != nil {
    for _, in := range (block.User_transaction.Inputs) {
      if inList, ok := ValidInputs[string(in.From)]; ok {

        CheckInputs(inList, in)

      }
    }

    for _, out := range (block.User_transaction.Outputs) {
      if inList, ok := ValidInputs[string(out.To)]; ok {

        ValidInputs[string(out.To)] = AddInput(inList, out, *block)

      } else {

        inputs := make([]input, 0)
        inputs = AddInput(inputs, out, *block)

        ValidInputs[string(out.To)] = inputs
      }
    }

  }
  if block.Miner_transaction != nil {
    if block.Miner_transaction.Outputs != nil {
      out := block.Miner_transaction.Outputs[0]
      if inList, ok := ValidInputs[string(out.To)]; ok {

        ValidInputs[string(out.To)] = AddInput(inList, out, *block)

      } else {

        inputs := make([]input, 0)
        inputs = AddInput(inputs, out, *block)
        ValidInputs[string(out.To)] = inputs
      }
    }

  }

}
