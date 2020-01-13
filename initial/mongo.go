package initial

import "gopkg.in/mgo.v2"

const MONGO_DB_NAME = "tilde"

var dataBase *mgo.Database

func init() {
	address := "localhost:27017"
	session, err := mgo.Dial(address)
	if err != nil {
		panic(err)
	}
	dataBase = session.DB(MONGO_DB_NAME)
	if err != nil {
		panic(err)
	}
}
