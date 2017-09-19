package tracherConfig

import (
	"os"
	"log"
	"fmt"
)

func Load() (cfg *Config, err error) {
	cfg = &Config{}
	cfg.Port = defaultConfig.Port
	cfg.SQL = defaultConfig.SQL
	sql := os.Getenv("SQL")
	if sql != ""  {
		cfg.SQL = sql
	}
	log.Println(sql)
	fmt.Println(sql)
	return cfg,nil
}