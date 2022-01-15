package recordRepository

import (
	"context"
	"log"
	"net/http"

	recordType "getir-case/internal/record/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Aggregate(req *recordType.GetRecordsDBModel) *recordType.GetRecordsResponse
}

type repository struct {
	mc *mongo.Collection
}

func NewRepository(mc *mongo.Collection) *repository {
	return &repository{mc: mc}
}

func (r *repository) Aggregate(model *recordType.GetRecordsDBModel) *recordType.GetRecordsResponse {

	ctx, cancelCtx := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancelCtx()

	res := new(recordType.GetRecordsResponse)
	res.Records = make([]recordType.Record, 0)

	pipeline := []bson.D{}

	//filtering for cretaedAt >= model.startDate && createdAt <= model.endDate
	pipeline = append(pipeline, bson.D{{
		"$match", bson.D{
			{"createdAt", bson.D{
				{"$gte", model.StartDate},
				{"$lte", model.EndDate},
			}}},
	}})

	//projection for record fields in response model
	pipeline = append(pipeline, bson.D{{
		"$project", bson.D{
			{"_id", 0},
			{"key", "$key"},
			{"createdAt", "$createdAt"},
			{"totalCount", bson.D{{"$sum", "$counts"}}},
		},
	}})

	//filtering for projected records

	pipeline = append(pipeline, bson.D{{
		"$match", bson.D{
			{"totalCount", bson.D{
				{"$gt", model.MinCount},
				{"$lt", model.MaxCount},
			}}}}})

	cur, err := r.mc.Aggregate(ctx, pipeline, nil)
	if err != nil {
		go log.Printf("Aggregate error: %s", err.Error())
		res.Msg = "MongoDB aggregate failed"
		res.Code = http.StatusInternalServerError
		return res
	}

	if err := cur.All(ctx, &res.Records); err != nil {
		res.Msg = "MongoDB cursor failed"
		res.Code = http.StatusInternalServerError
		go log.Printf("Mongo cursor error: %s", err.Error())
		return res
	}

	res.Msg = "Success"
	return res
}
