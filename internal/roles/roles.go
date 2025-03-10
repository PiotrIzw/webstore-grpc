package roles

type Role struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`        // e.g., "admin", "user"
	Permissions []string `json:"permissions"` // e.g., ["orders:write", "accounts:read"]
}

type UserRole struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}
