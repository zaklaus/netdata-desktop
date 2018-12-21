package main

import (
	"flag"

	"github.com/zserge/webview"
)

func main() {
	// Fetch the connection string (or use default one)
	connString := flag.String("conn", "http://localhost:19999", "connection string to netdata")
	flag.Parse()

	// Open a window navigating to our connection string
	webview.Open("NetData Desktop", *connString, 1280, 720, true)
}
