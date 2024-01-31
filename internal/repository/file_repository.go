package repository

// FileRepository is an interface defining methods for interacting with file-related operations.
type FileRepository interface {
	// GetFileList returns a list of file names available in the repository.
	GetFileList() ([]string, error)

	// GetFileInfo retrieves metadata information about a specific file identified by its filename.
	// The returned metadata type should encapsulate details like filename, size, etc.
	GetFileInfo(filename string) (interface{}, error)

	// GetFileContent retrieves the content of a specific file identified by its filename.
	GetFileContent(filename string) ([]byte, error)
}
