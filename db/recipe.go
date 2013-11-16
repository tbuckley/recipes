package db

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var Recipes *mgo.Collection

type Recipe struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Page    string        `bson:"page"`
	Content string        `bson:"content"`
}

func (r *Recipe) Save() error {
	return Recipes.Insert(r)
}

func AllRecipes() ([]Recipe, error) {
	recipes := make([]Recipe, 100)
	err := Recipes.Find(bson.M{}).All(&recipes)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}
