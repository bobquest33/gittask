package main

import (
	"fmt"

	"github.com/izqui/helpers"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Id bson.ObjectId `json:"_id" bson:"_id"`

	Email       string `json:"email,omitempty" bson:"email"`
	Username    string `json:"login" bson:"username"`
	Location    string `json:"location,omitempty" bson:"location"`
	AccessToken string `json:"" bson:"token"`
}

type Repo struct {
	Id     bson.ObjectId `bson:"_id"`
	UserId bson.ObjectId `bson:"user_id"`

	GithubId int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Language string `json:"language,omitempty"`
}

func CurrentUser(token string) (user *User) {

	userCollection := DB.C("users")

	if err := userCollection.Find(bson.M{"token": token}).One(&user); err != nil {

		panic(err)

	} else {

		return user
	}
}

func (u *User) Update() {

	userCollection := DB.C("users")

	change := mgo.Change{

		ReturnNew: true,
		Update:    helpers.StructToBSONMap(u),
	}
	fmt.Println(change.Update)
	if _, err := userCollection.Find(bson.M{"_id": u.Id}).Apply(change, u); err != nil {

		panic(err)
	}

	return
}

func (u *User) GetRepos() (repos []Repo) {

	repoCollection := DB.C("repos")

	if err := repoCollection.Find(bson.M{"user_id": u.Id}).All(&repos); err != nil {

		panic(err)
	} else {
		return
	}
}
