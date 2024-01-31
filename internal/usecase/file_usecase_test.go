package usecase

import (
	"errors"
	"filetransfer/internal/repository"
	"go.uber.org/mock/gomock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileUsecase_GetFileList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockFileRepository(ctrl)
	usecase := NewFileUsecase(mockRepo)

	mockRepo.EXPECT().GetFileList().Return([]string{"file1.txt", "file2.txt"}, nil)

	files, err := usecase.GetFileList()

	assert.NoError(t, err)
	assert.Equal(t, []string{"file1.txt", "file2.txt"}, files)
}

func TestFileUsecase_GetFileInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockFileRepository(ctrl)
	usecase := NewFileUsecase(mockRepo)

	mockRepo.EXPECT().GetFileInfo("file1.txt").Return("file info", nil)

	fileInfo, err := usecase.GetFileInfo("file1.txt")

	assert.NoError(t, err)
	assert.Equal(t, "file info", fileInfo)
}

func TestFileUsecase_GetFileContent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockFileRepository(ctrl)
	usecase := NewFileUsecase(mockRepo)

	mockRepo.EXPECT().GetFileContent("file1.txt").Return([]byte("file content"), nil)

	content, err := usecase.GetFileContent("file1.txt")

	assert.NoError(t, err)
	assert.Equal(t, []byte("file content"), content)
}

func TestFileUsecase_GetFileList_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockFileRepository(ctrl)
	usecase := NewFileUsecase(mockRepo)

	mockRepo.EXPECT().GetFileList().Return(nil, errors.New("mock error"))

	files, err := usecase.GetFileList()

	assert.Error(t, err)
	assert.Nil(t, files)
}
