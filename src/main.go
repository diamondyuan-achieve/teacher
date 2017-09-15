package main
import (
	"github.com/anacrolix/torrent"
	"fmt"
)



func main() {
	c, _ := torrent.NewClient(nil)
	t,_:=c.AddTorrentFromFile("1.torrent")
	fmt.Println(t.GotInfo())
}

