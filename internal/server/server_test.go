package server

import (
	"context"
	"errors"
	"filetransfer/api"
	"filetransfer/internal/logger"
	"filetransfer/internal/repository"
	"filetransfer/internal/usecase"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
)

func TestFileTransferServer_GetFileList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockFileRepository(ctrl)
	fileUsecase := usecase.NewFileUsecase(mockRepo)
	server := NewFileTransferServer(fileUsecase, &logger.MockServerLogger{})

	mockRepo.EXPECT().GetFileList().Return([]string{"file1.txt", "file2.txt"}, nil)

	resp, err := server.GetFileList(context.Background(), &api.FileListRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, []string{"file1.txt", "file2.txt"}, resp.Files)
}

func TestFileTransferServer_GetFileInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockFileRepository(ctrl)
	fileUsecase := usecase.NewFileUsecase(mockRepo)
	server := NewFileTransferServer(fileUsecase, &logger.MockServerLogger{})

	mockRepo.EXPECT().GetFileInfo("file1.txt").Return(&api.FileInfoResponse{Filename: "file1.txt", Size: 100}, nil)

	resp, err := server.GetFileInfo(context.Background(), &api.FileInfoRequest{Filename: "file1.txt"})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "file1.txt", resp.Filename)
	assert.Equal(t, uint64(100), resp.Size)
}

func TestFileTransferServer_GetFileContent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockFileRepository(ctrl)
	fileUsecase := usecase.NewFileUsecase(mockRepo)
	server := NewFileTransferServer(fileUsecase, &logger.MockServerLogger{})

	mockRepo.EXPECT().GetFileContent("file1.txt").Return([]byte("file content"), nil)

	resp, err := server.GetFileContent(context.Background(), &api.FileInfoRequest{Filename: "file1.txt"})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "file1.txt", resp.Filename)
	assert.Equal(t, []byte("file content"), resp.Content)
}

func TestFileTransferServer_GetFileList_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockFileRepository(ctrl)
	fileUsecase := usecase.NewFileUsecase(mockRepo)
	server := NewFileTransferServer(fileUsecase, &logger.MockServerLogger{})

	mockRepo.EXPECT().GetFileList().Return(nil, errors.New("mock error"))

	resp, err := server.GetFileList(context.Background(), &api.FileListRequest{})

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, codes.Internal, status.Code(err))
}
