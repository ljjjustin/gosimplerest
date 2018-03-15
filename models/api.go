package models

import (
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbname = "gosimplerest.db"
)

var (
	db     *xorm.Engine
	tables []interface{}
)

func init() {
	var err error
	// create database connection
	db, err = xorm.NewEngine("sqlite3", dbname)
	if err != nil {
		log.Fatal(err)
	}
	db.ShowSQL(true)

	tables = append(tables, new(User), new(Group))
	// register and update database models
	db.Sync2(tables...)
}
