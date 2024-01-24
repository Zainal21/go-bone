package dtos

import "time"

type PersonalAccessTokenDto struct {
	Name         string
	TokenableId  string
	ExpirationAt *time.Time
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

type NewToken struct {
	PlainText string
	Hash      string
}
