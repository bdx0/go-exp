package bounqlite

import (
	"fmt"
	unqlite "github.com/nobonobo/unqlitego"
	"io/ioutil"
)

func Run() {
	var db *unqlite.Database
	// var src []byte
	// src = []byte("value")

	f, err := ioutil.TempFile("", "sample.db")
	if err != nil {
		panic(err)
	}
  fmt.Println(f.Name())
	db, _ = unqlite.NewDatabase(f.Name())
  err = db.Begin()
	for i := 0; i < 100; i++ {
		db.Store([]byte(fmt.Sprintf("%d", i)), []byte("abcdefghijklmnopabcdefghijklmnopabcdefghijklmnopabcdefghijklmnop"))
	}
  err = db.Close()
}
