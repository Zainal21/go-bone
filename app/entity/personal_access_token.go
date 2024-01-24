package entity

import (
	"time"
)

type PersonalAccessToken struct {
	Id           int        `json:"id" db:"id"`
	UserId       string     `json:"tokenable_id" db:"tokenable_id"`
	Name         string     `json:"name" db:"name"`
	Token        string     `json:"token" db:"token"`
	LastUsedAt   *time.Time `json:"last_used_at" db:"last_used_at"`
	ExpirationAt *time.Time `json:"expires_at" db:"expires_at"`
	Abilities    []string   `json:"abilities" db:"abilities"`
	CreatedAt    *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
}
