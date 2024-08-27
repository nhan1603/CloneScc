package alert

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestRepository_impl_GetAlert(t *testing.T) {
	tcs := map[string]struct {
		alertID int64
		expRes  model.Alert
		expErr  error
	}{
		"success": {
			alertID: 300,
			expRes: model.Alert{
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
		},
		"not found": {
			alertID: 1,
			expErr:  ErrNotFound,
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			test.WithTxDB(t, func(tx *sql.Tx) {
				// Given:
				ctx := context.Background()
				// Setup
				test.LoadSqlTestFile(t, tx, "testdata/alerts.sql")

				// When:
				repo := New(tx)
				rs, err := repo.GetAlert(ctx, tc.alertID)

				// Then:
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					test.Compare(t, tc.expRes, rs, model.Alert{}, "IncidentAt", "CreatedAt", "UpdatedAt")
				}
			})
		})
	}
}
