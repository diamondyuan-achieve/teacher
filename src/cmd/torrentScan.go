package main

import (
	"path/filepath"
	"os"
	"fmt"
	"torrent"
	"sql"
	"dmmUtils"
	"strings"
)

func main() {
	getFileList("/Users/diamondyuan/python3/spider1024/yuanchuang")
}

func getFileList(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fmt.Println(path)
		if strings.Contains(path, ".torrent") {
			teacher := torrentUtils.GetTorrentNameList(path)
			fmt.Println(teacher)
			learn(teacher)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func learn(teacher torrentUtils.Teacher) {
	for index := 0; index < len(teacher.Files); index++ {
		if teacher.Files[index].Size/1024/1024 > 200 {
			medata, err := dmmUtils.Search(teacher.Files[index].Name)
			tracherSql.Test(medata)
			if err == nil {
				tracherSql.SaveTorrent(medata, teacher)
			}
		}
	}

}
