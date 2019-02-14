package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

var db *mgo.Session

func main() {

}

func dialdb() error {
	var err error
	log.Println("dialing mongodb: localhost:27018")
	db, err := mgo.Dial("localhost:27018")
	return err
}

func closedb() {
	db.Close()
	log.Println("closed database connection")
}
