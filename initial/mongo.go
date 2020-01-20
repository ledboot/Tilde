package initial

import "gopkg.in/mgo.v2"

const MONGO_DB_NAME = "tilde"
const MONGO_DB_URI = "localhost:27017"

var dataBase *mgo.Database

func GetMongoDB() *mgo.Database {
	return dataBase
}
