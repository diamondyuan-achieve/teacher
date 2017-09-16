package torrentUtils

import (
	"github.com/anacrolix/torrent/metainfo"
)

type Teacher struct {
	Magnet       string
	Name string
	Files []TFile
}

type TFile struct {
	Name string
	Size int64
}

func GetTorrentNameList(path string) Teacher {
	mi, _ := metainfo.LoadFromFile(path)
	info, _ := mi.UnmarshalInfo()
	sl := make([]TFile, len(info.Files))
	for index := 0; index < len(info.Files); index++ {
		filename := info.Files[index].DisplayPath(&info)
		sl[index] = TFile{
			Name:filename,
			Size:info.Files[index].Length,
		}
	}
	return Teacher{
		Files:sl,
		Name:info.Name,
		Magnet:"magnet:?xt=urn:btih:" + mi.HashInfoBytes().String(),
	}
}
