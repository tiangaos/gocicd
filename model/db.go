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

func putKVFunc(bucket string, k, v []byte) func(tx *nutsdb.Tx) error {
	return func(tx *nutsdb.Tx) error {
		if err := tx.Put(bucket, k, v, 0); err != nil {
			return err
		}
		return nil
	}
}
