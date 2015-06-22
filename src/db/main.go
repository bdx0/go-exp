package db

import (
	bolt "bolt"
	"log"
)

type Post struct {
	Created time.Time
	Title   string
	Content string
}

// Run the example
func Run() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("posts"))
		if err != nil {
			return err
		}
		return b.Put([]byte("2015-01-01"), []byte("My New Year post"))
	})
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))
		v := b.Get([]byte("2015-01-01"))
		fmt.Printf("%sn", v)
		return nil
	})
	defer db.Close()
}
