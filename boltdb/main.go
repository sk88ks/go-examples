package main

import (
	"github.com/boltdb/bolt"
	"github.com/k0kubun/pp"
)

var db *bolt.DB

const ()

func Set(database, key, value []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(database)
		if err != nil {
			return err
		}

		return b.Put(key, value)
	})
}

func SetMulti(databaseP, databaseC, key []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		bp := tx.Bucket(databaseP)
		if bp == nil {
			return nil
		}

		val := bp.Get(key)
		pp.Println(val)

		b, err := tx.CreateBucketIfNotExists(databaseC)
		if err != nil {
			return err
		}

		return b.Put(key, val)
	})
}

func main() {
	boltdb, err := bolt.Open("./bolt.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer boltdb.Close()

	db = boltdb

	keys := []string{
		"test0",
		"test1",
		"test2",
		"test3",
		"test4",
		"test5",
	}
	values := []string{
		"test0",
		"test1",
		"test2",
		"test3",
		"test4",
		"test5",
	}

	//db1 := []byte("test1")
	//db2 := []byte("test2")

	for i := range keys {
		//go func() {
		//err := Set([]byte("test"+strconv.Itoa(i)), []byte(keys[i]), []byte(values[i]))
		err := Set([]byte("testP"), []byte(keys[i]), []byte(values[i]))
		pp.Println(err)
		//}()
	}

	for i := range keys {
		err := SetMulti([]byte("testP"), []byte("testC"), []byte(keys[i]))
		pp.Println(err)
	}

	db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("testC"))

		for i := range keys {
			pp.Println(string(b.Get([]byte(keys[i]))))
		}
		return nil
	})
}
