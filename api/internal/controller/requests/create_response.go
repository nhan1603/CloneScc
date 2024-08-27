package requests

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/jlaffaye/ftp"
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository"
	"github.com/nhan1603/CloneScc/api/internal/repository/request"
)

// CreateResponseInput is the input type for the CreateResponse controller
type CreateResponseInput struct {
	RequestID int64
	Message   string
	Files     []*multipart.FileHeader
}

// CreateResponse Create a new response
func (i impl) CreateResponse(ctx context.Context, input CreateResponseInput) error {
	log.Println("[CreateResponse] START creating new response in ctrl")
	req, err := i.repo.Request().GetByID(ctx, input.RequestID)
	if err != nil {
		if errors.Is(err, request.ErrNotFound) {
			return ErrNotFound
		}
		log.Println(fmt.Printf("[CreateResponse] Error Request().GetByID: %v", err))
		return err
	}

	if req.Status == model.VerificationRequestStatusResolved {
		return ErrRequestResolved
	}

	// Save files to the media folder
	mediaData := make([]model.MediaData, 0, len(input.Files))
	ftpClient, err := initFTPClientFunc()
	if err != nil {
		log.Printf("error when entablishing ftp connection: %v\n", err)
		return err
	}

	for _, file := range input.Files {
		media, err := i.saveMediaData(ctx, ftpClient, input.RequestID, file)
		if err != nil {
			log.Println(fmt.Printf("[CreateResponse] Error saving media data: %v\n", err))
			return err
		}
		mediaData = append(mediaData, media)
	}

	return i.repo.DoInTx(ctx, func(txRepo repository.Registry) error {
		inp := model.VerificationRequestResponses{
			VerificationRequestID: input.RequestID,
			Message:               input.Message,
			MediaData:             mediaData,
			VerifiedAt:            timeNowWrapperFunc(),
		}

		if _, err := txRepo.Request().InsertRequestResponses(ctx, inp); err != nil {
			log.Println(fmt.Printf("[CreateResponse] Error calling Request.InsertRequestResponses %#v from repository: %v", inp, err))
			return err
		}

		if _, err := txRepo.Request().UpdateStatusAndEndTime(ctx, input.RequestID, model.VerificationRequestStatusResolved, timeNowWrapperFunc()); err != nil {
			log.Println(fmt.Printf("[CreateResponse] Error calling Request.UpdateRequestStatus %#v from repository: %v", inp, err))
			return err
		}

		return nil
	})
}

func (i impl) saveMediaData(ctx context.Context, ftpClient *ftp.ServerConn, reqID int64, file *multipart.FileHeader) (model.MediaData, error) {
	// Generate the final filename
	finalFileName := fmt.Sprintf("req-%d-%s", reqID, file.Filename)

	// Upload the file to the remote FTP server in the root directory
	remoteFilePath := finalFileName

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return model.MediaData{}, fmt.Errorf("error opening the uploaded file: %v", err)
	}
	defer src.Close()

	// Upload the file to the FTP server directly from the source (without saving it locally)
	err = ftpClient.Stor(remoteFilePath, src)
	if err != nil {
		return model.MediaData{}, fmt.Errorf("error uploading file to FTP server: %v", err)
	}

	// Create the URL for accessing the uploaded file on the FTP server
	fileURL := fmt.Sprintf("%s%s", os.Getenv("FILE_URL"), finalFileName)

	return model.MediaData{
		FileName:      finalFileName,
		FileExtension: filepath.Ext(file.Filename),
		URL:           fileURL,
	}, nil
}
