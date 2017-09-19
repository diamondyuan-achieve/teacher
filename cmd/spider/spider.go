package main

import (
	"fmt"
	"github.com/GeorgeYuen/dht"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	config := dht.NewCrawlConfig()
	config.OnAnnouncePeer = func(infoHash, ip string, port int) {
		fmt.Println(infoHash)
	}
	d := dht.New(config)
	d.Run()
}
