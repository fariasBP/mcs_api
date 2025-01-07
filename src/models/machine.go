package models

import (
	"mcs_api/src/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Machine struct {
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
