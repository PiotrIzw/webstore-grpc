package preferences

type Preferences struct {
	UserID        string `json:"user_id"`
	Theme         string `json:"theme"`         // e.g., "light", "dark"
	Notifications bool   `json:"notifications"` // e.g., true/false
	Locale        string `json:"locale"`        // e.g., "en", "fr"
}
