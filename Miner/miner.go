package main

import (
  "time"
  "os"
  "io/ioutil"
  "fmt"
  "bytes"
)

func mine() {
  var username string = GetUserNameFromAuthToken(oAuthToken)
  var previousHash []byte
  if _, err := os.Stat("previous-hash"); os.IsNotExist(err) {

    commit := GetLatestCommitForUser(username, oAuthToken)
    if (commit != nil) {
      previousHash = []byte(commit.ID)
      ioutil.WriteFile("previous-hash", previousHash, 0644)
      fmt.Printf("Found New Commit: %v\n With Message: %s", previousHash, commit.Message)
    } else {
      ioutil.WriteFile("previous-hash", make([]byte, 0), 0644)
    }
    CreateBlock(previousHash)
  } else {
    previousHash, _ = ioutil.ReadFile("previoius-hash")
  }

  for ;; {
    time.Sleep(time.Minute * 2)
    commit := GetLatestCommitForUser(username, oAuthToken)
    hash := []byte(commit.ID)
    if (bytes.Compare(hash, previousHash) != 0) {
      fmt.Printf("Found New Commit: %v\n With Message: %s", previousHash, commit.Message)
      previousHash = hash
      ioutil.WriteFile("previous-hash", previousHash, 0644)
      CreateBlock(hash)
    }
  }

}

