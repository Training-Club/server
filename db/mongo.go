package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"tc-server/config"
	"time"
)

// GetMongoContext returns a pre-configured context
// used explicitly for Mongo processes.
func GetMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// InitMongo initializes a new connection to the Mongo server.
func InitMongo(conf *config.MongoConfig) (*mongo.Client, error) {
	return mongo.Connect(context.Background(), options.Client().ApplyURI(conf.URI+conf.DatabaseName+"?authSource=admin"))
}

func FindDocumentById[K any](
	params MongoParams,
	id string,
) (K, error) {
	ctx, cancel := GetMongoContext()
	collection := params.Client.Database(params.DBName).Collection(params.CollectionName)
	defer cancel()

	var document K
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return document, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&document)
	return document, err
}

func FindDocumentByKeyValue[K any, V any](
	params MongoParams,
	k string,
	v K,
) (V, error) {
	ctx, cancel := GetMongoContext()
	collection := params.Client.Database(params.DBName).Collection(params.CollectionName)
	defer cancel()

	var document V
	err := collection.FindOne(ctx, bson.M{k: v}).Decode(&document)
	return document, err
}

func FindDocumentByFilter[K any](
	params MongoParams,
	filter bson.M,
) (K, error) {
	ctx, cancel := GetMongoContext()
	collection := params.Client.Database(params.DBName).Collection(params.CollectionName)
	defer cancel()

	var document K
	err := collection.FindOne(ctx, filter).Decode(&document)
	return document, err
}

func FindManyDocumentsByKeyValue[K any, V any](
	params MongoParams,
	k string,
	v K,
) ([]V, error) {
	ctx, cancel := GetMongoContext()
	collection := params.Client.Database(params.DBName).Collection(params.CollectionName)
	defer cancel()

	var documents []V
	cursor, err := collection.Find(ctx, bson.M{k: v})
	err = cursor.All(ctx, &documents)
	return documents, err
}

func FindManyDocumentsByFilter[K any](
	params MongoParams,
	filter interface{},
) ([]K, error) {
	ctx, cancel := GetMongoContext()
	collection := params.Client.Database(params.DBName).Collection(params.CollectionName)
	defer cancel()

	var documents []K
	cursor, err := collection.Find(ctx, filter)
	err = cursor.All(ctx, &documents)
	return documents, err
}

func FindManyDocumentsByFilterWithOpts[K any](
	params MongoParams,
	filter interface{},
	opts *options.FindOptions,
) ([]K, error) {
	ctx, cancel := GetMongoContext()
	collection := params.Client.Database(params.DBName).Collection(params.CollectionName)
	defer cancel()

	var documents []K
	cursor, err := collection.Find(ctx, filter, opts)
	err = cursor.All(ctx, &documents)
	return documents, err
}

func InsertDocument[K any](
	params MongoParams,
	document K,
) (string, error) {
	ctx, cancel := GetMongoContext()
	collection := params.Client.Database(params.DBName).Collection(params.CollectionName)
	defer cancel()

	result, err := collection.InsertOne(ctx, document)
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func ReplaceDocument[K any](
	params MongoParams,
	documentId primitive.ObjectID,
	replacement K) (*mongo.UpdateResult, error) {
	ctx, cancel := GetMongoContext()
	collection := params.Client.Database(params.DBName).Collection(params.CollectionName)
	defer cancel()

	result, err := collection.UpdateOne(ctx, bson.M{"_id": documentId}, bson.M{"$set": replacement})
	return result, err
}

func UpdateDocument[K any](
	params MongoParams,
	documentId primitive.ObjectID,
	key string,
	value K) (*mongo.UpdateResult, error) {
	ctx, cancel := GetMongoContext()
	collection := params.Client.Database(params.DBName).Collection(params.CollectionName)
	defer cancel()

	result, err := collection.UpdateOne(ctx, bson.M{"_id": documentId}, bson.D{{"$set", bson.D{{key, value}}}})
	return result, err
}

func UpdateDocumentByFilter[K any](
	params MongoParams,
	documentId primitive.ObjectID,
	updateFilter interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := GetMongoContext()
	collection := params.Client.Database(params.DBName).Collection(params.CollectionName)
	defer cancel()

	result, err := collection.UpdateOne(ctx, bson.M{"_id": documentId}, updateFilter)
	return result, err
}

func DeleteDocument[K any](
	params MongoParams,
	document K,
) (*mongo.DeleteResult, error) {
	ctx, cancel := GetMongoContext()
	collection := params.Client.Database(params.DBName).Collection(params.CollectionName)
	defer cancel()

	result, err := collection.DeleteOne(ctx, document)
	return result, err
}
