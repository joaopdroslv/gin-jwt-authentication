package errs

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")

	ErrResourceAlreadyExists = errors.New("this resource already exists")

	ErrInactiveUser                 = errors.New("inactive user, cannot login")
	ErrUserEmailConfirmationPending = errors.New("email confirmation is pending, do it before loggin in")
	ErrUserPasswordCreationPending  = errors.New("password creation is pending, do it before loggin in")
	ErrDeletedUser                  = errors.New("deleted user, cannot login")
)

// TODO: This could be improved to avoid multiple errors.Is conditions
func IsUserStatusRelated(err error) bool {
	return errors.Is(err, ErrInactiveUser) ||
		errors.Is(err, ErrUserEmailConfirmationPending) ||
		errors.Is(err, ErrUserPasswordCreationPending) ||
		errors.Is(err, ErrDeletedUser)
}
