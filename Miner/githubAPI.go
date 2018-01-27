package main

import (
  "fmt"
  "net/http"
  "bytes"
  "io/ioutil"
  "encoding/json"
)

type Commit struct {
  ID string
  Message string
  RepoName string
}
//
//func main() {
//  tokenPt1 := "93b1d16bf647dea"
//  tokenPt2 := "360b566e34eddbbad07b89f7a"
//  commit := GetLatestCommitForUser("jdjack", tokenPt1+tokenPt2)
//  CheckCommitExistanceForUser("jdjack", commit.ID, tokenPt1+tokenPt2)
//  GetUserNameFromAuthToken(tokenPt1+tokenPt2)
//}

////func main() {
//  tokenPt1 := "93b1d16bf647dea"
//  tokenPt2 := "360b566e34eddbbad07b89f7a"
//  GetLatestCommitForUser("jdjack", tokenPt1+tokenPt2)
//}

// Jacks token: - 67bbbbc9234c47bd77b1720c493ed436801d6cd4
func GetLatestCommitForUser(username string, oAuthToken string) *Commit {
  url := "https://api.github.com/graphql"

  var jsonStr = []byte(`{"query" : "query { viewer { repositories(last: 100) {nodes { name, ref(qualifiedName: \"master\") {target {... on Commit {history(first: 100, since: \"2018-01-20T01:01:00\") {edges {node {author {user {login }} committedDate, oid, messageHeadline}}}}}}}}}}"}`)
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
  req.Header.Set("Authorization", "Bearer " + oAuthToken)
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)

  var dat map[string]interface{}

  if err := json.Unmarshal(body, &dat); err != nil {
    panic(err)
  }

  mostRecentTime := "1970-01-01T01:01:01Z"
  var mostRecentCommit *Commit = nil

  commitList := dat["data"].(map[string]interface{})["viewer"].(map[string]interface{})["repositories"].(map[string]interface{})["nodes"].([]interface{})
  for _, commit := range commitList {
    castedCommit := commit.(map[string]interface{})
    repoName := castedCommit["name"].(string)
    edges := castedCommit["ref"].(map[string]interface{})["target"].(map[string]interface{})["history"].(map[string]interface{})["edges"].([]interface{})
    if len(edges) == 0 {
      continue
    }

    for _, edge := range edges {
      castedEdge := edge.(map[string]interface{})
      edgeUsername := castedEdge["node"].(map[string]interface{})["author"].(map[string]interface{})["user"].(map[string]interface{})["login"].(string)
      if username != edgeUsername {
        continue
      }

      date := castedEdge["node"].(map[string]interface{})["committedDate"].(string)
      oid := castedEdge["node"].(map[string]interface{})["oid"].(string)
      message := castedEdge["node"].(map[string]interface{})["messageHeadline"].(string)
      if date > mostRecentTime {
        mostRecentTime = date
        mostRecentCommit = &Commit{oid, message, repoName}
      }
    }

  }

  fmt.Println(mostRecentTime)
  fmt.Println(mostRecentCommit)

  return mostRecentCommit
}


func CheckCommitExistanceForUser(username string, hash string, oAuthToken string) bool {

  url := "https://api.github.com/search/commits?q=hash:" + hash

  req, err := http.NewRequest("GET", url, nil)
  req.Header.Set("Authorization", "Bearer " + oAuthToken)
  req.Header.Set("Accept", "application/vnd.github.cloak-preview+json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)

  var dat map[string]interface{}

  if err := json.Unmarshal(body, &dat); err != nil {
    panic(err)
  }

  foundItems := dat["items"].([]interface{})

  if len(foundItems) == 0 {
    return false
  }

  foundItem := foundItems[0].(map[string]interface{})
  shaHash := foundItem["sha"].(string)
  author := foundItem["author"].(map[string]interface{})["login"].(string)

  found := shaHash == hash && author == username

  fmt.Println(found)

  return found
}

func GetUserNameFromAuthToken(token string) string {

  url := "https://api.github.com/graphql"

  var jsonStr = []byte(`{"query" : "query { viewer { login }}"}`)
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
  req.Header.Set("Authorization", "Bearer " + token)
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)

  var dat map[string]interface{}

  if err := json.Unmarshal(body, &dat); err != nil {
    panic(err)
  }

  name := dat["data"].(map[string]interface{})["viewer"].(map[string]interface{})["login"].(string)

  fmt.Println(name)
  return name
}