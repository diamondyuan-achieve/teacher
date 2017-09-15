package torrentUtils

import (
	"fmt"
	"path/filepath"
	"os"
	"github.com/anacrolix/torrent/metainfo"
)

func getFileList(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		getTorrentNameList(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

}

func getTorrentNameList(path string) []string {
	mi, _ := metainfo.LoadFromFile(path)
	info, _ := mi.UnmarshalInfo()
	sl := make([]string, len(info.Files))
	fmt.Println("magnet:?xt=urn:btih:" + mi.HashInfoBytes().String())
	for index := 0; index < len(info.Files); index++ {
		filename := info.Files[index].DisplayPath(&info)
		if info.Files[index].Length / 1024/1024 > 200 {
			fmt.Println(filename)
		}
		sl[index] = filename
	}
	return sl
}
