package storages

import (
	"context"
	"errors"
	"fmt"

	"github.com/Matvey1109/LibraryManagementSystemCore/internal/models"
	"github.com/Matvey1109/LibraryManagementSystemCore/pkg/loadenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ! Abstract Factory
type StorageFactory interface {
	CreateStorage() (Storage, error)
}

func GetStorageFactory() (StorageFactory, error) {
	typeOfStorage, _, _ := loadenv.LoadGlobalEnv()

	if typeOfStorage == "local" {
		return &LocalStorageFactory{}, nil
	}

	if typeOfStorage == "mongodb" {
		return &MongoDBStorageFactory{}, nil
	}

	return nil, errors.New("typeOfStorage not found")
}

// ! Concrete Factories
type LocalStorageFactory struct{} // * Implements interface StorageFactory

func (f *LocalStorageFactory) CreateStorage() (Storage, error) {
	return &LocalStorage{
		members:    []models.Member{},
		books:      []models.Book{},
		borrowings: []models.Borrowing{},
	}, nil
}

type MongoDBStorageFactory struct{} // * Implements interface StorageFactory

func (f *MongoDBStorageFactory) CreateStorage() (Storage, error) {
	MONGO_INITDB_ROOT_USERNAME, MONGO_INITDB_ROOT_PASSWORD := loadenv.LoadMongoEnv()
	uri := fmt.Sprintf("mongodb://%s:%s@mongo:27017/", MONGO_INITDB_ROOT_USERNAME, MONGO_INITDB_ROOT_PASSWORD)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	database := client.Database("mydatabase")
	membersCollection := database.Collection("members")
	booksCollection := database.Collection("books")
	borrowingsCollection := database.Collection("borrowings")

	return &MongoDBStorage{
		client:               client,
		membersCollection:    membersCollection,
		booksCollection:      booksCollection,
		borrowingsCollection: borrowingsCollection,
	}, nil
}

// ! Abstract Product
type Storage interface {
	GetAllMembersStorage() ([]models.Member, error)
	GetMemberStorage(id string) (models.Member, error)
	AddMemberStorage(member models.Member) error
	UpdateMemberStorage(id string, member models.Member) error
	DeleteMemberStorage(id string) error

	GetAllBooksStorage() ([]models.Book, error)
	GetBookStorage(id string) (models.Book, error)
	AddBookStorage(book models.Book) error
	UpdateBookStorage(id string, book models.Book) error
	DeleteBookStorage(id string) error

	GetAllBorrowingsStorage() ([]models.Borrowing, error)
	GetBorrowingStorage(id string) (models.Borrowing, error)
	AddBorrowingStorage(borrowing models.Borrowing) error
	UpdateBorrowingStorage(id string, borrowing models.Borrowing) error
	DeleteBorrowingStorage(id string) error
}

var (
	ExportStorageFactory, _ = GetStorageFactory()
	ExportStorage, _        = ExportStorageFactory.CreateStorage()
)
