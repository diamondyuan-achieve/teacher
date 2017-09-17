package main

import (
	"gopkg.in/macaron.v1"
	"encoding/json"
	"github.com/GeorgeYuen/teacher/config"
	"github.com/GeorgeYuen/teacher/sql"
	"github.com/GeorgeYuen/teacher/dmmUtils"
)


func main() {
	cfg,_ := tracherConfig.Load()
	m := macaron.Classic()
	m.Get("/av/*", func(ctx *macaron.Context) {
		medata,err :=dmmUtils.Search(ctx.Req.URL.Path[4:])
		if err != nil {
			ctx.Resp.Write([]byte(err.Error()))
			return
		}
		go tracherSql.Test(medata)
		metaJson, _ := json.MarshalIndent(medata, "", "\t")
		ctx.Resp.Write([]byte(metaJson))
	})
	m.Run(cfg.Port)
}
