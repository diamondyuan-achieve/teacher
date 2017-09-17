package dmmUtils

import (
	"log"
	"time"
	"github.com/junzh0u/opendmm"
	"errors"
)

var movieMeta opendmm.MovieMeta

func Search(query string) (opendmm.MovieMeta,error) {
	log.Println("start search " + query)
	match := opendmm.Search(query)
	select {
	case meta, ok := <-match:
		if ok {
			return meta,nil
		} else {
			log.Println("Not found")
		}
	case <-time.After(5 * time.Second):
		log.Println("Time out")
	}
	return movieMeta,errors.New("not found")
}