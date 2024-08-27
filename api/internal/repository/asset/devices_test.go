package asset

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetDevices(t *testing.T) {
	tcs := map[string]struct {
		input        GetDevicesInput
		givenFixture string
		expRes       []model.Devices
		expTotal     int64
		expErr       error
	}{
		"success": {
			input: GetDevicesInput{
				PremiseID: 50,
			},
			givenFixture: "testdata/devices.sql",
			expRes: []model.Devices{
				{
					ID:          100,
					PremiseID:   50,
					DeviceName:  "Camera 50",
					DeviceCode:  "cctv_cam50",
					IsActive:    true,
					FloorNumber: 1,
				},
				{
					ID:          101,
					PremiseID:   50,
					DeviceName:  "Camera 51",
					DeviceCode:  "cctv_cam51",
					IsActive:    true,
					FloorNumber: 2,
				},
			},
			expTotal: 2,
		},
		"devices not found": {
			input: GetDevicesInput{
				PremiseID: 50,
			},
		},
		"sucess with limit and page size > 0": {
			input: GetDevicesInput{
				Name:      "Camera",
				PremiseID: 50,
				Limit:     1,
				Page:      1,
			},
			givenFixture: "testdata/devices.sql",
			expRes: []model.Devices{
				{
					ID:          100,
					PremiseID:   50,
					DeviceName:  "Camera 50",
					DeviceCode:  "cctv_cam50",
					IsActive:    true,
					FloorNumber: 1,
				},
			},
			expTotal: 2,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given:
			ctx := context.Background()
			test.WithTxDB(t, func(tx *sql.Tx) {
				defer tx.Rollback()
				if tc.givenFixture != "" {
					test.LoadSqlTestFile(t, tx, tc.givenFixture)
				}
				// When
				repo := New(tx)
				res, total, err := repo.GetDevices(ctx, tc.input)

				// Then:
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.Equal(t, len(tc.expRes), len(res))
					require.Equal(t, tc.expTotal, total)
				}
			})

		})

	}
}
