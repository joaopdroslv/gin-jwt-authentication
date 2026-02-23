package auth

type Credential struct {
	ID           int64
	Email        string
	PasswordHash string
}
