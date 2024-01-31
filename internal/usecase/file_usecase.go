package usecase

import (
	"filetransfer/internal/repository"
)

// FileUsecase represents the use case for file-related operations.
type FileUsecase struct {
	repository repository.FileRepository
}

// NewFileUsecase creates a new instance of FileUsecase with the provided repository.
func NewFileUsecase(repository repository.FileRepository) *FileUsecase {
	return &FileUsecase{
		repository: repository,
	}
}

// GetFileList retrieves the list of files from the underlying repository.
func (u *FileUsecase) GetFileList() ([]string, error) {
	files, err := u.repository.GetFileList()
	if err != nil {
		return nil, err
	}

	return files, nil
}

// GetFileInfo retrieves information about a specific file from the underlying repository.
func (u *FileUsecase) GetFileInfo(filename string) (interface{}, error) {
	fileMetadata, err := u.repository.GetFileInfo(filename)
	if err != nil {
		return nil, err
	}

	return fileMetadata, nil
}

// GetFileContent retrieves the content of a specific file from the underlying repository.
func (u *FileUsecase) GetFileContent(filename string) ([]byte, error) {
	content, err := u.repository.GetFileContent(filename)
	if err != nil {
		return nil, err
	}

	return content, nil
}
