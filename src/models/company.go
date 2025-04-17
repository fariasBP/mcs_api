package models

import (
	"mcs_api/src/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Company struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.DateTime `json:"updated_at" bson:"updated_at"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Manager     string             `json:"manager" bson:"manager,omitempty"`
	Latitude    string             `json:"latitude" bson:"latitude,omitempty"`
	Longitude   string             `json:"longitude" bson:"longitude,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Contact     string             `json:"contact" bson:"contact,omitempty"`
}

func CreateCompany(name, manager, latitude, longitude, description, contact string) error {
	ctx, client, coll := config.ConnectColl("companies")
	defer client.Disconnect(ctx)
	newCompany := &Company{
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now().UTC()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now().UTC()),
		Name:        name,
		Manager:     manager,
		Latitude:    latitude,
		Longitude:   longitude,
		Description: description,
		Contact:     contact,
	}
	_, err := coll.InsertOne(ctx, newCompany)
	return err
}

func ExistsCompany(idName string) bool {
	ctx, client, coll := config.ConnectColl("companies")
	defer client.Disconnect(ctx)
	company := &Company{}
	err := coll.FindOne(ctx, bson.M{"name": idName}).Decode(company)
	return err == nil
}

func ExistsCompanyById(idStr string) bool {
	ctx, client, coll := config.ConnectColl("companies")
	defer client.Disconnect(ctx)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return false
	}
	company := &Company{}
	err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(company)
	return err == nil
}

func GetCompanies(search string, limit, page int) ([]Company, int64, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("companies")
	defer client.Disconnect(ctx)
	// creando parametros de consulta
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(page - 1))
	query := bson.M{}
	if search != "" {
		query = bson.M{"name": primitive.Regex{
			Pattern: `(\s` + search + `|^` + search + `|\w` + search + `\w` + `|` + search + `$` + `|` + search + `\s)`, Options: "i",
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
	var companies []Company
	if err = cursor.All(ctx, &companies); err != nil {
		return nil, 0, err
	}
	return companies, count, nil
}

func GetCompanyById(id string) (*Company, error) {
	ctx, client, coll := config.ConnectColl("companies")
	defer client.Disconnect(ctx)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	company := &Company{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(company)
	return company, err
}
