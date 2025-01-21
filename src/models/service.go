package models

import (
	"fmt"
	"mcs_api/src/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	Name   string  `json:"name" bson:"name,omitempty"`
	Price  float64 `json:"price" bson:"price,omitempty"`
	Number int     `json:"number" bson:"number,omitempty"`
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

func GetServices(machine_id, endedAt string, limit, page int) ([]Service, int64, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("service")
	defer client.Disconnect(ctx)

	// creando parametros de consulta
	opts := options.Find().SetSort(bson.M{"ended_at": -1}).SetLimit(int64(limit)).SetSkip(int64(page - 1))
	query := bson.M{"machine_id": machine_id}
	if endedAt != "" {
		// trabajando en fechas
		ended, err := time.Parse(time.DateTime, endedAt)
		if err != nil {
			return nil, 0, err
		}
		// creando query con ended
		query = bson.M{"$and": []bson.M{
			{"machine_id": machine_id},
			{"ended_at": bson.M{"$lte": primitive.NewDateTimeFromTime(ended.UTC())}},
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
	var services []Service
	if err = cursor.All(ctx, &services); err != nil {
		return nil, 0, err
	}
	return services, count, nil
}
