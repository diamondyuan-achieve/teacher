package main

import (
	"fmt"
	"github.com/shiyanhui/dht"
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
