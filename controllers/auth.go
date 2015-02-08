package controllers

import (
	"../models"
	"../utilities"
	"encoding/json"
	"fmt"
	"github.com/goincremental/negroni-sessions"
	"net/http"
)

type Credentials struct {
	Email    string
	Password string
}

type Auth struct{}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	credentials := new(Credentials)
	err := decoder.Decode(&credentials)
	if err != nil {
		panic(err)
	}

	db := utilities.GetDB(r)
	user := new(models.User)
	err = user.Authenticate(db, credentials.Email, credentials.Password)
	if err == nil {
		session := sessions.GetSession(r)
		session.Set("user_id", user.ID.Hex())
		w.WriteHeader(202)

	} else {
		w.WriteHeader(404)
	}
}

func (a *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	session := sessions.GetSession(r)
	user_id := session.Get("user_id")
	fmt.Println(user_id)
	if user_id == nil {
		w.WriteHeader(403)
		http.Redirect(w, r, "/", 403)

	} else {
		session.Delete("user_id")
		http.Redirect(w, r, "/", 202)
	}

}

func (a *Auth) User(w http.ResponseWriter, r *http.Request) {
	db := utilities.GetDB(r)
	session := sessions.GetSession(r)
	user_id := session.Get("user_id")
	fmt.Println(user_id)
	if user_id == nil {
		w.WriteHeader(403)

	} else {
		user := new(models.User)
		user.Get(db, user_id.(string))
		fmt.Println(user)
		outData, _ := json.Marshal(user)
		w.Write(outData)
	}

}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	data := map[string]string{"name": "", "email": "", "password": ""}
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	db := utilities.GetDB(r)
	user := new(models.User)
	user.NewUser(db, data["name"], data["email"], data["password"])
	fmt.Println(user)

}
