package server

import (
	"context"
	"filetransfer/api"
	"filetransfer/internal/logger"
	"filetransfer/internal/server/server_interceptor"
	"filetransfer/internal/usecase"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"

	"google.golang.org/grpc"
)

// FileTransferServer represents the gRPC server for file transfer operations.
type FileTransferServer struct {
	fileUsecase *usecase.FileUsecase
	server      *grpc.Server
	logger      logger.ServerLogger
	api.UnimplementedFileTransferServer
}

// NewFileTransferServer creates a new instance of FileTransferServer.
func NewFileTransferServer(fileUsecase *usecase.FileUsecase, logger logger.ServerLogger) *FileTransferServer {
	return &FileTransferServer{
		fileUsecase: fileUsecase,
		logger:      logger,
	}
}

// Start starts the gRPC server on the specified port.
func (s *FileTransferServer) Start(port int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		s.logger.Printf("Error starting listener: %v", err)
		return err
	}

	s.server = grpc.NewServer(
		grpc.ChainUnaryInterceptor(server_interceptor.LoggingInterceptor(s.logger), server_interceptor.ValidationInterceptor()),
	)
	api.RegisterFileTransferServer(s.server, s)

	s.logger.Printf("gRPC server started on :%d\n", port)

	go func() {
		if err := s.server.Serve(listen); err != nil {
			s.logger.Printf("Error serving gRPC: %v", err)
		}
	}()

	return nil
}

// Stop stops the gRPC server gracefully.
func (s *FileTransferServer) Stop() {
	if s.server != nil {
		s.server.GracefulStop()
		s.logger.Printf("gRPC server stopped")
	}
}

// handleError handles errors and returns a gRPC status with the appropriate code.
func handleError(err error, msg string, code codes.Code) error {
	if err != nil {
		return status.Errorf(code, "%s: %v", msg, err)
	}
	return nil
}

// GetFileList retrieves the list of files from the repository.
func (s *FileTransferServer) GetFileList(ctx context.Context, req *api.FileListRequest) (*api.FileListResponse, error) {
	files, err := s.fileUsecase.GetFileList()
	if err != nil {
		return nil, handleError(err, "Error getting file list", codes.Internal)
	}

	return &api.FileListResponse{Files: files}, nil
}

// GetFileInfo retrieves information about a specific file from the repository.
func (s *FileTransferServer) GetFileInfo(ctx context.Context, req *api.FileInfoRequest) (*api.FileInfoResponse, error) {
	fileMetadata, err := s.fileUsecase.GetFileInfo(req.Filename)
	if err != nil {
		return nil, handleError(err, "Error getting file metadata", codes.NotFound)
	}

	return fileMetadata.(*api.FileInfoResponse), nil
}

// GetFileContent retrieves the content of a specific file from the repository.
func (s *FileTransferServer) GetFileContent(ctx context.Context, req *api.FileInfoRequest) (*api.FileContentResponse, error) {
	content, err := s.fileUsecase.GetFileContent(req.Filename)
	if err != nil {
		return nil, handleError(err, "Error getting file content", codes.Internal)
	}

	return &api.FileContentResponse{Filename: req.Filename, Content: content}, nil
}
