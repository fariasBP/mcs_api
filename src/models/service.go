package models

import (
	"errors"
	"fmt"
	"mcs_api/src/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type StatusService int

type StatusProtocol int

const (
	Operative   StatusProtocol = 0 // Operativo o en funcionamiento
	Limited     StatusProtocol = 1 // Operacion limitada o funcionamiento parcial
	NoOperative StatusProtocol = 2 // No Operativo o no en funcionamiento
	Unknown     StatusProtocol = 3 // Desconocido o que no se sabe en que estado esta
)

const (
	Recived    StatusService = 0 // Recibido y pendiente
	Inspection StatusService = 1 // Inspeccion y diagnostico
	Execute    StatusService = 2 // Ejecutado o ejecutando
	Test       StatusService = 3 // Pruebas, Evaluacion y verificacion
	Delivery   StatusService = 4 // Registro y Entrega
	Cancelled  StatusService = 5 // Cancelado
)

type Status2Service int

const (
	Nono      Status2Service = 0 // cualquier status
	Pending   Status2Service = 1 // pendiente (status: 0, 1)
	Ejecution Status2Service = 2 // en ejecucion (status: 2, 3)
	Finish    Status2Service = 3 // finalizado (status: 4, 5)
)

type (
	Service struct {
		ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		MachineId string             `json:"machine_id" bson:"machine_id,omitempty"`
		StartedAt primitive.DateTime `json:"started_at" bson:"started_at,omitempty"`
		EndedAt   primitive.DateTime `json:"ended_at" bson:"ended_at,omitempty"`
		Keepers   []string           `json:"keepers" bson:"keepers,omitempty"`
		Status    StatusService      `json:"status" bson:"status,omitempty"`
	}
	ServiceRebuild struct {
		ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		StartedAt primitive.DateTime `json:"started_at" bson:"started_at,omitempty"`
		EndedAt   primitive.DateTime `json:"ended_at" bson:"ended_at,omitempty"`
		Keepers   []string           `json:"keepers" bson:"keepers,omitempty"`
		Status    StatusService      `json:"status" bson:"status,omitempty"`
		Machine   MachineRebuild     `json:"machine" bson:"machine,omitempty"`
	}
)

func NewService(machineId, keeper string) error {
	// creando estructura
	newService := &Service{
		MachineId: machineId,
		StartedAt: primitive.NewDateTimeFromTime(time.Now().UTC()),
		Status:    Recived,
		Keepers:   []string{keeper},
	}
	// conectar a la bbdd
	ctx, client, coll := config.ConnectColl("services")
	defer client.Disconnect(ctx)
	// insertando
	_, err := coll.InsertOne(ctx, newService)

	return err
}

func ExistsServiceById(id string) bool {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("services")
	defer client.Disconnect(ctx)
	// obteniendo id
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// buscando servicio
	service := &Service{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(service)

	return err == nil
}

func IsActiveServiceById(id string) bool {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("services")
	defer client.Disconnect(ctx)
	// obteniendo id
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	// buscando servicio
	service := &Service{}
	query := bson.M{"$and": []bson.M{
		{"_id": idObj},
		{"$or": []bson.M{
			{"status": bson.M{"$gte": 0, "$lt": 4}},
			{"status": nil},
		}},
	}}
	err = coll.FindOne(ctx, query).Decode(service)

	return err == nil
}

func IsActiveService(idMachine string) bool {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("services")
	defer client.Disconnect(ctx)
	// buscando servicio
	service := &Service{}
	query := bson.M{"$and": []bson.M{
		{"machine_id": idMachine},
		{"$or": []bson.M{
			{"status": bson.M{"$gte": 0, "$lt": 4}},
			{"status": nil},
		}},
	}}
	err := coll.FindOne(ctx, query).Decode(service)

	return err == nil
}

func SleepService(id string) error {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("services")
	defer client.Disconnect(ctx)
	// obteniendo id
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	// acualizando status y fecha
	update := bson.M{"$set": bson.M{
		"status": Recived,
	}}
	_, err = coll.UpdateOne(ctx, bson.M{"_id": idObj}, update)

	return err
}

func UpdateProgressService(idService string, status int) error {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("services")
	defer client.Disconnect(ctx)
	// obteniendo id
	id, err := primitive.ObjectIDFromHex(idService)
	if err != nil {
		return err
	}
	// verifica el status como en progreso
	if status < int(Inspection) && status > int(Test) {
		return errors.New("status param not valid")
	}
	// acualizando status y fecha
	update := bson.M{"$set": bson.M{
		"status": status,
	}}
	_, err = coll.UpdateOne(ctx, bson.M{"_id": id}, update)

	return err
}

func FinishService(idService string, cancelled bool) error {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("services")
	defer client.Disconnect(ctx)
	// obteniendo id
	id, err := primitive.ObjectIDFromHex(idService)
	if err != nil {
		return err
	}
	// estableciendo si es para cancelar o entregar
	status := Delivery
	if cancelled {
		status = Cancelled
	}
	// acualizando status y fecha
	update := bson.M{"$set": bson.M{
		"ended_at": primitive.NewDateTimeFromTime(time.Now().UTC()),
		"status":   status,
	}}
	_, err = coll.UpdateOne(ctx, bson.M{"_id": id}, update)

	return err
}

func IsFinishedService(id string) bool {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("services")
	defer client.Disconnect(ctx)
	// obteniendo id
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	// buscando servicio
	service := &Service{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(service)
	// verificando si se ha entregado o cancelado
	if err == nil && (service.Status == Delivery || service.Status == Cancelled) {
		return true
	}

	return false
}

func GetServiceById(id string) (*Service, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("services")
	defer client.Disconnect(ctx)
	// obteniendo id
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	// buscando servicio
	service := &Service{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(service)

	return service, err
}

func GetServices(startedAt, endedAt time.Time, ascending bool, status Status2Service, limit, page int) ([]Service, int64, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("services")
	defer client.Disconnect(ctx)
	// creando parametros de consulta
	asc := -1
	if ascending {
		asc = 1
	}
	opts := options.Find().SetSort(bson.M{"started_at": asc}).SetLimit(int64(limit)).SetSkip(int64(page - 1))
	// creando consulta
	query := bson.M{}
	if !startedAt.IsZero() && !endedAt.IsZero() {
		query = bson.M{"started_at": bson.M{"$gte": primitive.NewDateTimeFromTime(startedAt.UTC()), "$lte": primitive.NewDateTimeFromTime(endedAt.UTC())}}
		if status == 1 {
			fmt.Println("quiiii")
			query = bson.M{"$and": []bson.M{
				{"started_at": bson.M{"$gte": primitive.NewDateTimeFromTime(startedAt.UTC()), "$lte": primitive.NewDateTimeFromTime(endedAt.UTC())}},
				{"$or": []bson.M{
					{"status": bson.M{"$eq": 0}},
					{"status": bson.M{"$eq": 1}},
					{"status": bson.M{"$exists": false}},
				}},
			}}
		} else if status == 2 {
			query = bson.M{"$and": []bson.M{
				{"started_at": bson.M{"$gte": primitive.NewDateTimeFromTime(startedAt.UTC()), "$lte": primitive.NewDateTimeFromTime(endedAt.UTC())}},
				{"status": bson.M{"$gte": 2, "$lte": 3}},
			}}
		} else if status == 3 {
			query = bson.M{"$and": []bson.M{
				{"started_at": bson.M{"$gte": primitive.NewDateTimeFromTime(startedAt.UTC()), "$lte": primitive.NewDateTimeFromTime(endedAt.UTC())}},
				{"status": bson.M{"$gte": 4, "$lte": 5}},
			}}
		}
	} else {
		if status == Pending {
			query = bson.M{"$or": []bson.M{
				{"status": bson.M{"$eq": 0}},
				{"status": bson.M{"$eq": 1}},
				{"status": bson.M{"$exists": false}},
			}}
		} else if status == Ejecution {
			query = bson.M{"status": bson.M{"$gte": 2, "$lte": 3}}
		} else if status == Finish {
			query = bson.M{"status": bson.M{"$gte": 4, "$lte": 5}}
		}
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
