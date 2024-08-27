package asset

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetPremises(t *testing.T) {
	tcs := map[string]struct {
		input        GetPremisesInput
		givenFixture string
		expRes       []model.Premises
		expErr       error
	}{
		"success": {
			input: GetPremisesInput{
				Name: "MyTower",
			},
			givenFixture: "testdata/premises.sql",
			expRes: []model.Premises{
				{
					ID:           50,
					Name:         "Sunrise Tower",
					Location:     "307/12 Nguyen Van Troi St, W1, Tan Binh",
					PremisesCode: "P001",
					Description:  "Sunrise Tower",
					CCTVCount:    4,
				},
				{
					ID:           51,
					Name:         "Bitexco Financial Tower",
					Location:     "2 Hai Ba Trung St, Ben Nghe, District 1",
					PremisesCode: "P002",
					Description:  "Bitexco Financial Tower",
					CCTVCount:    4,
				},
			},
		},
		"Premises not found": {
			input: GetPremisesInput{
				Name: "MyTower",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given:
			ctx := context.Background()
			test.WithTxDB(t, func(tx *sql.Tx) {
				defer tx.Rollback()
				// Setup
				if tc.givenFixture != "" {
					test.LoadSqlTestFile(t, tx, tc.givenFixture)
				}
				repo := New(tx)

				// When:
				res, err := repo.GetPremises(ctx, tc.input)

				// Then:
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.Equal(t, len(tc.expRes), len(res))
				}
			})
		})

	}
}
