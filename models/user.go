package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

func (u *User) NewUser(db *mgo.Database, name string, email string, password string) {
	u.Name = name
	u.Email = email
	u.ID = bson.NewObjectId()
	h := md5.New()
	io.WriteString(h, password)
	u.Password = hex.EncodeToString(h.Sum(nil))
	c := db.C("user")
	c.Insert(&u)
}

func (u *User) Get(db *mgo.Database, id string) error {
	if bson.IsObjectIdHex(id) {
		return db.C("user").FindId(bson.ObjectIdHex(id)).One(&u)
	} else {
		return errors.New("It is not ID")
	}
}

func (u *User) Authenticate(db *mgo.Database, email string, password string) error {
	h := md5.New()
	io.WriteString(h, password)
	hex_password := hex.EncodeToString(h.Sum(nil))
	err := db.C("user").Find(map[string]string{
		"password": hex_password,
		"email":    email,
	}).One(&u)
	return err
}
