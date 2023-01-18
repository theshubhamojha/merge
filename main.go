package main

import (
	"sync"

	"github.com/megre/server"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go server.StartHTTPServer()

	wg.Wait()
}
