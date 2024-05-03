package repositories

import (
	"errors"
	"time"

	"github.com/Matvey1109/LibraryManagementSystemCore/internal/models"
)

type BorrowingRepository struct{}

func (br *BorrowingRepository) GetAllBorrowings() ([]models.Borrowing, error) {
	borrowings, err := storage.GetAllBorrowingsStorage()
	if err != nil {
		return borrowings, err
	}
	return borrowings, nil
}

func (br *BorrowingRepository) GetMemberBooks(memberID string) ([]models.Book, error) {
	borrowings, err := storage.GetAllBorrowingsStorage()
	if err != nil {
		return nil, err
	}

	memberBorrowings := make([]models.Borrowing, 0)
	for _, borrowing := range borrowings {
		if borrowing.MemberID == memberID {
			memberBorrowings = append(memberBorrowings, borrowing)
		}
	}

	books := make([]models.Book, 0)
	for _, borrowing := range memberBorrowings {
		book, err := storage.GetBookStorage(borrowing.BookID)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (br *BorrowingRepository) BorrowBook(bookID string, memberID string, borrowYear int) error {
	if borrowYear > time.Now().Year() {
		return errors.New("invalid year")
	}

	newID := GenerateID()
	newBorrowing := models.Borrowing{
		ID:         newID,
		BookID:     bookID,
		MemberID:   memberID,
		BorrowYear: borrowYear,
		ReturnYear: -1,
	}
	err := storage.AddBorrowingStorage(newBorrowing)
	if err != nil {
		return err
	}
	return nil
}

func (br *BorrowingRepository) ReturnBook(id string) error {
	borrowing, err := storage.GetBorrowingStorage(id)
	if err != nil {
		return err
	}

	borrowing.ReturnYear = time.Now().Year()

	err = storage.UpdateBorrowingStorage(id, borrowing)
	if err != nil {
		return err
	}
	return nil
}
