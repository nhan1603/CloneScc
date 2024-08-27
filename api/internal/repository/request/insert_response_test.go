package request

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestImpl_InsertRequestResponses(t *testing.T) {
	type mockGenerateRequestResponseIDFuncGenerator struct {
		expID  int64
		expErr error
	}
	type arg struct {
		mockGenerateRequestResponseIDFuncGenerator mockGenerateRequestResponseIDFuncGenerator
		givenModel                                 model.VerificationRequestResponses
		expResult                                  model.VerificationRequestResponses
		expErr                                     error
	}

	tcs := map[string]arg{
		"err - from snowflake": {
			mockGenerateRequestResponseIDFuncGenerator: mockGenerateRequestResponseIDFuncGenerator{
				expID:  0,
				expErr: errors.New(`snowflake mock error`),
			},
			expErr: errors.New(`snowflake mock error`),
		},
		"err - request id not exist": {
			givenModel: model.VerificationRequestResponses{
				VerificationRequestID: 989898988,
				Message:               "abc",
				MediaData: []model.MediaData{
					{
						FileName:      "abc",
						FileExtension: ".mp4",
						URL:           "http://abc.com/abc.mp4",
					},
				},
				VerifiedAt: time.Now(),
			},
			mockGenerateRequestResponseIDFuncGenerator: mockGenerateRequestResponseIDFuncGenerator{
				expID: 2,
			},
			expErr: errors.New("dbmodel: unable to insert into verification_request_responses: pq: insert or update on table \"verification_request_responses\" violates foreign key constraint \"verification_request_responses_verification_request_id_fkey\""),
		},
		"success": {
			givenModel: model.VerificationRequestResponses{
				VerificationRequestID: 400,
				Message:               "abc",
				MediaData: []model.MediaData{
					{
						FileName:      "abc",
						FileExtension: ".mp4",
						URL:           "http://abc.com/abc.mp4",
					},
				},
				VerifiedAt: time.Now(),
			},
			mockGenerateRequestResponseIDFuncGenerator: mockGenerateRequestResponseIDFuncGenerator{
				expID: 2,
			},
			expResult: model.VerificationRequestResponses{
				ID:                    2,
				VerificationRequestID: 400,
				Message:               "abc",
				MediaData: []model.MediaData{
					{
						FileName:      "abc",
						FileExtension: ".mp4",
						URL:           "http://abc.com/abc.mp4",
					},
				},
				VerifiedAt: time.Now(),
			},
		},
	}
	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			ctx := context.Background()
			test.WithTxDB(t, func(tx *sql.Tx) {
				// Given
				test.LoadSqlTestFile(t, tx, "testdata/insert_request_response.sql")

				// Mock
				generateRequestResponseIDFunc = func() (int64, error) {
					return tc.mockGenerateRequestResponseIDFuncGenerator.expID, tc.mockGenerateRequestResponseIDFuncGenerator.expErr
				}
				defer func() { generateRequestResponseIDFunc = generateRequestID }()

				// When
				instance := New(tx)
				result, err := instance.InsertRequestResponses(ctx, tc.givenModel)

				// Then
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.NotNil(t, result.ID)
					require.NotZero(t, result.CreatedAt)
					require.NotZero(t, result.UpdatedAt)
					require.NotZero(t, result.VerifiedAt)
					if !cmp.Equal(tc.expResult, result,
						cmpopts.IgnoreFields(model.VerificationRequestResponses{}, "CreatedAt", "UpdatedAt", "VerifiedAt")) {
						t.Errorf("\n VerificationRequest mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expResult, result,
							cmp.Diff(tc.expResult, result, cmpopts.IgnoreFields(model.VerificationRequestResponses{}, "CreatedAt", "UpdatedAt", "VerifiedAt")))
						t.FailNow()
					}
				}
			})
		})
	}
}
