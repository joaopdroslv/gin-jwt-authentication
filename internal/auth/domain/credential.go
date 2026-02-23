package auth

import "time"

type UserInfo struct {
	ID int64
}

type Credential struct {
	ID           int64
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	UserInfo UserInfo
}
