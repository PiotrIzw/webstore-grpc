package repository

import (
	"database/sql"
	"github.com/PiotrIzw/webstore-grcp/internal/preferences"
)

type PreferencesRepository struct {
	db *sql.DB
}

func NewPreferencesRepository(db *sql.DB) *PreferencesRepository {
	return &PreferencesRepository{db: db}
}

func (r *PreferencesRepository) UpdatePreferences(pref *preferences.Preferences) error {
	query := `INSERT INTO preferences (user_id, theme, notifications, locale)
              VALUES ($1, $2, $3, $4)
              ON CONFLICT (user_id) DO UPDATE
              SET theme = $2, notifications = $3, locale = $4`
	_, err := r.db.Exec(query, pref.UserID, pref.Theme, pref.Notifications, pref.Locale)
	return err
}

func (r *PreferencesRepository) GetPreferences(userID string) (*preferences.Preferences, error) {
	var pref preferences.Preferences
	query := `SELECT theme, notifications, locale FROM preferences WHERE user_id = $1`
	err := r.db.QueryRow(query, userID).Scan(&pref.Theme, &pref.Notifications, &pref.Locale)
	if err != nil {
		return nil, err
	}
	return &pref, nil
}
