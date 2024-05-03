package repositories

import (
	"errors"

	"github.com/Matvey1109/LibraryManagementSystemCore/core/models"
)

type BookRepository struct{}

func (br *BookRepository) GetAllBooks() ([]models.Book, error) {
	books, err := storage.GetAllBooksStorage()
	if err != nil {
		return books, err
	}
	return books, nil
}

func (br *BookRepository) GetBook(id string) (models.Book, error) {
	book, err := storage.GetBookStorage(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (br *BookRepository) AddBook(title string, author string, publicationYear int, genre string, totalCopies int) error {
	newID := GenerateID()
	newBook := models.Book{
		ID:              newID,
		Title:           title,
		Author:          author,
		PublicationYear: publicationYear,
		Genre:           genre,
		AvailableCopies: totalCopies,
		TotalCopies:     totalCopies,
	}
	err := storage.AddBookStorage(newBook)
	if err != nil {
		return err
	}
	return nil
}

func (br *BookRepository) UpdateBook(id string, title *string, author *string, publicationYear *int, genre *string, availableCopies *int, totalCopies *int) error {
	book, err := storage.GetBookStorage(id)
	if err != nil {
		return err
	}

	if title != nil {
		book.Title = *title
	}
	if author != nil {
		book.Author = *author
	}
	if publicationYear != nil {
		book.PublicationYear = *publicationYear
	}
	if genre != nil {
		book.Genre = *genre
	}
	if availableCopies != nil {
		if totalCopies != nil {
			if *availableCopies > *totalCopies {
				return errors.New("available copies cannot be greater than total copies")
			}
			book.TotalCopies = *totalCopies
		} else {
			if *availableCopies > book.TotalCopies {
				return errors.New("available copies cannot be greater than total copies")
			}
		}
		book.AvailableCopies = *availableCopies
	}
	if totalCopies != nil {
		book.TotalCopies = *totalCopies
	}

	if err != nil {
		return err
	}

	err = storage.UpdateBookStorage(id, book)
	if err != nil {
		return err
	}
	return err
}

func (br *BookRepository) DeleteBook(id string) error {
	err := storage.DeleteBookStorage(id)
	if err != nil {
		return err
	}
	return nil
}
