package models

import "time"

type Member struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func (m *Member) GetMemberID() string {
	return m.ID
}
