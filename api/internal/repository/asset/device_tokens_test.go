package asset

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/nhan1603/CloneScc/api/internal/repository/generator"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetDeviceToken(t *testing.T) {
	tcs := map[string]struct {
		givenUserID int64
		expToken    string
		expErr      error
	}{
		"success": {
			givenUserID: 250,
			expToken:    "qUsCeUjQ5Z6010",
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given:
			ctx := context.Background()
			test.WithTxDB(t, func(tx *sql.Tx) {
				// Setup:
				test.LoadSqlTestFile(t, tx, "testdata/device_tokens.sql")
				repo := New(tx)

				// When:
				deviceToken, err := repo.GetDeviceToken(ctx, tc.givenUserID)

				// Then:
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.expToken, deviceToken)
				}
			})
		})
	}
}

func TestRepository_UpsertDeviceToken(t *testing.T) {
	tcs := map[string]struct {
		givenInput UpsertDeviceTokenInput
		expErr     error
	}{
		"success": {
			givenInput: UpsertDeviceTokenInput{
				UserID:      252,
				DeviceToken: "ABCD",
				Platform:    "ios",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given:
			ctx := context.Background()
			test.WithTxDB(t, func(tx *sql.Tx) {
				require.Nil(t, generator.InitSnowflakeGenerators())
				// Setup
				test.LoadSqlTestFile(t, tx, "testdata/device_tokens.sql")
				repo := New(tx)

				// When
				gotErr := repo.UpsertDeviceToken(ctx, tc.givenInput)

				// Then
				require.NoError(t, gotErr)
			})
		})
	}
}
