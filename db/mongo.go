package db

import (
	"labix.org/v2/mgo"
)

var session *mgo.Session

func Connect(url, db string) error {
	var err error
	session, err = mgo.Dial(url)
	if err != nil {
		return err
	}

	Recipes = session.DB(db).C("scraper")
	return nil
}

func Close() {
	session.Close()
}
