package main

import (
	"bufio"
	"os"
)

var oAuthToken string
func main() {

  oAuthToken = os.Args[2]
	// Start the server
	server := StartHTTPServer()
	defer ShutdownHTTPServer(server)

	// Load always-on peers
	alwaysOnPeers = LoadAlwaysOnPeers()
	livePeers = FetchLivePeers()

	// Start mining
	go mine()

	// Listen for a command from the front-end
	var text string
	for text != "shutdown\n" {
		reader := bufio.NewReader(os.Stdin)
		text, _ = reader.ReadString('\n')
	}

}
