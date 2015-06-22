package bounqlite

import (
	"fmt"
	bolt "github.com/boltdb/bolt"
	"io/ioutil"
)

func Run() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
