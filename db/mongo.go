package db

import (
	"flag"

	"labix.org/v2/mgo"
)

var (
	mongoURL = flag.String("mongourl", "127.0.0.1", "")
	mongoDB  = flag.String("mongodb", "recipes", "")

	session *mgo.Session
)

func Connect() error {
	var err error
	session, err = mgo.Dial(*mongoURL)
	if err != nil {
		return err
	}

	Recipes = session.DB(*mongoDB).C("scraper")
	return nil
}

func Close() {
	session.Close()
}
