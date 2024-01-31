package repository

import (
	"filetransfer/api"
	"os"
	"path/filepath"
)

// LocalFileRepository is an implementation of the FileRepository interface for local file storage.
type LocalFileRepository struct {
	storagePath string
}

// NewLocalFileRepository creates a new instance of LocalFileRepository with the specified storage path.
func NewLocalFileRepository(storagePath string) *LocalFileRepository {
	return &LocalFileRepository{
		storagePath: storagePath,
	}
}

// GetFileList retrieves a list of file names available in the local storage.
func (r *LocalFileRepository) GetFileList() ([]string, error) {
	files, err := os.ReadDir(r.storagePath)
	if err != nil {
		return nil, err
	}

	var fileList []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileList = append(fileList, file.Name())
	}

	return fileList, nil
}

// GetFileInfo retrieves metadata information about a specific file from the local storage.
// It returns an interface{}, which encapsulates details like filename and size.
func (r *LocalFileRepository) GetFileInfo(filename string) (interface{}, error) {
	filePath := filepath.Join(r.storagePath, filename)
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	fileMetadata := &api.FileInfoResponse{
		Filename: filename,
		Size:     uint64(fileInfo.Size()),
	}
	return fileMetadata, nil
}

// GetFileContent retrieves the content of a specific file from the local storage.
func (r *LocalFileRepository) GetFileContent(filename string) ([]byte, error) {
	filePath := filepath.Join(r.storagePath, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}
