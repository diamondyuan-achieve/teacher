package tracherSql

import (
	"time"
	"fmt"
)

func SaveData(hash []byte,content []byte,torrent []byte) {
	stmt, err := db.Prepare("" +
		"INSERT magent_table " +
		"SET " +
		"info_hash=?, " +
		"torrent=?, " +
		"meta_data_info=?," +
		"created_at=?")
	if err != nil {
		fmt.Print(err.Error())
	}
	stmt.Exec(
		hash,content,torrent,
		time.Now(),
	)
}


func SaveHash(hash string,ip string) {
	stmt, err := db.Prepare(
		"INSERT Hash " +
			"SET hash = ?, "+
			"ip = ?")
	if err != nil {
		fmt.Print(err.Error())
	}
	_,err2 := stmt.Exec(
		hash,
		ip,
	)
	if err != nil {
		fmt.Println(err2)
	}
}
