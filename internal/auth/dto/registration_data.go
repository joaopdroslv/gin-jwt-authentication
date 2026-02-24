package auth

import "time"

type RegistrationData struct {
	UserStatusID int64
	Name         string
	Birthdate    time.Time
	Email        string
	PasswordHash string
}
