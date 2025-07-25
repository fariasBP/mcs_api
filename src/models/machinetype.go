package models

import (
	"mcs_api/src/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type MachineType struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.DateTime `json:"updated_at" bson:"updated_at"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
}

func CreateMachineType(name, description string) error {
	newMachineType := &MachineType{
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now().UTC()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now().UTC()),
		Name:        name,
		Description: description,
	}
	ctx, client, coll := config.ConnectColl("machinetype")
	defer client.Disconnect(ctx)
	_, err := coll.InsertOne(ctx, newMachineType)
	return err
}

func GetMachineTypeById(id string) (*MachineType, error) {
	ctx, client, coll := config.ConnectColl("machinetype")
	defer client.Disconnect(ctx)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	machineType := &MachineType{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(machineType)
	return machineType, err
}

func ExistsMachineType(name string) bool {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machinetype")
	defer client.Disconnect(ctx)
	// consultando
	machineType := &MachineType{}
	err := coll.FindOne(ctx, bson.M{"name": name}).Decode(machineType)
	return err == nil
}

func ExistsMachineTypeById(idStr string) bool {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machinetype")
	defer client.Disconnect(ctx)
	// obteniendo id
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return false
	}
	// consultando
	machineType := &MachineType{}
	err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(machineType)
	return err == nil
}

func GetMachineTypes(search string, limit, page int) ([]MachineType, int64, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machinetype")
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
	var machinetypes []MachineType
	if err = cursor.All(ctx, &machinetypes); err != nil {
		return nil, 0, err
	}
	return machinetypes, count, nil
}

func UpdateMachineType(machineType *MachineType) error {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machinetype")
	defer client.Disconnect(ctx)
	// actualizando
	machineType.UpdatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	_, err := coll.UpdateOne(ctx, bson.M{"_id": machineType.ID}, bson.M{"$set": machineType})
	return err
}
