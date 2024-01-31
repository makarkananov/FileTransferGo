package main

import (
	"filetransfer/internal/client"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// Create a new CLI application
	app := cli.NewApp()
	app.Name = "FileTransferClient"
	app.Usage = "CLI Client for File Transfer gRPC Service"

	// Define a command-line flag for specifying the gRPC server address
	var serverAddress string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "server",
			Value:       "localhost:50051",
			Usage:       "Address of the gRPC server",
			Destination: &serverAddress,
		},
	}

	// Define CLI commands for interacting with the file transfer service
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "Get the list of files on the server",
			Action: func(c *cli.Context) error {
				// Create a logger for the client
				clientLogger := log.New(os.Stdout, "[Client] ", log.LstdFlags)

				// Create a new file transfer client
				fileTransferClient, err := client.NewFileTransferClient(serverAddress, clientLogger)
				if err != nil {
					return err
				}
				defer fileTransferClient.Close()

				// Retrieve the list of files from the server
				files, err := fileTransferClient.GetFileList()
				if err != nil {
					return err
				}

				// Print the list of files
				fmt.Println("Files:")
				for _, file := range files {
					fmt.Printf("- %s\n", file)
				}

				return nil
			},
		},
		{
			Name:    "info",
			Aliases: []string{"i"},
			Usage:   "Get information about a specific file",
			Action: func(c *cli.Context) error {
				// Create a logger for the client
				clientLogger := log.New(os.Stdout, "[Client] ", log.LstdFlags)

				// Create a new file transfer client
				fileTransferClient, err := client.NewFileTransferClient(serverAddress, clientLogger)
				if err != nil {
					return err
				}
				defer fileTransferClient.Close()

				// Retrieve the filename from the command-line arguments
				filename := c.Args().First()
				if filename == "" {
					return fmt.Errorf("please provide a filename")
				}

				// Retrieve information about the specified file from the server
				fileInfo, err := fileTransferClient.GetFileInfo(filename)
				if err != nil {
					return err
				}

				// Print the file information
				fmt.Printf("File information for %s:\n%s\n", filename, fileInfo)

				return nil
			},
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Get content of a specific file",
			Action: func(c *cli.Context) error {
				// Create a logger for the client
				clientLogger := log.New(os.Stdout, "[Client] ", log.LstdFlags)

				// Create a new file transfer client
				fileTransferClient, err := client.NewFileTransferClient(serverAddress, clientLogger)
				if err != nil {
					return err
				}
				defer fileTransferClient.Close()

				// Retrieve the filename from the command-line arguments
				filename := c.Args().First()
				if filename == "" {
					return fmt.Errorf("please provide a filename")
				}

				// Retrieve the content of the specified file from the server
				fileContent, err := fileTransferClient.GetFileContent(filename)
				if err != nil {
					return err
				}

				// Print the file content
				fmt.Printf("File content for %s:\n%s\n", filename, fileContent.Content)

				return nil
			},
		},
	}

	// Run the CLI application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
