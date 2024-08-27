package requests

import (
	"context"
	"errors"
	"mime/multipart"
	"testing"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/nhan1603/CloneScc/api/internal/model"
	ftpPkg "github.com/nhan1603/CloneScc/api/internal/pkg/ftp"
	"github.com/nhan1603/CloneScc/api/internal/repository"
	"github.com/nhan1603/CloneScc/api/internal/repository/request"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestImpl_CreateResponse(t *testing.T) {
	type mockGetByID struct {
		mockResp model.VerificationRequest
		mockErr  error
	}

	type mockSaveMedia struct {
		wantMock bool
		mockResp model.MediaData
		mockErr  error
	}

	timeMock := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	timeNowWrapperFunc = func() time.Time {
		return timeMock
	}

	defer func() {
		timeNowWrapperFunc = timeNowWrapper
	}()

	type mockInsertRequestResponses struct {
		wantMock bool
		mockErr  error
	}

	type mockUpdateStatusAndEndTime struct {
		wantMock bool
		mockErr  error
	}

	type arg struct {
		givenInput                 CreateResponseInput
		mockGetByID                mockGetByID
		mockSaveMedia              mockSaveMedia
		mockInsertRequestResponses mockInsertRequestResponses
		mockUpdateStatusAndEndTime mockUpdateStatusAndEndTime
		expErr                     error
	}

	tcs := map[string]arg{
		"success": {
			givenInput: CreateResponseInput{
				RequestID: 1,
				Message:   "message",
				Files: []*multipart.FileHeader{
					{
						Filename: "abc",
						Header:   nil,
						Size:     1233,
					},
				},
			},
			mockGetByID: mockGetByID{
				mockResp: model.VerificationRequest{
					ID:     1,
					Status: model.VerificationRequestStatusNew,
				},
			},
			mockSaveMedia: mockSaveMedia{
				wantMock: true,
				mockResp: model.MediaData{
					FileName:      "abc",
					FileExtension: ".mp4",
					URL:           "http://abc.com",
				},
			},
			mockInsertRequestResponses: mockInsertRequestResponses{
				wantMock: true,
			},
			mockUpdateStatusAndEndTime: mockUpdateStatusAndEndTime{
				wantMock: true,
			},
		},
		"error update status and end time": {
			givenInput: CreateResponseInput{
				RequestID: 1,
				Message:   "message",
				Files: []*multipart.FileHeader{
					{
						Filename: "abc",
						Header:   nil,
						Size:     1233,
					},
				},
			},
			mockGetByID: mockGetByID{
				mockResp: model.VerificationRequest{
					ID:     1,
					Status: model.VerificationRequestStatusNew,
				},
			},
			mockSaveMedia: mockSaveMedia{
				wantMock: true,
				mockResp: model.MediaData{
					FileName:      "abc",
					FileExtension: ".mp4",
					URL:           "http://abc.com",
				},
			},
			mockInsertRequestResponses: mockInsertRequestResponses{
				wantMock: true,
			},
			mockUpdateStatusAndEndTime: mockUpdateStatusAndEndTime{
				wantMock: true,
				mockErr:  errors.New("error"),
			},
			expErr: errors.New("error"),
		},
		"error insert request response": {
			givenInput: CreateResponseInput{
				RequestID: 1,
				Message:   "message",
				Files: []*multipart.FileHeader{
					{
						Filename: "abc",
						Header:   nil,
						Size:     1233,
					},
				},
			},
			mockGetByID: mockGetByID{
				mockResp: model.VerificationRequest{
					ID:     1,
					Status: model.VerificationRequestStatusNew,
				},
			},
			mockSaveMedia: mockSaveMedia{
				wantMock: true,
				mockResp: model.MediaData{
					FileName:      "abc",
					FileExtension: ".mp4",
					URL:           "http://abc.com",
				},
			},
			mockInsertRequestResponses: mockInsertRequestResponses{
				wantMock: true,
				mockErr:  errors.New("error"),
			},
			expErr: errors.New("error"),
		},
		"error save media file": {
			givenInput: CreateResponseInput{
				RequestID: 1,
				Message:   "message",
				Files: []*multipart.FileHeader{
					{
						Filename: "abc",
						Header:   nil,
						Size:     1233,
					},
				},
			},
			mockGetByID: mockGetByID{
				mockResp: model.VerificationRequest{
					ID:     1,
					Status: model.VerificationRequestStatusNew,
				},
			},
			mockSaveMedia: mockSaveMedia{
				wantMock: true,
				mockErr:  errors.New("error"),
			},
			expErr: errors.New("error"),
		},
		"error status is resolved": {
			givenInput: CreateResponseInput{
				RequestID: 1,
				Message:   "message",
				Files: []*multipart.FileHeader{
					{
						Filename: "abc",
						Header:   nil,
						Size:     1233,
					},
				},
			},
			mockGetByID: mockGetByID{
				mockResp: model.VerificationRequest{
					ID:     1,
					Status: model.VerificationRequestStatusResolved,
				},
			},
			expErr: ErrRequestResolved,
		},
		"error get by id": {
			givenInput: CreateResponseInput{
				RequestID: 1,
				Message:   "message",
				Files: []*multipart.FileHeader{
					{
						Filename: "abc",
						Header:   nil,
						Size:     1233,
					},
				},
			},
			mockGetByID: mockGetByID{
				mockErr: errors.New("error"),
			},
			expErr: errors.New("error"),
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			// Given & Mock
			mockRegistry := repository.NewMockRegistry(t)
			mockRequest := request.NewMockRepository(t)
			mockRegistry.On("Request").Return(mockRequest)
			mockRequest.On("GetByID", mock.Anything, tc.givenInput.RequestID).Return(tc.mockGetByID.mockResp, tc.mockGetByID.mockErr)

			if tc.mockSaveMedia.wantMock {
				initFTPClientFunc = func() (*ftp.ServerConn, error) {
					return &ftp.ServerConn{}, nil
				}
				defer func() {
					initFTPClientFunc = ftpPkg.InitFTPClient
				}()

				saveMediaDataFunc = func(ctx context.Context, ftpClient *ftp.ServerConn, requestID int64, file *multipart.FileHeader) (model.MediaData, error) {
					return tc.mockSaveMedia.mockResp, tc.mockSaveMedia.mockErr
				}

				defer func() {
					saveMediaDataFunc = ftpPkg.SaveMediaData
				}()
			}

			if tc.mockInsertRequestResponses.wantMock {
				mockRegistry.On("DoInTx", mock.Anything, mock.Anything).Return(
					func(ctx context.Context, txFunc repository.TxFunc) error {
						return txFunc(mockRegistry)
					})
				mockRequest.On("InsertRequestResponses", mock.Anything, model.VerificationRequestResponses{
					VerificationRequestID: tc.givenInput.RequestID,
					Message:               tc.givenInput.Message,
					MediaData:             []model.MediaData{tc.mockSaveMedia.mockResp},
					VerifiedAt:            timeMock,
				}).Return(model.VerificationRequestResponses{}, tc.mockInsertRequestResponses.mockErr)
			}

			if tc.mockUpdateStatusAndEndTime.wantMock {
				mockRequest.On("UpdateStatusAndEndTime", mock.Anything,
					tc.givenInput.RequestID, model.VerificationRequestStatusResolved, timeMock).Return(model.VerificationRequest{}, tc.mockUpdateStatusAndEndTime.mockErr)
			}
			instance := New(mockRegistry, nil, nil, nil)
			err := instance.CreateResponse(context.Background(), tc.givenInput)
			if err != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
