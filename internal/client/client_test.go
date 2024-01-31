package client

import (
	"errors"
	"filetransfer/api"
	"filetransfer/internal/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestFileTransferClient_GetFileList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := logger.NewMockClientLogger(ctrl)
	mockClient := api.NewMockFileTransferClient(ctrl)

	client := FileTransferClient{
		client: mockClient,
		logger: mockLogger,
	}

	mockLogger.EXPECT().Printf(gomock.Any(), gomock.Any()).AnyTimes()
	mockClient.EXPECT().GetFileList(gomock.Any(), gomock.Any()).Return(&api.FileListResponse{Files: []string{"file1.txt", "file2.txt"}}, nil)

	files, err := client.GetFileList()

	assert.NoError(t, err)
	assert.Equal(t, []string{"file1.txt", "file2.txt"}, files)
}

func TestFileTransferClient_GetFileInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := logger.NewMockClientLogger(ctrl)
	mockClient := api.NewMockFileTransferClient(ctrl)

	client := &FileTransferClient{
		client: mockClient,
		logger: mockLogger,
	}

	mockLogger.EXPECT().Printf(gomock.Any(), gomock.Any()).AnyTimes()
	mockClient.EXPECT().GetFileInfo(gomock.Any(), gomock.Any()).Return(&api.FileInfoResponse{Filename: "file1.txt", Size: 100}, nil)

	fileInfo, err := client.GetFileInfo("file1.txt")

	assert.NoError(t, err)
	assert.NotNil(t, fileInfo)
	assert.Equal(t, "file1.txt", fileInfo.Filename)
	assert.Equal(t, uint64(100), fileInfo.Size)
}

func TestFileTransferClient_GetFileContent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := logger.NewMockClientLogger(ctrl)
	mockClient := api.NewMockFileTransferClient(ctrl)

	client := &FileTransferClient{
		client: mockClient,
		logger: mockLogger,
	}

	mockLogger.EXPECT().Printf(gomock.Any(), gomock.Any()).AnyTimes()
	mockClient.EXPECT().GetFileContent(gomock.Any(), gomock.Any()).Return(&api.FileContentResponse{Filename: "file1.txt", Content: []byte("file content")}, nil)

	fileContent, err := client.GetFileContent("file1.txt")

	assert.NoError(t, err)
	assert.NotNil(t, fileContent)
	assert.Equal(t, "file1.txt", fileContent.Filename)
	assert.Equal(t, []byte("file content"), fileContent.Content)
}

func TestFileTransferClient_GetFileList_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := logger.NewMockClientLogger(ctrl)
	mockClient := api.NewMockFileTransferClient(ctrl)

	client := &FileTransferClient{
		client: mockClient,
		logger: mockLogger,
	}

	mockLogger.EXPECT().Printf(gomock.Any(), gomock.Any()).AnyTimes()
	mockClient.EXPECT().GetFileList(gomock.Any(), gomock.Any()).Return(nil, errors.New("mock error"))

	files, err := client.GetFileList()

	assert.Error(t, err)
	assert.Nil(t, files)
}
