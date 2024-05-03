package models

type Book struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publicationYear"`
	Genre           string `json:"genre"`
	AvailableCopies int    `json:"availableCopies"`
	TotalCopies     int    `json:"totalCopies"`
}

func (b *Book) GetBookID() string {
	return b.ID
}
