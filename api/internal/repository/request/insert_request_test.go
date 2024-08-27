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

func TestImpl_Insert(t *testing.T) {
	type mockGenerateRequestIDFuncGenerator struct {
		expID  int64
		expErr error
	}
	type arg struct {
		mockGenerateRequestIDFuncGenerator mockGenerateRequestIDFuncGenerator
		givenModel                         model.VerificationRequest
		expResult                          model.VerificationRequest
		expErr                             error
	}

	tcs := map[string]arg{
		"err - from snowflake": {
			mockGenerateRequestIDFuncGenerator: mockGenerateRequestIDFuncGenerator{
				expID:  0,
				expErr: errors.New(`snowflake mock error`),
			},
			expErr: errors.New(`snowflake mock error`),
		},
		"err - alert id not exist": {
			givenModel: model.VerificationRequest{
				AlertID:        1,
				RequestBy:      2,
				AssignedUserID: 3,
				Message:        "abc",
				Status:         model.VerificationRequestStatusNew,
				StartTime:      time.Now(),
			},
			mockGenerateRequestIDFuncGenerator: mockGenerateRequestIDFuncGenerator{
				expID: 2,
			},
			expErr: errors.New("dbmodel: unable to insert into verification_requests: pq: insert or update on table \"verification_requests\" violates foreign key constraint \"verification_requests_alert_id_fkey\""),
		},
		"success": {
			givenModel: model.VerificationRequest{
				AlertID:        300,
				RequestBy:      102,
				AssignedUserID: 101,
				Message:        "abc",
				Status:         model.VerificationRequestStatusNew,
				StartTime:      time.Now(),
			},
			mockGenerateRequestIDFuncGenerator: mockGenerateRequestIDFuncGenerator{
				expID: 2,
			},
			expResult: model.VerificationRequest{
				ID:             2,
				AlertID:        300,
				RequestBy:      102,
				AssignedUserID: 101,
				Message:        "abc",
				Status:         model.VerificationRequestStatusNew,
				StartTime:      time.Now(),
			},
		},
	}
	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			ctx := context.Background()
			test.WithTxDB(t, func(tx *sql.Tx) {
				// Given
				test.LoadSqlTestFile(t, tx, "testdata/insert_request.sql")

				// Mock
				generateRequestIDFunc = func() (int64, error) {
					return tc.mockGenerateRequestIDFuncGenerator.expID, tc.mockGenerateRequestIDFuncGenerator.expErr
				}
				defer func() { generateRequestIDFunc = generateRequestID }()

				// When
				instance := New(tx)
				result, err := instance.Insert(ctx, tc.givenModel)

				// Then
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.NotNil(t, result.ID)
					require.NotZero(t, result.CreatedAt)
					require.NotZero(t, result.UpdatedAt)
					require.NotZero(t, result.StartTime)
					require.Zero(t, result.EndTime)
					if !cmp.Equal(tc.expResult, result,
						cmpopts.IgnoreFields(model.VerificationRequest{}, "CreatedAt", "UpdatedAt", "StartTime", "EndTime")) {
						t.Errorf("\n VerificationRequest mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expResult, result,
							cmp.Diff(tc.expResult, result, cmpopts.IgnoreFields(model.VerificationRequest{}, "CreatedAt", "UpdatedAt", "StartTime", "EndTime")))
						t.FailNow()
					}
				}
			})
		})
	}
}
