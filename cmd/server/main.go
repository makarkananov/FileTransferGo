package main

import (
	"filetransfer/internal/repository"
	"filetransfer/internal/server"
	"filetransfer/internal/usecase"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initialize the server logger
	logger := log.New(os.Stdout, "[Server] ", log.LstdFlags)

	// Create a new instance of the local file repository with the root directory "/"
	fileRepository := repository.NewLocalFileRepository("/")

	// Create a new file usecase with the file repository
	fileUsecase := usecase.NewFileUsecase(fileRepository)

	// Create a new file transfer server with the file usecase and logger
	fileServer := server.NewFileTransferServer(fileUsecase, logger)

	// Start the server in a separate goroutine
	go func() {
		if err := fileServer.Start(50051); err != nil {
			logger.Printf("Error starting server: %v", err)
		}
	}()

	// Set up a signal channel to handle interrupt and termination signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a signal
	sig := <-sigCh
	logger.Printf("Received signal %v. Shutting down...", sig)

	// Stop the server gracefully
	fileServer.Stop()
}
