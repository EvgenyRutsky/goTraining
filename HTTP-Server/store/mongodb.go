package store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/model"
)

const (
	collection = "books"
	dbname = "dev"
)

type MongoRepository struct {
	config *Config
	context context.Context
	client *mongo.Client
}

func NewMongoRepository(config *Config) Repository {
	return &MongoRepository{
		config:   config,
		context: context.Background(),
	}
}

func (m *MongoRepository) open() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.config.DatabaseURL))
	if  err != nil {
		return err
	}

	if err := client.Connect(m.context); err != nil {
		return  err
	}

	if err := client.Ping(m.context, nil); err != nil {
		return err
	}

	m.client = client
	return nil
}

func (m *MongoRepository) close() error {
	if err := m.client.Disconnect(m.context); err != nil {
		return err
	}
	return nil
}

func (m *MongoRepository) Create(b *model.Book) (interface{}, error) {
	if err := m.open(); err != nil {
		return nil, err
	}

	collection := m.client.Database(dbname).Collection(collection)

	insertResult, err := collection.InsertOne(m.context, b)
	if err != nil {
		return nil, err
	}

	if err := m.close(); err != nil {
		return nil, err
	}

	return insertResult.InsertedID, nil
}

func (m *MongoRepository) FindByID(id int) (*model.Book, error) {

	if err := m.open(); err != nil {
		return nil, err
	}

	filter := bson.D{{"id",id}}
	var result *model.Book

	collection := m.client.Database(dbname).Collection(collection)
	err := collection.FindOne(m.context, filter).Decode(result)
	if err != nil {
		return  nil, err
	}

	if err := m.close(); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MongoRepository) GetAll() ([]*model.Book, error) {
	if err := m.open(); err != nil {
		return nil, err
	}
	var results []*model.Book
	collection := m.client.Database(dbname).Collection(collection)
	filter := bson.M{}
	option := options.Find()

	cur, err := collection.Find(m.context, filter, option)
	if err != nil {
		return nil, err
	}

	for cur.Next(m.context) {
		var elem *model.Book
		err := cur.Decode(elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(m.context)

	if err := m.close(); err != nil {
		return nil, err
	}
	return results, nil
}

func (m *MongoRepository) Delete(id int) (int, error) {
	if err := m.open(); err != nil {
		return 0, err
	}

	collection := m.client.Database(dbname).Collection(collection)
	filter := bson.D{{"id",id}}

	result, err := collection.DeleteOne(m.context, filter)
	if err != nil {
		return 0, err
	}

	if err := m.close(); err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}

func (m *MongoRepository) Update(b *model.Book) (int, error) {
	if err := m.open(); err != nil {
		return 0, err
	}

	collection := m.client.Database(dbname).Collection(collection)
	filter := bson.D{{"id",b.ID}}
	update := bson.D{
		{"$set",bson.D{
				{"name", b.Name},
				{"author", b.Author},
			}},
	}

	result, err := collection.UpdateOne(m.context, filter, update)
	if err != nil {
		return 0, err
	}

	if err := m.close(); err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}


