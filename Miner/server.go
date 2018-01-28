package main

import (
	"log"
	"net/http"
)

// Starts the HTTP server
func StartHTTPServer() *http.Server {

	// Binds the port
	srv := &http.Server{Addr: ":8081"}

	// Reroute the functions
	http.HandleFunc("/getPeers", GetPeersReq)
	http.HandleFunc("/getBlockchain", GetBlockchainReq)
	http.HandleFunc("/authorizeBlock", AuthorizeBlockReq)

	// Start the listening thread
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	// returning reference so caller can call Shutdown()
	return srv
}

// Shutdown the HTTP server
func ShutdownHTTPServer(server *http.Server) {

	// Try and shutdown the server
	if err := server.Shutdown(nil); err != nil {
		log.Fatal(err)
	}

}
