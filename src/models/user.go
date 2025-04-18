package models

import (
	"mcs_api/src/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Permission int

const (
	Super    Permission = 3 // crea lee modifica y elimina
	Admin    Permission = 2 // crea lee y modifica
	Operator Permission = 1 // crea y lee
	Public   Permission = 0 // solo lee
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
	Nick      string             `json:"nick" bson:"nick,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty"`
	Lname     string             `json:"lname" bson:"lname,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Pwd       string             `json:"pwd" bson:"pwd,omitempty"`
	Bth       string             `json:"bth" bson:"bth,omitempty"`
	Perm      Permission         `json:"perm" bson:"perm,omitempty"`
}

func CreateSuperUser(nick, name, lname, email, pwd, bth string) error {
	newUser := &User{
		Nick:      nick,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now().UTC()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now().UTC()),
		Name:      name,
		Lname:     lname,
		Email:     email,
		Pwd:       pwd,
		Bth:       bth,
		Perm:      Super,
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
	err := coll.FindOne(ctx, bson.M{"perm": Super}).Decode(user)
	return err == nil
}

func CreateUser(nick, name, lname, email, pwd, bth string, perm Permission) error {
	newUser := &User{
		Nick:      nick,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now().UTC()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now().UTC()),
		Name:      name,
		Lname:     lname,
		Email:     email,
		Pwd:       pwd,
		Bth:       bth,
		Perm:      perm,
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

func GetUserAndPwd(identifier string) (*User, error) {
	ctx, client, coll := config.ConnectColl("users")
	defer client.Disconnect(ctx)
	user := &User{}
	err := coll.FindOne(ctx, bson.M{"$or": []bson.M{
		{"nick": identifier},
		{"email": identifier},
	}}).Decode(user)
	return user, err
}

func GetUserById(id string) (*User, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("users")
	defer client.Disconnect(ctx)
	// obteniendo id
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	// buscando usuario
	user := &User{}
	opts := options.FindOne().SetProjection(bson.M{"pwd": 0})
	err = coll.FindOne(ctx, bson.M{"_id": idObj}, opts).Decode(user)

	return user, err
}

func UpdateUser(user *User) error {
	ctx, client, coll := config.ConnectColl("users")
	defer client.Disconnect(ctx)
	idObj, err := primitive.ObjectIDFromHex(user.ID.Hex())
	if err != nil {
		return err
	}
	user.ID = idObj
	user.UpdatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	_, err = coll.UpdateOne(ctx, bson.M{"_id": idObj}, bson.M{"$set": user})
	if user.Perm == Public {
		_, err = coll.UpdateOne(ctx, bson.M{"_id": idObj}, bson.M{"$set": bson.M{"perm": Public}})
	}
	return err
}
