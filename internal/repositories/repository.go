package repositories

import (
	"github.com/Matvey1109/LibraryManagementSystemCore/internal/storages"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var storage = storages.ExportStorage

type Repository struct {
	MemberRepository
	BookRepository
	BorrowingRepository
}

func NewRepository() *Repository {
	return &Repository{}
}

func GenerateID() string {
	switch storage.(type) {
	case *storages.LocalStorage:
		objectID := primitive.NewObjectID()
		stringID := objectID.Hex()
		return stringID
	default:
		return ""
	}
}

var (
	ExportRepository = NewRepository()
)
