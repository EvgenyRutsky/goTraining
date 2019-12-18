package infrastructure

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"httpserver/domain"
	"httpserver/infrastructure/client"
	"time"
)

type BooksRepository interface {
	InsertBook(book *domain.Book) (int, error)
	UpdateBook(book *domain.Book) (int, error)
	DeleteBook(id int) (int, error)
	GetBookByID(id int) (*domain.Book, error)
	GetBooks() ([]*domain.Book, error)
}

type booksRepository struct {
	context context.Context
	config *client.Config
}

func NewBookRepository(config *client.Config) BooksRepository {
	return &booksRepository{
		context: context.Background(),
		config: config,
	}
}

func (br *booksRepository) InsertBook(book *domain.Book) (int, error) {
	ctx, cancel := context.WithTimeout(br.context, 5*time.Second)
	defer cancel()
	opencl, err := client.NewClient(br.config).Open(ctx)
	if err != nil {
		return 0, err
	}
	collection := opencl.Database(br.config.Dbname).Collection(br.config.BooksCollection)
	_, err = collection.InsertOne(ctx, book)
	if err != nil {
		return 0, err
	}

	if err := client.Close(ctx, opencl); err != nil {
		return 0, err
	}

	return book.ID, nil
}

func (br *booksRepository) UpdateBook(book *domain.Book) (int, error) {
	ctx, cancel := context.WithTimeout(br.context, 5*time.Second)
	defer cancel()
	opencl, err := client.NewClient(br.config).Open(ctx)
	if err != nil {
		return 0, err
	}

	collection := opencl.Database(br.config.Dbname).Collection(br.config.BooksCollection)
	filter := bson.D{{"id",book.ID}}
	update := bson.D{
		{"$set",bson.D{
			{"name", book.Name},
			{"author", book.Author},
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

func (br *booksRepository) DeleteBook(id int) (int, error) {
	ctx, cancel := context.WithTimeout(br.context, 5*time.Second)
	defer cancel()
	opencl, err := client.NewClient(br.config).Open(ctx)
	if err != nil {
		return 0, err
	}

	collection := opencl.Database(br.config.Dbname).Collection(br.config.BooksCollection)
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

func (br *booksRepository) GetBookByID(id int) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(br.context, 5*time.Second)
	defer cancel()
	opencl, err := client.NewClient(br.config).Open(ctx)
	if err != nil {
		return nil, err
	}
	var book domain.Book
	collection := opencl.Database(br.config.Dbname).Collection(br.config.BooksCollection)
	filter := bson.D{{"id",id}}
	err = collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		return  nil, err
	}

	if err := client.Close(ctx, opencl); err != nil {
		return nil, err
	}
	return &book, nil
}

func (br *booksRepository) GetBooks() ([]*domain.Book, error) {
	ctx, cancel := context.WithTimeout(br.context, 5*time.Second)
	defer cancel()
	opencl, err := client.NewClient(br.config).Open(ctx)
	if err != nil {
		return nil, err
	}
	collection := opencl.Database(br.config.Dbname).Collection(br.config.BooksCollection)
	cur, err := collection.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}
	var results []*domain.Book
	for cur.Next(ctx) {
		var elem *domain.Book
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



