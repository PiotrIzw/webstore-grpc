package repository

import (
	"database/sql"
)

type FileRepository interface {
	SaveFileMetadata(fileID, fileName, fileType string, fileSize int64) error
}

type fileRepository struct {
	db *sql.DB
}

func NewFileRepository(db *sql.DB) FileRepository {
	return &fileRepository{db: db}
}

func (r *fileRepository) SaveFileMetadata(fileID, fileName, fileType string, fileSize int64) error {
	query := `INSERT INTO files (id, name, type, size) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, fileID, fileName, fileType, fileSize)
	return err
}
