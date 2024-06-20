package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Post struct {
		ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
		From    string        `json:"from" bson:"from"`
		To      string        `json:"to" bson:"to"`
		Message string        `json:"message" bson:"message"`
	}
)
