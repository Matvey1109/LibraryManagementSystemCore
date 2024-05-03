package models

type Borrowing struct {
	ID         string `json:"id"`
	BookID     string `json:"bookId"`
	MemberID   string `json:"memberId"`
	BorrowYear int    `json:"borrowYear"`
	ReturnYear int    `json:"returnYear,omitempty"`
}

func (b *Borrowing) GetBorrowingID() string {
	return b.ID
}
