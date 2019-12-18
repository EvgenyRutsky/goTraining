package infrastructure

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"httpserver/domain"
	"httpserver/infrastructure/client"
	"time"
)

type ReaderRepository interface {
	InsertReader(reader *domain.Reader) (int, error)
	UpdateReader(reader *domain.Reader) (int, error)
	DeleteReader(id int) (int, error)
	GetReaderByID(id int) (*domain.Reader, error)
	GetReaders() ([]*domain.Reader, error)
}

type readersRepository struct {
	context context.Context
	config *client.Config
}

func NewReaderRepository(config *client.Config) ReaderRepository {
	return &readersRepository {
		context: context.Background(),
		config: config,
	}
}

func (r *readersRepository) InsertReader(reader *domain.Reader) (int, error) {
	ctx, cancel := context.WithTimeout(r.context, 5*time.Second)
	defer cancel()
	opencl, err := client.NewClient(r.config).Open(ctx)
	if err != nil {
		return 0, err
	}
	collection := opencl.Database(r.config.Dbname).Collection(r.config.ReadersCollection)
	_, err = collection.InsertOne(ctx, reader)
	if err != nil {
		return 0, err
	}

	if err := client.Close(ctx, opencl); err != nil {
		return 0, err
	}

	return reader.ID, nil
}

func (r *readersRepository) UpdateReader(reader *domain.Reader) (int, error) {
	ctx, cancel := context.WithTimeout(r.context, 5*time.Second)
	defer cancel()
	opencl, err := client.NewClient(r.config).Open(ctx)
	if err != nil {
		return 0, err
	}

	collection := opencl.Database(r.config.Dbname).Collection(r.config.ReadersCollection)
	filter := bson.D{{"id",reader.ID}}
	update := bson.D{
		{"$set",bson.D{
			{"name", reader.Name},
			{"book_id", reader.BookID},
		}},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	if err := client.Close(ctx, opencl); err != nil {
		return 0, err
	}
	return int(result.ModifiedCount), nil
}

func (r *readersRepository) DeleteReader(id int) (int, error) {
	ctx, cancel := context.WithTimeout(r.context, 5*time.Second)
	defer cancel()
	opencl, err := client.NewClient(r.config).Open(ctx)
	if err != nil {
		return 0, err
	}

	collection := opencl.Database(r.config.Dbname).Collection(r.config.ReadersCollection)
	filter := bson.D{{"id",id}}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}

	if err := client.Close(ctx, opencl); err != nil {
		return 0, err
	}
	return int(result.DeletedCount), nil
}

func (r *readersRepository) GetReaderByID(id int) (*domain.Reader, error) {
	ctx, cancel := context.WithTimeout(r.context, 5*time.Second)
	defer cancel()
	opencl, err := client.NewClient(r.config).Open(ctx)
	if err != nil {
		return nil, err
	}
	var reader domain.Reader
	collection := opencl.Database(r.config.Dbname).Collection(r.config.ReadersCollection)
	filter := bson.D{{"id",id}}
	err = collection.FindOne(ctx, filter).Decode(&reader)
	if err != nil {
		return  nil, err
	}

	if err := client.Close(ctx, opencl); err != nil {
		return nil, err
	}
	return &reader, nil
}

func (r *readersRepository) GetReaders() ([]*domain.Reader, error) {
	ctx, cancel := context.WithTimeout(r.context, 5*time.Second)
	defer cancel()
	opencl, err := client.NewClient(r.config).Open(ctx)
	if err != nil {
		return nil, err
	}
	collection := opencl.Database(r.config.Dbname).Collection(r.config.ReadersCollection)
	cur, err := collection.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}
	var results []*domain.Reader
	for cur.Next(ctx) {
		var elem *domain.Reader
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	if err := client.Close(ctx, opencl); err != nil {
		return nil, err
	}
	return results, nil
}


