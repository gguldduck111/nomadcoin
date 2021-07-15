package db

import (
	"github.com/boltdb/bolt"
	"github.com/gguldduck111/nomadcoin/Util"
)

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		//init db
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		Util.HandleErr(err)
		db = dbPointer

		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
			Util.HandleErr(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		Util.HandleErr(err)
	}
	return db
}
