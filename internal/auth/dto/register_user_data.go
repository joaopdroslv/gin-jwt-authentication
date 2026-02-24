package auth

import "time"

type RegisterUserData struct {
	UserStatusID int64
	Name         string
	Birthdate    time.Time
	Email        string
	PasswordHash string
}
