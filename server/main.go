package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.Ltime | log.Llongfile)
}

func main() {
	server := server{
		exitCh: make(chan struct{}, 1),
		wg:     new(sync.WaitGroup),
	}
	if err := server.run(acceptSignals()); err != nil {
		log.Fatalln(err)
	}
}
