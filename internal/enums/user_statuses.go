package enums

type UserStatusID int64

const (
	Active                   UserStatusID = 1
	Inactive                 UserStatusID = 2
	PasswordCreationPending  UserStatusID = 3
	EmailConfirmationPending UserStatusID = 4
	Deleted                  UserStatusID = 5
)
