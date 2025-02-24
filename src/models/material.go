package models

import (
	"mcs_api/src/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Material struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SeriviceId string             `json:"service_id" bson:"service_id,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Price      int                `json:"price" bson:"price,omitempty"`
	Number     int                `json:"number" bson:"number,omitempty"`
}

func CreateMaterial(idService, name string, price int, number int) error {
	// conectando a la bbdd
	ctx, client, coll := config.ConnectColl("materials")
	defer client.Disconnect(ctx)
	// insertando
	newMaterial := &Material{
		SeriviceId: idService,
		Name:       name,
		Price:      price,
		Number:     number,
	}
	_, err := coll.InsertOne(ctx, newMaterial)

	return err
}

func ExistsMaterialById(id string) bool {
	ctx, client, coll := config.ConnectColl("materials")
	defer client.Disconnect(ctx)
	// obteniendo id
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	// buscando material
	material := &Material{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(material)

	return err == nil
}

func UpdateMaterial(idMaterial, name string, price, number int) error {
	// conectando a la bbdd
	ctx, client, coll := config.ConnectColl("materials")
	defer client.Disconnect(ctx)
	// obteniendo id
	id, err := primitive.ObjectIDFromHex(idMaterial)
	if err != nil {
		return err
	}
	// actualizando material
	update := bson.M{"$set": bson.M{"name": name, "price": price, "number": number}}
	_, err = coll.UpdateOne(ctx, bson.M{"_id": id}, update)

	return err
}

func GetMaterialsByServiceId(idService string) ([]Material, error) {
	// conectando a la bbdd
	ctx, client, coll := config.ConnectColl("materials")
	defer client.Disconnect(ctx)
	// creando parametros de consulta
	query := bson.M{"service_id": idService}
	// consultando
	cursor, err := coll.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	// modelando datos
	var materials []Material
	if err = cursor.All(ctx, &materials); err != nil {
		return nil, err
	}

	return materials, nil
}
