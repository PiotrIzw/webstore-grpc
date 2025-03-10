package repository

import (
	"database/sql"
)

type RolesRepository struct {
	db *sql.DB
}

func NewRolesRepository(db *sql.DB) *RolesRepository {
	return &RolesRepository{db: db}
}

func (r *RolesRepository) AssignRole(userID, roleName string) error {
	query := `INSERT INTO user_roles (user_id, role_id)
              VALUES ($1, (SELECT id FROM roles WHERE name = $2))`
	_, err := r.db.Exec(query, userID, roleName)
	return err
}

func (r *RolesRepository) CheckPermission(userID, permission string) (bool, error) {
	var allowed bool
	query := `SELECT EXISTS (
                  SELECT 1 FROM user_roles ur
                  JOIN roles r ON ur.role_id = r.id
                  WHERE ur.user_id = $1 AND $2 = ANY(r.permissions)
              )`
	err := r.db.QueryRow(query, userID, permission).Scan(&allowed)
	return allowed, err
}
