package ftp

import (
	"fmt"
	"os"

	"github.com/jlaffaye/ftp"
)

// InitFTPClient to initialize FTP client
func InitFTPClient() (*ftp.ServerConn, error) {
	// Define the FTP server details
	ftpServer := os.Getenv("FTP_SERVER")
	ftpPort := 21
	ftpUsername := os.Getenv("FTP_USER")
	ftpPassword := os.Getenv("FTP_PASSWORD")

	// Connect to the FTP server
	client, err := ftp.Dial(fmt.Sprintf("%s:%d", ftpServer, ftpPort))
	if err != nil {
		return nil, fmt.Errorf("error connecting to FTP server: %v", err)
	}

	// Login to the FTP server
	err = client.Login(ftpUsername, ftpPassword)
	if err != nil {
		client.Quit()
		return nil, fmt.Errorf("error logging in to FTP server: %v", err)
	}

	return client, nil
}
