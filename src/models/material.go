package models

import (
	"mcs_api/src/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Material struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt  primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt  primitive.DateTime `json:"updated_at" bson:"updated_at"`
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
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now().UTC()),
		UpdatedAt:  primitive.NewDateTimeFromTime(time.Now().UTC()),
		Name:       name,
		Price:      price,
		Number:     number,
	}
	_, err := coll.InsertOne(ctx, newMaterial)

	return err
}

func GetMaterialById(id string) (*Material, error) {
	// conectando a la bbdd
	ctx, client, coll := config.ConnectColl("materials")
	defer client.Disconnect(ctx)
	// obteniendo id
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	// buscando material
	material := &Material{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(material)

	return material, err
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

func UpdateMaterial(material *Material) error {
	// conectando a la bbdd
	ctx, client, coll := config.ConnectColl("materials")
	defer client.Disconnect(ctx)
	// actualizando material
	material.UpdatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	_, err := coll.UpdateOne(ctx, bson.M{"_id": material.ID}, bson.M{"$set": material})

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
