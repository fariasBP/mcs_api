package models

import (
	"mcs_api/src/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type (
	Machine struct {
		ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		CreatedAt     primitive.DateTime `json:"created_at" bson:"created_at"`
		UpdatedAt     primitive.DateTime `json:"updated_at" bson:"updated_at"`
		CompanyId     string             `json:"company_id" bson:"company_id,omitempty"`           // id de la empresa
		MachineTypeId string             `json:"machine_type_id" bson:"machine_type_id,omitempty"` // id del tipo de maquina
		BrandId       string             `json:"brand_id" bson:"brand_id,omitempty"`               // id de la marca
		Serial        string             `json:"serial" bson:"serial,omitempty"`
		Model         string             `json:"model" bson:"model,omitempty"`
		Services      []Service          `json:"services" bson:"services,omitempty"`
	}
	MachineRebuild struct {
		ID          primitive.ObjectID `json:"id"`
		CreatedAt   primitive.DateTime `json:"created_at"`
		UpdatedAt   primitive.DateTime `json:"updated_at"`
		Company     string             `json:"company"`
		MachineType string             `json:"machine_type"`
		Brand       string             `json:"brand"`
		Serial      string             `json:"serial"`
		Model       string             `json:"model"`
		Services    []Service          `json:"services"`
	}
	MachineRebuildBasic struct {
		ID          primitive.ObjectID `json:"id"`
		CreatedAt   primitive.DateTime `json:"created_at"`
		UpdatedAt   primitive.DateTime `json:"updated_at"`
		Company     string             `json:"company"`
		MachineType string             `json:"machine_type"`
		Brand       string             `json:"brand"`
		Serial      string             `json:"serial"`
		Model       string             `json:"model"`
	}
)

func CreateMachine(companyId, machineTypeId, brandId, serial, model string) error {
	newMachine := &Machine{
		CreatedAt:     primitive.NewDateTimeFromTime(time.Now().UTC()),
		UpdatedAt:     primitive.NewDateTimeFromTime(time.Now().UTC()),
		CompanyId:     companyId,
		MachineTypeId: machineTypeId,
		BrandId:       brandId,
		Serial:        serial,
		Model:         model,
	}
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	_, err := coll.InsertOne(ctx, newMachine)
	return err
}

func ExistsMachineById(idStr string) bool {
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return false
	}
	machine := &Machine{}
	err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(machine)
	return err == nil
}

func GetMachineById(id string) (*Machine, error) {
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	machine := &Machine{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(machine)
	return machine, err
}

func GetMachinesByCompanyIdAndSerial(companyId, serial string, limit, page int) ([]Machine, int64, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	// creando parametros de consulta
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(page - 1))
	query := bson.M{"company_id": companyId}
	if serial != "" {
		query = bson.M{"$and": []bson.M{
			{"company_id": companyId},
			{"serial": primitive.Regex{
				Pattern: `(\s` + serial + `|^` + serial + `|\w` + serial + `\w` + `|` + serial + `$` + `|` + serial + `\s)`, Options: "i",
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
	var machines []Machine
	if err = cursor.All(ctx, &machines); err != nil {
		return nil, 0, err
	}
	return machines, count, nil
}

func GetMachinesBySerial(serial string, limit, page int) ([]Machine, int64, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	// creando parametros de consulta
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(page - 1))
	query := bson.M{}
	if serial != "" {
		query = bson.M{"serial": primitive.Regex{
			Pattern: `(\s` + serial + `|^` + serial + `|\w` + serial + `\w` + `|` + serial + `$` + `|` + serial + `\s)`, Options: "i",
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
	var machines []Machine
	if err = cursor.All(ctx, &machines); err != nil {
		return nil, 0, err
	}
	return machines, count, nil
}

func GetMachinesBasicBySerial(serial string, limit, page int) ([]Machine, int64, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	// creando parametros de consulta
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(page - 1)).SetProjection(bson.M{"services": 0})
	query := bson.M{}
	if serial != "" {
		query = bson.M{"serial": primitive.Regex{
			Pattern: `(\s` + serial + `|^` + serial + `|\w` + serial + `\w` + `|` + serial + `$` + `|` + serial + `\s)`, Options: "i",
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
	var machines []Machine
	if err = cursor.All(ctx, &machines); err != nil {
		return nil, 0, err
	}
	return machines, count, nil
}
