package request

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetByID(t *testing.T) {
	tcs := map[string]struct {
		givenInput int64
		expRes     model.VerificationRequest
		err        error
	}{
		"success": {
			givenInput: 400,
			expRes: model.VerificationRequest{
				ID:             400,
				AlertID:        300,
				RequestBy:      100,
				AssignedUserID: 101,
				Message:        "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
				Status:         "NEW",
			},
		},
		"error when no rows": {
			givenInput: 1,
			err:        ErrNotFound,
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
				rs, err := repo.GetByID(ctx, tc.givenInput)

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
