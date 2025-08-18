package authentication

import "github.com/rnd-varnion/utils/permission"

var (
	ACCESS_SECRET_KEY           string = "ACCESS_SECRET_KEY"
	REFRESH_SECRET_KEY          string = "REFRESH_SECRET_KEY"
	INTERNAL_ACCESS_SECRET_KEY  string = "INTERNAL_ACCESS_SECRET_KEY"
	JWT_EXPIRY_INTERNAL_SECONDS string = "JWT_EXPIRY_INTERNAL_SECONDS"
)

const (
	UserIDKey         = "user_id"
	IsFromInternalKey = "is_from_internal"
	SessionKey        = "session"
)

type UserSession struct {
	Name        string                   `json:"name"`
	Email       string                   `json:"email"`
	Role        string                   `json:"role"`
	Permissions []*permission.Permission `json:"permissions"`
}
