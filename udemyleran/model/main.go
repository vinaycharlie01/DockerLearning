package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Favorite is a struct to represent a favorite in the database.
type Favorite struct {
	ID   bson.ObjectId `bson:"_id"`
	Type string
	Name string
	URL  string
}

// Collection returns the MongoDB collection for this struct.
func (f *Favorite) Collection() *mgo.Collection {
	return db.C("favorites")
}

// Save inserts or updates a Favorite in the database.
func (f *Favorite) Save() error {
	if f.ID == "" {
		f.ID = bson.NewObjectId()
	}
	_, err := f.Collection().UpsertId(f.ID, f)
	return err
}

// FindFavorites returns all favorites from the database.
func FindFavorites() ([]*Favorite, error) {
	var favorites []*Favorite
	err := db.C("favorites").Find(nil).All(&favorites)
	return favorites, err
}

var db *mgo.Database

// Connect connects to the MongoDB database.
func Connect(url string) error {
	session, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	db = session.DB("")
	return nil
}
