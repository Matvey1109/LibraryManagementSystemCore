package storages

import (
	"context"
	"errors"
	"time"

	"github.com/Matvey1109/LibraryManagementSystemCore/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ! Implements interface Storage
type MongoDBStorage struct {
	client               *mongo.Client
	membersCollection    *mongo.Collection
	booksCollection      *mongo.Collection
	borrowingsCollection *mongo.Collection
}

var _ Storage = (*MongoDBStorage)(nil) // Checker

// * Member
func (ms *MongoDBStorage) GetAllMembersStorage() ([]models.Member, error) {
	cursor, err := ms.membersCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var members []models.Member
	for cursor.Next(context.Background()) {
		var (
			mongoMemberMap map[string]interface{}
			member         models.Member
		)

		err := cursor.Decode(&mongoMemberMap)
		if err != nil {
			return nil, err
		}

		for key, value := range mongoMemberMap {
			if key == "_id" {
				member.ID = value.(primitive.ObjectID).Hex()
			} else if key == "name" {
				member.Name = value.(string)
			} else if key == "address" {
				member.Address = value.(string)
			} else if key == "email" {
				member.Email = value.(string)
			} else if key == "createdAt" {
				curTime := value.(primitive.DateTime).Time()
				member.CreatedAt, _ = time.Parse(time.DateTime, curTime.Format(time.DateTime))
			}
		}
		members = append(members, member)
	}
	return members, nil
}

func (ms *MongoDBStorage) GetMemberStorage(id string) (models.Member, error) {
	var member models.Member
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return member, err
	}

	filter := bson.M{"_id": objectID}
	err = ms.membersCollection.FindOne(context.Background(), filter).Decode(&member)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return member, errors.New("member not found")
		}
		return member, err
	}
	member.ID = id

	return member, nil
}

func (ms *MongoDBStorage) AddMemberStorage(member models.Member) error {
	newMember := bson.D{
		{Key: "name", Value: member.Name},
		{Key: "address", Value: member.Address},
		{Key: "email", Value: member.Email},
		{Key: "createdAt", Value: member.CreatedAt},
	}
	_, err := ms.membersCollection.InsertOne(context.Background(), newMember)
	return err
}

func (ms *MongoDBStorage) UpdateMemberStorage(id string, member models.Member) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	updatedMember := bson.D{
		{Key: "name", Value: member.Name},
		{Key: "address", Value: member.Address},
		{Key: "email", Value: member.Email},
	}

	_, err = ms.membersCollection.UpdateOne(context.Background(), filter, bson.D{{Key: "$set", Value: updatedMember}})
	if err != nil {
		return err
	}

	return nil
}

func (ms *MongoDBStorage) DeleteMemberStorage(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	result, err := ms.membersCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("member not found")
	}

	return nil
}

// * Book
func (ms *MongoDBStorage) GetAllBooksStorage() ([]models.Book, error) {
	cursor, err := ms.booksCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var books []models.Book
	for cursor.Next(context.Background()) {
		var (
			mongoBookMap map[string]interface{}
			book         models.Book
		)

		err := cursor.Decode(&mongoBookMap)
		if err != nil {
			return nil, err
		}

		for key, value := range mongoBookMap {
			if key == "_id" {
				book.ID = value.(primitive.ObjectID).Hex()
			} else if key == "title" {
				book.Title = value.(string)
			} else if key == "author" {
				book.Author = value.(string)
			} else if key == "publicationYear" {
				if value, ok := value.(int32); ok {
					book.PublicationYear = int(value)
				}
			} else if key == "genre" {
				book.Genre = value.(string)
			} else if key == "availableCopies" {
				if value, ok := value.(int32); ok {
					book.AvailableCopies = int(value)
				}
			} else if key == "totalCopies" {
				if value, ok := value.(int32); ok {
					book.TotalCopies = int(value)
				}
			}
		}
		books = append(books, book)
	}
	return books, nil
}

func (ms *MongoDBStorage) GetBookStorage(id string) (models.Book, error) {
	var book models.Book
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return book, err
	}

	filter := bson.M{"_id": objectID}
	err = ms.booksCollection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return book, errors.New("book not found")
		}
		return book, err
	}
	book.ID = id

	return book, nil
}

