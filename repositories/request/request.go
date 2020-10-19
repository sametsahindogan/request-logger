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

	query := bson.M{"domain": request.Domain}

	if requestHasDate(request.Created) {
		query["created_at"] = bson.M{
			"$lt":  request.Created.AddDate(0, 0, 1),
			"$gte": request.Created,
		}
	}

	if request.UserId != "" {
		query["user_id"] = request.UserId
	}

	if request.Uri != "" {
		query["uri"] = request.Uri
	}

	if request.Ip != "" {
		query["ip_address"] = request.Ip
	}

	model := requestModel.NewRequest()

	total, err := model.CountDocuments(context.TODO(), query)

	if err != nil {
		return nil, nil, err
	}

	queryOptions, page, perPage := prepareQueryOptions(request)

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

	queryInformation := map[string]interface{}{
		"totalRows":   total,
		"currentRows": int64(len(results)),
		"page":        page,
		"perPage":     perPage,
		"totalPage":   getTotalPage(total, perPage),
		"hasNextPage": total > (page * perPage),
	}

	return results, queryInformation, nil
}

func requestHasDate(time time.Time) bool {
	return !reflect.DeepEqual(time, reflect.Zero(reflect.TypeOf(time)).Interface())
}

func getTotalPage(total int64, perPage int64) int64 {

	totalPage := total / perPage

	if (total % perPage) != 0 {
		totalPage++
	}

	return totalPage
}

func prepareQueryOptions(request *requestTypes.GetByDomainRequestValidation) (*options.FindOptions, int64, int64) {
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

	return queryOptions, page, perPage
}
