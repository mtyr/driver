package db

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

var DB = newDB()

type dbType struct {
	db *scribble.Driver
}

func newDB() dbType {
	db, err := scribble.New("./db", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	return dbType{db: db}
}

func (d *dbType) Get() *scribble.Driver {
	return d.db
}
