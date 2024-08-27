package asset

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetCCTVKeyByName(t *testing.T) {
	tcs := map[string]struct {
		givenFixture string
		cctvName     string
		expRes       int64
		expErr       error
	}{
		"success": {
			givenFixture: "testdata/devices.sql",
			cctvName:     "Camera 50",
			expRes:       100,
		},
		"cttv not found": {
			cctvName: "Camera 50",
			expErr:   ErrCctvNotFound,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given:
			ctx := context.Background()
			test.WithTxDB(t, func(tx *sql.Tx) {
				// Setup:
				if tc.givenFixture != "" {
					test.LoadSqlTestFile(t, tx, tc.givenFixture)
				}
				// When:
				repo := New(tx)
				id, err := repo.GetCCTVKeyByName(ctx, tc.cctvName)

				// Then:
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.expRes, id)
				}
			})
		})
	}
}
func TestRepository_GetAllCCTV(t *testing.T) {
	tcs := map[string]struct {
		cctvName string
		expRes   int
		expErr   error
	}{
		"success": {
			expRes: 25,
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given:
			ctx := context.Background()
			test.WithTxDB(t, func(tx *sql.Tx) {
				// Setup:
				test.LoadSqlTestFile(t, tx, "testdata/devices.sql")
				// When:
				repo := New(tx)
				actRes, err := repo.GetAllCCTV(ctx)

				// Then:
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.expRes, len(actRes))
				}
			})
		})
	}
}
