package request

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"request-logger/database"
	"time"
)

var CollectionName = "requests"

type Request struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    string             `bson:"user_id"`
	IpAddress string             `bson:"ip_address"`
	Uri       string             `bson:"uri"`
	Domain    string             `bson:"domain"`
	CreatedAt time.Time          `bson:"created_at"`
}

/*func (r Request) Formatted(tag string, request *Request) map[string]interface{} {

	return database.ModelStructFieldTagResolver(tag, reflect.ValueOf(request), reflect.TypeOf(*request))
}*/

func NewRequest() *mongo.Collection {

	client := database.GetConnection()

	return client.Collection(CollectionName)
}
