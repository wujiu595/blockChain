package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

/*

db, err := bolt.Open("my.db", 0600, nil)
if err != nil {
	log.Fatal(err)
}
defer db.Close()


err := db.Update(func(tx *bolt.Tx) error {
	...
	return nil
})

err := db.View(func(tx *bolt.Tx) error {
	...
	return nil
})

err := db.Batch(func(tx *bolt.Tx) error {
	...
	return nil
})

//transaction
tx, err := db.Begin(true)
if err != nil {
    return err
}
defer tx.Rollback()

_, err := tx.CreateBucket([]byte("MyBucket"))
if err != nil {
    return err
}

// Commit the transaction and check for error.
if err := tx.Commit(); err != nil {
    return err
}
*/

func main()  {
	db, err := bolt.Open("./my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("answer"))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
}
