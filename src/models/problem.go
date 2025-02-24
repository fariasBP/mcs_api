package models

import (
	"mcs_api/src/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Problem struct {
	Id         string `json:"id" bson:"_id,omitempty"`
	ServiceId  string `json:"service_id" bson:"service_id,omitempty"`
	ProtocolId string `json:"protocol_id" bson:"protocol_id,omitempty"`
	Problem    string `json:"problem" bson:"problem,omitempty"`
	Solution   string `json:"solution" bson:"solution,omitempty"`
}

func NewProblem(idService, idProtocol, problem string) error {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("problems")
	defer client.Disconnect(ctx)
	// insertando
	newProblem := &Problem{
		ServiceId:  idService,
		ProtocolId: idProtocol,
		Problem:    problem,
	}
	_, err := coll.InsertOne(ctx, newProblem)

	return err
}

func ExistsProblemById(id string) bool {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("problems")
	defer client.Disconnect(ctx)
	// obteniendo id
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	// buscando problem
	problem := &Problem{}
	err = coll.FindOne(ctx, bson.M{"_id": idObj}).Decode(problem)

	return err == nil
}

func NewSolution(idProblem, solution string) error {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("problems")
	defer client.Disconnect(ctx)
	// obteniendo id
	id, err := primitive.ObjectIDFromHex(idProblem)
	if err != nil {
		return err
	}
	// insertando
	update := bson.M{"$set": bson.M{"solution": solution}}
	_, err = coll.UpdateOne(ctx, bson.M{"_id": id}, update)

	return err
}

func UpdateProblem(idProblem, problem string) error {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("problems")
	defer client.Disconnect(ctx)
	// obteniendo id
	id, err := primitive.ObjectIDFromHex(idProblem)
	if err != nil {
		return err
	}
	// insertando
	update := bson.M{"$set": bson.M{"problem": problem}}
	_, err = coll.UpdateOne(ctx, bson.M{"_id": id}, update)

	return err
}

func UpdateSolution(idProblem, solution string) error {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("problems")
	defer client.Disconnect(ctx)
	// obteniendo id
	id, err := primitive.ObjectIDFromHex(idProblem)
	if err != nil {
		return err
	}
	// insertando
	update := bson.M{"$set": bson.M{"solution": solution}}
	_, err = coll.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func GetProblemsByServiceId(idService string) ([]Problem, error) {
	// conectando a la base de datos
	ctx, client, coll := config.ConnectColl("problems")
	defer client.Disconnect(ctx)
	// creando parametros de consulta
	query := bson.M{"service_id": idService}
	// consultando
	cursor, err := coll.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	// modelando datos
	var problems []Problem
	if err = cursor.All(ctx, &problems); err != nil {
		return nil, err
	}

	return problems, err
}
