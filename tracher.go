package main

import (
	"encoding/json"
	"github.com/go-macaron/macaron"
	"github.com/DiamondYuan/teacher/config"
	"github.com/DiamondYuan/teacher/sql"
	"github.com/DiamondYuan/teacher/dmmUtils"
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
