package alert

import (
	"context"
	"testing"
	"time"

	"database/sql"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/nhan1603/CloneScc/api/internal/repository/generator"
	"github.com/stretchr/testify/require"
)

func TestRepository_impl_CreateAlert(t *testing.T) {
	tcs := map[string]struct {
		cctvDeviceID int64
		dataMessage  model.AlertMessage
		expErr       error
	}{
		"success": {
			cctvDeviceID: 200,
			dataMessage: model.AlertMessage{
				ID:          "30",
				CCTVName:    "CCTV 1",
				FloorNumber: "1",
				Type:        model.AlertTypePropertyDamage.ToString(),
				Description: "Test",
				IncidentAt:  time.Date(2023, 8, 4, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			test.WithTxDB(t, func(tx *sql.Tx) {
				// Given:
				ctx := context.Background()
				require.Nil(t, generator.InitSnowflakeGenerators())

				// Setup
				test.LoadSqlTestFile(t, tx, "testdata/alerts.sql")

				// When:
				repo := New(tx)
				_, err := repo.CreateAlert(ctx, tc.dataMessage, tc.cctvDeviceID)

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
