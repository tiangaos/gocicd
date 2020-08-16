package model

import (
	"log"
	"github.com/xujiajun/nutsdb"	
)

var db *nutsdb.DB

func init() {
	var err error
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, err = nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
}
