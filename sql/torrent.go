package tracherSql

import (
	"time"
	"fmt"
	"strings"
)

type HashIp  struct{
	Hash string
	Ip string
}



func SaveData(hash []byte, content []byte, torrent []byte) {
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
		hash, content, torrent,
		time.Now(),
	)
}


func SaveHash(unsavedRows []HashIp) {
	fmt.Println(len(unsavedRows))
	valueStrings := make([]string, 0, len(unsavedRows))
	valueArgs := make([]interface{}, 0, len(unsavedRows) * 2)
	for _, post := range unsavedRows {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, post.Hash)
		valueArgs = append(valueArgs, post.Ip)
	}
	stmt := fmt.Sprintf("INSERT INTO Hash (hash, ip) VALUES %s", strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	if err!=nil {
		fmt.Println(err)
	}
}