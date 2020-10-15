package request

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
	requestModel "request-logger/models/request"
	requestTypes "request-logger/requests/request"
	"time"
)

type RequestRepository struct{}

func (r RequestRepository) Create(request *requestTypes.StoreProcessRequestValidation) error {

	data := requestModel.Request{}
	data.UserId = request.UserId
	data.IpAddress = request.IpAddress
	data.Uri = request.Uri
	data.Domain = request.Domain
	data.CreatedAt = time.Now()

	_, err := requestModel.NewRequest().InsertOne(context.TODO(), data)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (r RequestRepository) GetByDomainName(request *requestTypes.GetByDomainRequestValidation) ([]requestModel.Request, map[string]interface{}, error) {

	data := requestModel.Request{}
	data.Domain = request.Domain

	query := bson.M{"domain": data.Domain}

	if requestHasDate(request.Created) {
		data.CreatedAt = request.Created

		query["created_at"] = bson.M{
			"$lt":  data.CreatedAt.AddDate(0, 0, 1),
			"$gte": data.CreatedAt,
		}
	}

	if request.UserId != "" {
		data.UserId = request.UserId

		query["user_id"] = data.UserId
	}

	total, err := requestModel.NewRequest().CountDocuments(context.TODO(), query)

	if err != nil {
		return nil, nil, err
	}

	queryOptions := options.Find()
	queryOptions.SetSkip(request.Offset)
	queryOptions.SetLimit(request.Limit)
	sort := 1
	if request.Sort == "DESC" {
		sort = -1
	}

	queryOptions.SetSort(bson.M{"created_at": sort})

	cursor, errs := requestModel.NewRequest().Find(context.TODO(), query, queryOptions)

	if errs != nil {
		return nil, nil, err
	}

	var results []requestModel.Request

	var mapping requestModel.Request

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&mapping)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, mapping)
	}

	queryInformation := map[string]interface{}{
		"total":        total,
		"offset":       request.Offset,
		"limit":        request.Limit,
		"current_rows": len(results),
	}

	return results, queryInformation, nil
}

func requestHasDate(time time.Time) bool {
	return !reflect.DeepEqual(time, reflect.Zero(reflect.TypeOf(time)).Interface())
}