func (ms *MongoDBStorage) AddBookStorage(book models.Book) error {
	newBook := bson.D{
		{Key: "title", Value: book.Title},
		{Key: "author", Value: book.Author},
		{Key: "publicationYear", Value: book.PublicationYear},
		{Key: "genre", Value: book.Genre},
		{Key: "availableCopies", Value: book.AvailableCopies},
		{Key: "totalCopies", Value: book.TotalCopies},
	}
	_, err := ms.booksCollection.InsertOne(context.Background(), newBook)
	return err
}

func (ms *MongoDBStorage) UpdateBookStorage(id string, book models.Book) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	updatedBook := bson.D{
		{Key: "title", Value: book.Title},
		{Key: "author", Value: book.Author},
		{Key: "publicationYear", Value: book.PublicationYear},
		{Key: "genre", Value: book.Genre},
		{Key: "availableCopies", Value: book.AvailableCopies},
		{Key: "totalCopies", Value: book.TotalCopies},
	}

	_, err = ms.booksCollection.UpdateOne(context.Background(), filter, bson.D{{Key: "$set", Value: updatedBook}})
	if err != nil {
		return err
	}

	return nil
}

func (ms *MongoDBStorage) DeleteBookStorage(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	result, err := ms.booksCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("book not found")
	}

	return nil
}

// * Borrowing
func (ms *MongoDBStorage) GetAllBorrowingsStorage() ([]models.Borrowing, error) {
	cursor, err := ms.borrowingsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var borrowings []models.Borrowing
	for cursor.Next(context.Background()) {
		var (
			mongoBorrowingMap map[string]interface{}
			borrowing         models.Borrowing
		)

		err := cursor.Decode(&mongoBorrowingMap)
		if err != nil {
			return nil, err
		}

		for key, value := range mongoBorrowingMap {
			if key == "_id" {
				borrowing.ID = value.(primitive.ObjectID).Hex()
			} else if key == "bookId" {
				borrowing.BookID = value.(string)
			} else if key == "memberId" {
				borrowing.MemberID = value.(string)
			} else if key == "borrowYear" {
				if value, ok := value.(int32); ok {
					borrowing.BorrowYear = int(value)
				}
			} else if key == "returnYear" {
				if value, ok := value.(int32); ok {
					borrowing.ReturnYear = int(value)
				}
			}
		}
		borrowings = append(borrowings, borrowing)
	}
	return borrowings, nil
}

func (ms *MongoDBStorage) GetBorrowingStorage(id string) (models.Borrowing, error) {
	var borrowing models.Borrowing
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return borrowing, err
	}

	filter := bson.M{"_id": objectID}
	err = ms.borrowingsCollection.FindOne(context.Background(), filter).Decode(&borrowing)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return borrowing, errors.New("borrowing not found")
		}
		return borrowing, err
	}
	borrowing.ID = id

	return borrowing, nil
}

func (ms *MongoDBStorage) AddBorrowingStorage(borrowing models.Borrowing) error {
	newBorrowing := bson.D{
		{Key: "bookId", Value: borrowing.BookID},
		{Key: "memberId", Value: borrowing.MemberID},
		{Key: "borrowYear", Value: borrowing.BorrowYear},
		{Key: "returnYear", Value: borrowing.ReturnYear},
	}
	_, err := ms.borrowingsCollection.InsertOne(context.Background(), newBorrowing)
	return err
}

func (ms *MongoDBStorage) UpdateBorrowingStorage(id string, borrowing models.Borrowing) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	updatedBorrowing := bson.D{
		{Key: "bookId", Value: borrowing.BookID},
		{Key: "memberId", Value: borrowing.MemberID},
		{Key: "borrowYear", Value: borrowing.BorrowYear},
		{Key: "returnYear", Value: borrowing.ReturnYear},
	}

	_, err = ms.borrowingsCollection.UpdateOne(context.Background(), filter, bson.D{{Key: "$set", Value: updatedBorrowing}})
	if err != nil {
		return err
	}

	return nil
}

func (ms *MongoDBStorage) DeleteBorrowingStorage(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	result, err := ms.borrowingsCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("borrowing not found")
	}

	return nil
}
