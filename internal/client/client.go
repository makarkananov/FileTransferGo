package client

import (
	"context"
	"filetransfer/api"
	"filetransfer/internal/client/client_interceptor"
	"filetransfer/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

// FileTransferClient represents a gRPC client for file transfer operations.
type FileTransferClient struct {
	conn   *grpc.ClientConn
	client api.FileTransferClient
	logger logger.ClientLogger
}

// NewFileTransferClient creates a new FileTransferClient instance.
// It establishes a connection to the gRPC server at the specified address.
func NewFileTransferClient(serverAddress string, logger logger.ClientLogger) (*FileTransferClient, error) {
	conn, err := grpc.Dial(
		serverAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(client_interceptor.ClientLoggingInterceptor(logger)),
	)
	if err != nil {
		return nil, err
	}

	fileTransferClient := api.NewFileTransferClient(conn)

	return &FileTransferClient{
		conn:   conn,
		client: fileTransferClient,
		logger: logger,
	}, nil
}

// Close closes the connection to the gRPC server.
func (c *FileTransferClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// GetFileList retrieves the list of files from the gRPC server.
func (c *FileTransferClient) GetFileList() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Microsecond)
	defer cancel()

	req := &api.FileListRequest{}
	resp, err := c.client.GetFileList(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Files, nil
}

// GetFileInfo retrieves information about a specific file from the gRPC server.
func (c *FileTransferClient) GetFileInfo(filename string) (*api.FileInfoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &api.FileInfoRequest{
		Filename: filename,
	}
	resp, err := c.client.GetFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetFileContent retrieves the content of a specific file from the gRPC server.
func (c *FileTransferClient) GetFileContent(filename string) (*api.FileContentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &api.FileInfoRequest{
		Filename: filename,
	}
	resp, err := c.client.GetFileContent(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
