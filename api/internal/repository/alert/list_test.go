package alert

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestRepository_impl_GetAlerts(t *testing.T) {
	tcs := map[string]struct {
		input  model.GetAlertsInput
		expRes []model.Alert
		expErr error
	}{
		"success|get alerts with premise": {
			input: model.GetAlertsInput{
				PremiseID: 100,
				Limit:     50,
				Page:      1,
			},
			expRes: []model.Alert{
				{
					ID:              300,
					CCTVDeviceID:    200,
					Type:            "Suspicious Activities",
					Description:     "TEST",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					CCTVDevice:      "CCTV 1",
					CCTVDeviceFloor: 1,
					IsAcknowledged:  true,
				},
				{
					ID:              301,
					CCTVDeviceID:    200,
					Type:            "Unauthorized Access",
					Description:     "TEST",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					CCTVDevice:      "CCTV 1",
					CCTVDeviceFloor: 1,
					IsAcknowledged:  true,
				},
				{
					ID:              302,
					CCTVDeviceID:    200,
					Type:            "Suspicious Activities",
					Description:     "TEST",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					CCTVDevice:      "CCTV 1",
					CCTVDeviceFloor: 1,
					IsAcknowledged:  true,
				},
			},
		},
		"success|limit 2, page 2": {
			input: model.GetAlertsInput{
				PremiseID: 100,
				Limit:     2,
				Page:      2,
			},
			expRes: []model.Alert{
				{
					ID:              302,
					CCTVDeviceID:    200,
					Type:            "Suspicious Activities",
					Description:     "TEST",
					PremiseName:     "Sunrise Tower",
					PremiseLocation: "307/12 Nguyen Van Troi St, W1, Tan Binh",
					CCTVDevice:      "CCTV 1",
					CCTVDeviceFloor: 1,
					IsAcknowledged:  true,
				},
			},
		},
		"not found": {
			input: model.GetAlertsInput{
				PremiseID: 200,
				Limit:     50,
				Page:      1,
			},
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			test.WithTxDB(t, func(tx *sql.Tx) {
				// Given:
				ctx := context.Background()
				test.LoadSqlTestFile(t, tx, "testdata/alerts.sql")

				// When:
				repo := New(tx)
				_, _, err := repo.GetAlerts(ctx, tc.input)

				// Then:
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
				}
			})
		})
	}
}
