package utilities

import (
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
	"net/http"
)

func GetDB(r *http.Request) *mgo.Database {
	db := context.Get(r, "db")
	return db.(*mgo.Database)
}

func SetDB(r *http.Request, db *mgo.Database) {
	context.Set(r, "db", db)

}
