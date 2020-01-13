package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

func main() {
	address := "localhost:27017"
	session, err := mgo.Dial(address)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	colls, err := session.DB("tidle").CollectionNames()
	if err != nil {
		panic(err)
	}
	fmt.Println(colls)
}
