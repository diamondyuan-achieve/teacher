package tracherSql

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/junzh0u/opendmm"
	"time"
	"strings"
	"encoding/json"
	"log"
	"github.com/DiamondYuan/teacher/config"
	"github.com/DiamondYuan/teacher/torrent"
)

var db *sql.DB

func init() {
	cfg, _ := tracherConfig.Load()
	db, _ = sql.Open("mysql", cfg.SQL)
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
}

func Test(meta opendmm.MovieMeta) {
	if Check(meta) != 0 {
		updateCount(meta)
		return
	}
	stmt, err := db.Prepare("" +
		"INSERT av " +
		"SET " +
		"code=?, " +
		"title=?, " +
		"cover_image=?, " +
		"maker=?, " +

		"actresses=?," +
		"genres=?," +
		"directors=?," +
		"created_at=?," +
		"modified_at=?," +
		"release_date=?," +
		"count = 0")

	checkErr(err)

	stmt.Exec(
		meta.Code,
		meta.Title,
		meta.CoverImage,
		meta.Maker,
		strings.Join(meta.Actresses, ","),
		strings.Join(meta.Genres, ","),
		strings.Join(meta.Directors, ","),
		time.Now(),
		time.Now(),
		meta.ReleaseDate)
}

func Check(meta opendmm.MovieMeta) int {
	rows, err := db.Query("SELECT count(*) as count FROM av WHERE code = ?", meta.Code)
	checkErr(err)
	return checkCount(rows)
}

func updateCount(meta opendmm.MovieMeta) int {
	rows, err := db.Query("UPDATE av SET  count = count + 1 WHERE code = ?", meta.Code)
	checkErr(err)
	return checkCount(rows)
}

func SaveTorrent(meta opendmm.MovieMeta, teacher torrentUtils.Teacher) {
	content, _ := json.MarshalIndent(teacher, "", "")
	stmt, err := db.Prepare("" +
		"INSERT av_torrent " +
		"SET " +
		"code=?, " +
		"name=?, " +
		"magnet=?, " +
		"content=?," +
		"created_at=?")
	checkErr(err)
	stmt.Exec(
		meta.Code,
		teacher.Name,
		teacher.Magnet,
		content,
		time.Now(),
	)
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
