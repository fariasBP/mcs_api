package models

import (
	"mcs_api/src/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Brand struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
	Name      string             `json:"name" bson:"name,omitempty"`
}

func CreateBrand(name string) error {
	newBrand := &Brand{
		CreatedAt: primitive.NewDateTimeFromTime(time.Now().UTC()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now().UTC()),
		Name:      name,
	}
	ctx, client, coll := config.ConnectColl("brand")
	defer client.Disconnect(ctx)
	_, err := coll.InsertOne(ctx, newBrand)
	return err
}

func ExistsBrand(name string) bool {
	ctx, client, coll := config.ConnectColl("brand")
	defer client.Disconnect(ctx)
	brand := &Brand{}
	err := coll.FindOne(ctx, bson.M{"name": name}).Decode(brand)
	return err == nil
}

func ExistsBrandById(idStr string) bool {
	ctx, client, coll := config.ConnectColl("brand")
	defer client.Disconnect(ctx)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return false
	}
	brand := &Brand{}
	err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(brand)
	return err == nil
}

func GetBrands(name string, limit, page int) ([]Brand, int64, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("brand")
	defer client.Disconnect(ctx)
	// creando parametros de consulta
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(page - 1))
	query := bson.M{}
	if name != "" {
		query = bson.M{"name": primitive.Regex{
			Pattern: `(\s` + name + `|^` + name + `|\w` + name + `\w` + `|` + name + `$` + `|` + name + `\s)`, Options: "i",
		}}
	}
	// consultando cantidad de datos
	count, err := coll.CountDocuments(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	// consultando
	cursor, err := coll.Find(ctx, query, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)
	// modelando datos
	var brands []Brand
	if err = cursor.All(ctx, &brands); err != nil {
		return nil, 0, err
	}
	return brands, count, nil
}
