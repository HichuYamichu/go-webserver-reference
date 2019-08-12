package context

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Context struct {
	DB *mongo.Database
}

func NewContext(db *mongo.Database) *Context {
	c := &Context{
		DB: db,
	}
	return c
}
