package main

import (
	"net/http"
	"time"
)

// Default go http client doesn't specify a timeout,
// needs a wrapper to specify the client.
var netClient = &http.Client{
	Timeout: time.Second * 10,
}
