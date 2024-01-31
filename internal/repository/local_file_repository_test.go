package repository

import (
	"filetransfer/api"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalFileRepository_GetFileList(t *testing.T) {
	tempDir := t.TempDir()
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")

	err := os.WriteFile(file1, []byte("content1"), 0644)
	assert.NoError(t, err)
	err = os.WriteFile(file2, []byte("content2"), 0644)
	assert.NoError(t, err)

	repo := NewLocalFileRepository(tempDir)

	fileList, err := repo.GetFileList()
	assert.NoError(t, err)

	assert.ElementsMatch(t, []string{"file1.txt", "file2.txt"}, fileList)
}

func TestLocalFileRepository_GetFileInfo(t *testing.T) {
	tempDir := t.TempDir()
	file := filepath.Join(tempDir, "file.txt")

	err := os.WriteFile(file, []byte("content"), 0644)
	assert.NoError(t, err)

	repo := NewLocalFileRepository(tempDir)

	fileInfoInterface, err := repo.GetFileInfo("file.txt")
	assert.NoError(t, err)

	fileInfo, ok := fileInfoInterface.(*api.FileInfoResponse)
	assert.True(t, ok, "expected *FileInfoResponse")

	expectedFileInfo := &api.FileInfoResponse{
		Filename: "file.txt",
		Size:     7,
	}

	assert.Equal(t, expectedFileInfo, fileInfo)
}

func TestLocalFileRepository_GetFileContent(t *testing.T) {
	tempDir := t.TempDir()
	file := filepath.Join(tempDir, "file.txt")

	err := os.WriteFile(file, []byte("content"), 0644)
	assert.NoError(t, err)

	repo := NewLocalFileRepository(tempDir)

	content, err := repo.GetFileContent("file.txt")
	assert.NoError(t, err)

	assert.Equal(t, []byte("content"), content)
}
