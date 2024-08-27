package request

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestRepository_impl_GetRequest(t *testing.T) {
	tcs := map[string]struct {
		requestID int64
		expRes    model.RequestDetail
		expErr    error
	}{
		"success": {
			requestID: 400,
			expRes: model.RequestDetail{
				ID:       400,
				Title:    "CCTV 1-300",
				Author:   "John",
				Assignee: "Thomas",
				Message:  "Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.",
				AlertDetail: model.Alert{
					ID:              300,
					Type:            "Suspicious Activities",
					Description:     "TEST",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					CCTVDevice:      "CCTV 1",
					CCTVDeviceFloor: 1,
					IsAcknowledged:  true,
				},
				Respond: &model.RequestRespond{
					ID:      500,
					User:    "Thomas",
					Message: "Ive checked the monitored area, and it's a false alarm. No property damage or suspicious activities were found.",
				},
			},
		},
		"not found": {
			requestID: 1,
			expErr:    ErrNotFound,
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
				rs, err := repo.GetRequest(ctx, tc.requestID)

				// Then:
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					test.Compare(t, tc.expRes, rs, model.RequestDetail{}, "StartTime", "AlertDetail", "Respond")
					test.Compare(t, tc.expRes.AlertDetail, rs.AlertDetail, model.Alert{}, "CCTVDeviceID", "IncidentAt", "CreatedAt", "UpdatedAt")
					test.Compare(t, tc.expRes.Respond, rs.Respond, model.RequestRespond{}, "MediaData", "VerifiedAt")
				}
			})
		})
	}
}
