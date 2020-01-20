package initial

import "gopkg.in/mgo.v2"

func init() {
	initMongo()
	initPostgres()
}

func initMongo() {
	session, err := mgo.Dial(MONGO_DB_URI)
	if err != nil {
		panic(err)
	}
	dataBase = session.DB(MONGO_DB_NAME)
	if err != nil {
		panic(err)
	}
}

func initPostgres() {

}
