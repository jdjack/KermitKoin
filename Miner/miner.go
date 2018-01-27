package main

import (
  "time"
  "os"
  "io/ioutil"
)

func mine() {
  var username string = getUserNameFromAuthTokken(oAuthToken)
  if _, err := os.Stat("previous-hash"); os.IsNotExist(err) {

    ioutil.WriteFile("previous-hash", []byte(GetLatestCommitForUser(username, oAuthToken).ID), 0644)
  }

  var text string
  for text != "shutdown" {
    
  }

}

func checkCommits() {

}
