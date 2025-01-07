package models

import (
	"mcs_api/src/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func AddPermission(idUser, idPermission string) error {
	ctx, client, coll := config.ConnectColl("users")
	defer client.Disconnect(ctx)
	id, err := primitive.ObjectIDFromHex(idUser)
	if err != nil {
		return err
	}
	update := bson.M{"$set": bson.M{"permissions": idPermission}}
	_, err = coll.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}
