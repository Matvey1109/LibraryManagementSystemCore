package storages

import (
	"errors"

	"github.com/Matvey1109/LibraryManagementSystemCore/core/models"
)

// ! Implements interface Storage
type LocalStorage struct {
	members    []models.Member
	books      []models.Book
	borrowings []models.Borrowing
}

var _ Storage = (*LocalStorage)(nil) // Checker

// * Member
func (ls *LocalStorage) GetAllMembersStorage() ([]models.Member, error) {
	return ls.members, nil
}

func (ls *LocalStorage) GetMemberStorage(id string) (models.Member, error) {
	for _, member := range ls.members {
		if member.ID == id {
			return member, nil
		}
	}
	return models.Member{}, errors.New("member not found")
}

func (ls *LocalStorage) AddMemberStorage(member models.Member) error {
	ls.members = append(ls.members, member)
	return nil
}

func (ls *LocalStorage) UpdateMemberStorage(id string, member models.Member) error {
	for i, m := range ls.members {
		if m.ID == id {
			ls.members[i] = member
			return nil
		}
	}
	return errors.New("member not found")
}

func (ls *LocalStorage) DeleteMemberStorage(id string) error {
	for i, m := range ls.members {
		if m.ID == id {
			ls.members = append(ls.members[:i], ls.members[i+1:]...)
			return nil
		}
	}
	return errors.New("member not found")
}

// * Book
func (ls *LocalStorage) GetAllBooksStorage() ([]models.Book, error) {
	return ls.books, nil
}

func (ls *LocalStorage) GetBookStorage(id string) (models.Book, error) {
	for _, book := range ls.books {
		if book.ID == id {
			return book, nil
		}
	}
	return models.Book{}, errors.New("book not found")
}

func (ls *LocalStorage) AddBookStorage(book models.Book) error {
	ls.books = append(ls.books, book)
	return nil
}

func (ls *LocalStorage) UpdateBookStorage(id string, book models.Book) error {
	for i, b := range ls.books {
		if b.ID == id {
			ls.books[i] = book
			return nil
		}
	}
	return errors.New("book not found")
}

func (ls *LocalStorage) DeleteBookStorage(id string) error {
	for i, b := range ls.books {
		if b.ID == id {
			ls.books = append(ls.books[:i], ls.books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

// * Borrowing
func (ls *LocalStorage) GetAllBorrowingsStorage() ([]models.Borrowing, error) {
	return ls.borrowings, nil
}

func (ls *LocalStorage) GetBorrowingStorage(id string) (models.Borrowing, error) {
	for _, borrowing := range ls.borrowings {
		if borrowing.ID == id {
			return borrowing, nil
		}
	}
	return models.Borrowing{}, errors.New("borrowing not found")
}

func (ls *LocalStorage) AddBorrowingStorage(borrowing models.Borrowing) error {
	ls.borrowings = append(ls.borrowings, borrowing)
	return nil
}

func (ls *LocalStorage) UpdateBorrowingStorage(id string, borrowing models.Borrowing) error {
	for i, b := range ls.borrowings {
		if b.ID == id {
			ls.borrowings[i] = borrowing
			return nil
		}
	}
	return errors.New("borrowing not found")
}

func (ls *LocalStorage) DeleteBorrowingStorage(id string) error {
	for i, b := range ls.borrowings {
		if b.ID == id {
			ls.borrowings = append(ls.borrowings[:i], ls.borrowings[i+1:]...)
			return nil
		}
	}
	return errors.New("borrowing not found")
}
