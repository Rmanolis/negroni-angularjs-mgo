package main

import (
	"./utilities"
	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {

	store := cookiestore.New([]byte("secretkey789"))
	router := LoadRoutes()
	n := negroni.Classic()
	static := negroni.NewStatic(http.Dir("static"))
	static.Prefix = "/static"
	n.Use(static)
	n.Use(negroni.HandlerFunc(MgoMiddleware))
	n.Use(sessions.Sessions("global_session_store", store))
	n.UseHandler(router)
	n.Run(":3000")

}

func MgoMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		panic(err)
	}

	reqSession := session.Clone()
	defer reqSession.Close()
	db := reqSession.DB("test")
	utilities.SetDB(r, db)
	next(rw, r)
}
