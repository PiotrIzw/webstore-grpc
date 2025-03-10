package service

import (
	"errors"
	"fmt"
	"github.com/PiotrIzw/webstore-grcp/internal/pb"
	"github.com/PiotrIzw/webstore-grcp/internal/repository"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileService struct {
	pb.UnimplementedFileServiceServer
	repo repository.FileRepository
}

func NewFileService(repo repository.FileRepository) *FileService {
	return &FileService{repo: repo}
}

// UploadFile handles chunked file uploads.
func (s *FileService) UploadFile(stream pb.FileService_UploadFileServer) error {
	var (
		fileData     []byte
		fileType     string
		fileName     string
		chunkCount   int
		maxChunkSize = 5 * 1024 * 1024 // 5MB
	)

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Validate chunk size
		if len(chunk.Data) > maxChunkSize {
			return errors.New("chunk size exceeds 5MB limit")
		}

		// Append chunk to file data
		fileData = append(fileData, chunk.Data...)
		chunkCount++

		// Capture metadata from the first chunk
		if chunkCount == 1 {
			fileType = chunk.FileType
			fileName = chunk.FileName

			// Validate file type
			if !isValidFileType(fileType, fileName) {
				return errors.New("invalid file type. Only PNG, JPEG, and PDF are allowed")
			}
		}
	}

	// Save the file to storage
	fileID, err := s.saveFile(fileName, fileType, fileData)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	// Save metadata to the database
	err = s.repo.SaveFileMetadata(fileID, fileName, fileType, int64(len(fileData)))
	if err != nil {
		return fmt.Errorf("failed to save file metadata: %v", err)
	}

	// Send response
	return stream.SendAndClose(&pb.UploadFileResponse{FileId: fileID})
}

// saveFile saves the file to disk and returns the file ID.
func (s *FileService) saveFile(fileName, fileType string, data []byte) (string, error) {
	// Generate a unique file ID
	fileID := generateFileID(fileName)

	// Create the file path
	filePath := filepath.Join("uploads", fileID+"_"+fileName)

	// Ensure the uploads directory exists
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create uploads directory: %v", err)
	}

	// Write the file to disk
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return "", fmt.Errorf("failed to write file: %v", err)
	}

	return fileID, nil
}

// isValidFileType checks if the file type is allowed.
func isValidFileType(fileType, fileName string) bool {
	// Validate file type based on MIME type
	switch fileType {
	case "image/png", "image/jpeg", "application/pdf":
		return true
	}

	// Fallback: Validate file type based on file extension
	ext := strings.ToLower(filepath.Ext(fileName))
	switch ext {
	case ".png", ".jpg", ".jpeg", ".pdf":
		return true
	}

	return false
}

// generateFileID generates a unique file ID.
func generateFileID(fileName string) string {
	// Use a combination of timestamp and file name hash
	return fmt.Sprintf("%d_%s", time.Now().Unix(), fileName)
}
