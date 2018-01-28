package main

import (
  "time"
  "os"
  "io/ioutil"
  "fmt"
  "bytes"
  "container/list"
)

var TransactionQueue *list.List

func mine() {

  TransactionQueue = list.New()

  var username string = GetUserNameFromAuthToken(oAuthToken)
  fmt.Println("Welcome, " + username)
  fmt.Println("Mining will start in 2 minutes")
  var previousHash []byte
  if _, err := os.Stat("previous-hash"); os.IsNotExist(err) {

    commit := GetLatestCommitForUser(username, oAuthToken)
    if commit != nil {
      previousHash = []byte(commit.ID)
      ioutil.WriteFile("previous-hash", previousHash, 0644)
      fmt.Printf("Found New Commit: %v\n With Message: %s", previousHash, commit.Message)
    } else {
      ioutil.WriteFile("previous-hash", make([]byte, 0), 0644)
    }
    CreateBlock(previousHash)
  } else {
    previousHash, _ = ioutil.ReadFile("previous-hash")
  }

  for ;; {

    currentIndex := CurrentChain.getLatestBlock().Index
    time.Sleep(time.Second * 10)
    commit := GetLatestCommitForUser(username, oAuthToken)
    hash := []byte(commit.ID)

    if bytes.Compare(hash, previousHash) == 0 {
      if currentIndex == CurrentChain.getLatestBlock().Index {
        fmt.Printf("Found New Commit: %v\n With Message: %s - Mining its block!", hash, commit.Message)
        previousHash = hash
        ioutil.WriteFile("previous-hash", previousHash, 0644)
        if CreateBlock(hash) {
          SendBlock(CurrentChain.getLatestBlock())
        }
      }
    }
  }

}


