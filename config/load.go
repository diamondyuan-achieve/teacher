package tracherConfig

import (
	"os"
	"bufio"
	"io"
	"strings"
	"strconv"
	"log"
)

func Load() (cfg *Config, err error) {
	cfg = &Config{}
	cfg.Port = defaultConfig.Port
	cfg.SQL = defaultConfig.SQL
	sql := os.Getenv("SQL")
	initConfig(cfg,"teacher.ini")
	if sql != ""  {
		cfg.SQL = sql
	}
	log.Println(sql)
	return cfg,nil
}


func initConfig(cfg *Config,path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if key == "sql"{
			cfg.SQL = value
		}
		if key == "port" {
			cfg.Port,_ = strconv.Atoi(value)
		}
	}
}