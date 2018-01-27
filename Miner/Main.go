package main

import (
	"bufio"
	"os"
)

func main() {

	// Start the server
	server := StartHTTPServer()
	defer ShutdownHTTPServer(server)

	// Listen for a command from the front-end
	var text string
	for text != "shutdown\n" {
		reader := bufio.NewReader(os.Stdin)
		text, _ = reader.ReadString('\n')
	}

}
