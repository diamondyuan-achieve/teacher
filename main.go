package main

import (
	"gopkg.in/macaron.v1"
	"dmmUtils"
	"encoding/json"
)


func main() {
	m := macaron.Classic()
	m.Get("/av/*", func(ctx *macaron.Context) {
		medata,err :=dmmUtils.Search(ctx.Req.URL.Path[4:])
		if err != nil {
			ctx.Resp.Write([]byte(err.Error()))
			return
		}
		metaJson, _ := json.MarshalIndent(medata, "", "\t")
		ctx.Resp.Write([]byte(metaJson))
	})
	m.Run(9091)
}
