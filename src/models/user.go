package models

import (
	"mcs_api/src/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.DateTime `json:"updated_at" bson:"updated_at"`
	Nick        string             `json:"nick" bson:"nick,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Lname       string             `json:"lname" bson:"lname,omitempty"`
	Email       string             `json:"email" bson:"email,omitempty"`
	Pwd         string             `json:"pwd" bson:"pwd,omitempty"`
	Bth         string             `json:"bth" bson:"bth,omitempty"`
	Permissions []int8             `json:"permissions" bson:"permissions,omitempty"`
}

func CreateSuperUser(nick, name, lname, email, pwd, bth string) error {
	newUser := &User{
		Nick:        nick,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now().UTC()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now().UTC()),
		Name:        name,
		Lname:       lname,
		Email:       email,
		Pwd:         pwd,
		Bth:         bth,
		Permissions: []int8{0},
	}
	ctx, client, coll := config.ConnectColl("users")
	defer client.Disconnect(ctx)
	_, err := coll.InsertOne(ctx, newUser)
	return err
}

func ExistsSuperUser() bool {
	user := &User{}
	ctx, client, coll := config.ConnectColl("users")
	defer client.Disconnect(ctx)
	err := coll.FindOne(ctx, bson.M{"permissions": bson.M{
		"$all": []int8{0},
	}}).Decode(user)
	return err == nil
}

func CreateUser(nick, name, lname, email, pwd, bth string) error {
	newUser := &User{
		Nick:        nick,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now().UTC()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now().UTC()),
		Name:        name,
		Lname:       lname,
		Email:       email,
		Pwd:         pwd,
		Bth:         bth,
		Permissions: []int8{1},
	}
	ctx, client, coll := config.ConnectColl("users")
	defer client.Disconnect(ctx)
	_, err := coll.InsertOne(ctx, newUser)
	return err
}

func ExistsUser(nick, email string) bool {
	ctx, client, coll := config.ConnectColl("users")
	defer client.Disconnect(ctx)
	user := &User{}
	err := coll.FindOne(ctx, bson.M{"$or": []bson.M{
		{"nick": nick},
		{"email": email},
	}}).Decode(user)
	return err == nil
}

func ExistsUserById(id string) bool {
	ctx, client, coll := config.ConnectColl("users")
	defer client.Disconnect(ctx)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	user := &User{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(user)
	return err == nil
}

func GetUser(identifier string) (*User, error) {
	ctx, client, coll := config.ConnectColl("users")
	defer client.Disconnect(ctx)
	user := &User{}
	err := coll.FindOne(ctx, bson.M{"$or": []bson.M{
		{"nick": identifier},
		{"email": identifier},
	}}).Decode(user)
	return user, err
}
