package models

import (
	"fmt"
	"mcs_api/src/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Problem struct {
	Problem  string `json:"problem" bson:"problem,omitempty"`
	Solution string `json:"solution" bson:"solution,omitempty"`
}

type ProtocolData struct {
	Protocol string                `json:"protocol" bson:"protocol,omitempty"`
	Status   config.StatusProtocol `json:"status" bson:"status,omitempty"`
	Note     string                `json:"note" bson:"note,omitempty"`
	Problems []Problem             `json:"problems" bson:"problems,omitempty"`
}
type Material struct {
	Name   string `json:"name" bson:"name,omitempty"`
	Price  int    `json:"price" bson:"price,omitempty"`
	Number int    `json:"number" bson:"number,omitempty"`
}
type Service struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	MachineId string               `json:"machine_id" bson:"machine_id,omitempty"`
	StartedAt primitive.DateTime   `json:"started_at" bson:"started_at"`
	EndedAt   primitive.DateTime   `json:"ended_at" bson:"ended_at"`
	Protocols []ProtocolData       `json:"protocols" bson:"protocols,omitempty"`
	Materials []Material           `json:"materials" bson:"materials,omitempty"`
	Comments  string               `json:"comments" bson:"comments,omitempty"`
	Status    config.StatusService `json:"status" bson:"status,omitempty"`
}

func CreateService(machineId, comments string, startedAt, endedAt time.Time, protocols []ProtocolData, materials []Material, status config.StatusService) error {
	newService := &Service{
		MachineId: machineId,
		StartedAt: primitive.NewDateTimeFromTime(startedAt.UTC()),
		EndedAt:   primitive.NewDateTimeFromTime(endedAt.UTC()),
		Comments:  comments,
		Protocols: protocols,
		Materials: materials,
		Status:    status,
	}
	ctx, client, coll := config.ConnectColl("service")
	defer client.Disconnect(ctx)
	_, err := coll.InsertOne(ctx, newService)
	return err
}

func ExistsServiceById(id string) bool {
	ctx, client, coll := config.ConnectColl("service")
	defer client.Disconnect(ctx)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	service := &Service{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(service)
	return err == nil
}

func GetServiceById(id string) (*Service, error) {
	ctx, client, coll := config.ConnectColl("service")
	defer client.Disconnect(ctx)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	service := &Service{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(service)
	return service, err
}
