package request

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestRepository_UpdateStatusAndEndTime(t *testing.T) {
	tcs := map[string]struct {
		givenInputReqID int64
		givenStatus     model.VerificationRequestStatus
		givenEndTime    time.Time
		expRes          model.VerificationRequest
		err             error
	}{
		"success": {
			givenInputReqID: 400,
			givenStatus:     model.VerificationRequestStatusResolved,
			givenEndTime:    time.Now(),
			expRes: model.VerificationRequest{
				ID:             400,
				AlertID:        300,
				RequestBy:      100,
				AssignedUserID: 101,
				Message:        "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
				Status:         model.VerificationRequestStatusResolved,
			},
		},
		"error when no rows": {
			givenInputReqID: 1,
			err:             ErrNotFound,
		},
		"error when update empty status": {
			givenInputReqID: 400,
			givenStatus:     "",
			err:             errors.New(`dbmodel: unable to update verification_requests row: pq: new row for relation "verification_requests" violates check constraint "verification_requests_status_check"`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given:
			ctx := context.Background()

			// Setup
			test.WithTxDB(t, func(tx *sql.Tx) {
				test.LoadSqlTestFile(t, tx, "testdata/requests.sql")

				// When:
				repo := New(tx)
				rs, err := repo.UpdateStatusAndEndTime(ctx, tc.givenInputReqID, tc.givenStatus, tc.givenEndTime)

				// Then:
				if tc.err != nil {
					require.EqualError(t, err, tc.err.Error())
				} else {
					test.Compare(t, tc.expRes, rs, model.VerificationRequest{}, "CreatedAt", "UpdatedAt", "StartTime", "EndTime")
				}
			})
		})
	}
}
