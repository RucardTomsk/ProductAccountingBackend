package enum

type Roles int

const (
	ADMIN Roles = iota
	NOT_ROLE
)

func (i Roles) String() string {
	return [...]string{"admin", "not_role"}[i]
}

func ParseRoles(rolesString string) Roles {
	switch rolesString {
	case ADMIN.String():
		return ADMIN

	default:
		return NOT_ROLE
	}
}