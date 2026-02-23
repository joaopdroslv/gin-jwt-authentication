package auth

type LoginBody struct {
	Email    string
	Password string
}

type RegisterBody struct {
	Name      string
	Birthdate string
	Email     string
	Password  string
}
