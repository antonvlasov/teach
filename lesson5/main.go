package main

import "time"

const (
	port = 4001
)

func main() {
	// go serveUDP()
	// go doPingsUDP()

	go serveTCP()
	time.Sleep(500 * time.Millisecond)
	go doPingsTCP()

	select {}
}
