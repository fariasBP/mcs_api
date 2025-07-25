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
	}
	MachineRebuild struct {
		ID              primitive.ObjectID `json:"id"`
		CreatedAt       primitive.DateTime `json:"created_at"`
		UpdatedAt       primitive.DateTime `json:"updated_at"`
		CompanyId       string             `json:"company_id"`
		CompanyName     string             `json:"company_name"`
		CompanyManager  string             `json:"company_manager"`
		CompanyContact  string             `json:"company_contact"`
		MachineTypeId   string             `json:"machine_type_id"`
		MachineTypeName string             `json:"machine_type_name"`
		BrandId         string             `json:"brand_id"`
		BrandName       string             `json:"brand_name"`
		Serial          string             `json:"serial"`
		Model           string             `json:"model"`
	}
)

type ParamGetMachines string

const (
	CompanyParam     ParamGetMachines = "company_id"
	MachineTypeParam ParamGetMachines = "machine_type_id"
	BrandParam       ParamGetMachines = "brand_id"
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
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	// obteniendo id
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return false
	}
	// consultando
	machine := &Machine{}
	err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(machine)
	return err == nil
}

func ExistsMachine(serial string) bool {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	// consultando
	machine := &Machine{}
	err := coll.FindOne(ctx, bson.M{"serial": serial}).Decode(machine)
	return err == nil
}

// existe otra maquina con el mismo serial sin contar este
func ExistsOtherMachine(machineId, serial string) bool {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	// obteniendo id
	id, err := primitive.ObjectIDFromHex(machineId)
	if err != nil {
		return false
	}
	// consultando
	machine := &Machine{}
	query := bson.M{"$and": []bson.M{
		{"_id": bson.M{"$ne": id}},
		{"serial": serial},
	}}
	err = coll.FindOne(ctx, query).Decode(machine)
	return err == nil
}

func GetMachineById(id string) (*Machine, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	// obteniendo id
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	// consultando
	machine := &Machine{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(machine)
	return machine, err
}

func GetMachinesByModelAndSerial(search string, limit, page int) ([]Machine, int64, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	// creando parametros de consulta
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(page - 1))
	query := bson.M{}
	if search != "" {
		query = bson.M{"$or": []bson.M{
			{"serial": primitive.Regex{
				Pattern: `(\s` + search + `|^` + search + `|\w` + search + `\w` + `|` + search + `$` + `|` + search + `\s)`, Options: "i",
			}},
			{"model": primitive.Regex{
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
	var machines []Machine
	if err = cursor.All(ctx, &machines); err != nil {
		return nil, 0, err
	}
	return machines, count, nil
}

func GetMachinesByCompanyOrBrandOrMachineTypeId(paramId string, param ParamGetMachines, limit, page int) ([]Machine, int64, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	// creando parametros de consulta
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(page - 1))
	query := bson.M{string(param): paramId}
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

func UpdateMachine(machine *Machine) error {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("machine")
	defer client.Disconnect(ctx)
	// actualizando maquina
	machine.UpdatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	_, err := coll.UpdateOne(ctx, bson.M{"_id": machine.ID}, bson.M{"$set": machine})
	return err
}
