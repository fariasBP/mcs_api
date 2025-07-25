package models

import (
	"mcs_api/src/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type (
	Protocol struct {
		ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		CreatedAt     primitive.DateTime `json:"created_at" bson:"created_at"`
		UpdatedAt     primitive.DateTime `json:"updated_at" bson:"updated_at"`
		MachineTypeId string             `json:"machine_type_id" bson:"machine_type_id,omitempty"`
		Acronym       string             `json:"acronym" bson:"acronym,omitempty"`
		Name          string             `json:"name" bson:"name,omitempty"`
		Description   string             `json:"description" bson:"description,omitempty"`
	}
	ProtocolRebuild struct {
		ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		CreatedAt       primitive.DateTime `json:"created_at" bson:"created_at"`
		UpdatedAt       primitive.DateTime `json:"updated_at" bson:"updated_at"`
		MachineTypeId   string             `json:"machine_type_id" bson:"machine_type_id,omitempty"`
		MachineTypeName string             `json:"machine_type_name" bson:"machine_type_name,omitempty"`
		Acronym         string             `json:"acronym" bson:"acronym,omitempty"`
		Name            string             `json:"name" bson:"name,omitempty"`
		Description     string             `json:"description" bson:"description,omitempty"`
	}
)

func CreateProtocol(machineTypeId, acronym, name, description string) error {
	newProtocol := &Protocol{
		CreatedAt:     primitive.NewDateTimeFromTime(time.Now().UTC()),
		UpdatedAt:     primitive.NewDateTimeFromTime(time.Now().UTC()),
		MachineTypeId: machineTypeId,
		Acronym:       acronym,
		Name:          name,
		Description:   description,
	}
	ctx, client, coll := config.ConnectColl("protocol")
	defer client.Disconnect(ctx)
	_, err := coll.InsertOne(ctx, newProtocol)
	return err
}

func GetProtocolById(idStr string) (*Protocol, error) {
	ctx, client, coll := config.ConnectColl("protocol")
	defer client.Disconnect(ctx)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, err
	}
	protocol := &Protocol{}
	err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(protocol)
	return protocol, err
}

func ExistsProtocol(machineTypeId, acronym, name string) bool {
	ctx, client, coll := config.ConnectColl("protocol")
	defer client.Disconnect(ctx)
	protocol := &Protocol{}
	query := bson.M{"$and": []bson.M{
		{"machine_type_id": machineTypeId},
		{"$or": []bson.M{
			{"name": name},
			{"acronym": acronym},
		}},
	}}
	err := coll.FindOne(ctx, query).Decode(protocol)

	return err == nil
}

func ExistsProtocolById(idStr string) bool {
	ctx, client, coll := config.ConnectColl("protocol")
	defer client.Disconnect(ctx)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return false
	}
	protocol := &Protocol{}
	err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(protocol)
	return err == nil
}

// existe otro sin contar este
func ExistsOtherProtocol(protocolId, machineTypeId, acronym, name string) bool { //
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("protocol")
	defer client.Disconnect(ctx)
	// obteniendo id
	id, err := primitive.ObjectIDFromHex(protocolId)
	if err != nil {
		return false
	}
	// consultando
	protocol := &Protocol{}
	query := bson.M{"$and": []bson.M{
		{"_id": bson.M{"$ne": id}},
		{"machine_type_id": machineTypeId},
		{"$or": []bson.M{
			{"name": name},
			{"acronym": acronym},
		}},
	}}
	err = coll.FindOne(ctx, query).Decode(protocol)
	return err == nil
}

func GetProtocols(search string, limit, page int) ([]Protocol, int64, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("protocol")
	defer client.Disconnect(ctx)
	// creando parametros de consulta
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(page - 1))
	query := bson.M{}
	if search != "" {
		query = bson.M{"$or": []bson.M{
			{"acronym": primitive.Regex{
				Pattern: `(\s` + search + `|^` + search + `|\w` + search + `\w` + `|` + search + `$` + `|` + search + `\s)`, Options: "i",
			}},
			{"name": primitive.Regex{
				Pattern: `(\s` + search + `|^` + search + `|\w` + search + `\w` + `|` + search + `$` + `|` + search + `\s)`, Options: "i",
			}},
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
	var protocols []Protocol
	if err = cursor.All(ctx, &protocols); err != nil {
		return nil, 0, err
	}
	return protocols, count, nil
}

func UpdateProtocol(protocol *Protocol) error {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("protocol")
	defer client.Disconnect(ctx)
	// actualizando protocolo
	protocol.UpdatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	_, err := coll.UpdateOne(ctx, bson.M{"_id": protocol.ID}, bson.M{"$set": protocol})
	return err
}
