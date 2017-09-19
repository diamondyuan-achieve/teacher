package tracherSql

import (
	"time"
)

func SaveData(hash []byte,content []byte,torrent []byte) {
	stmt, err := db.Prepare("" +
		"INSERT magent_table " +
		"SET " +
		"info_hash=?, " +
		"torrent=?, " +
		"meta_data_info=?" +
		"created_at=?")
	checkErr(err)
	stmt.Exec(
		hash,content,torrent,
		time.Now(),
	)
}
