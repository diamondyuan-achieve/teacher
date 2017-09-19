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

	w := dht.NewWire(65536, 1024, 256)

	go func() {
		for resp := range w.Response() {
			fmt.Println(resp.IP+" "+string(resp.Port))
			fmt.Println(resp.Request.IP+" "+string(resp.Request.Port))
			fmt.Println(resp.InfoHash)
		}
	}()
	go w.Run()

	config := dht.NewCrawlConfig()
	config.OnAnnouncePeer = func(infoHash, ip string, port int) {
		fmt.Println(ip,port)
		w.Request([]byte(infoHash), ip, port)
	}
	d := dht.New(config)
	d.Run()
}
