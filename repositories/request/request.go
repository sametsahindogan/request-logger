package request

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	data.Id = primitive.NewObjectID()
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

	model := requestModel.NewRequest()

	total, err := model.CountDocuments(context.TODO(), query)

	if err != nil {
		return nil, nil, err
	}

	page := request.Page
	perPage := request.PerPage

	queryOptions := options.Find()
	queryOptions.SetSkip((page - 1) * perPage)
	queryOptions.SetLimit(perPage)
	sort := 1
	if request.Sort == "DESC" {
		sort = -1
	}

	queryOptions.SetSort(bson.M{"created_at": sort})

	cursor, errs := model.Find(context.TODO(), query, queryOptions)

	if errs != nil {
		return nil, nil, errs
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

	totalPage := total / perPage

	if (total % perPage) != 0 {
		totalPage++
	}

	queryInformation := map[string]interface{}{
		"total_rows":    total,
		"current_rows":  int64(len(results)),
		"page":          page,
		"per_page":      perPage,
		"total_page":    totalPage,
		"has_next_page": total > (page * perPage),
	}

	return results, queryInformation, nil
}

func requestHasDate(time time.Time) bool {
	return !reflect.DeepEqual(time, reflect.Zero(reflect.TypeOf(time)).Interface())
}
