package user

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetByCriteria(t *testing.T) {
	tcs := map[string]struct {
		givenInput GetUserInput
		expRes     model.User
		err        error
	}{
		"error when sql no rows": {
			givenInput: GetUserInput{
				Email: "test",
			},
			err: ErrNotFound,
		},
		"success": {
			givenInput: GetUserInput{
				Email: "john@scc.com",
				Role:  model.UserRoleOperationUser,
			},
			expRes: model.User{
				ID:          100,
				DisplayName: "John",
				Email:       "john@scc.com",
				Password:    "test",
				Role:        "OPERATION_USER",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given:
			ctx := context.Background()

			// Setup
			test.WithTxDB(t, func(tx *sql.Tx) {
				test.LoadSqlTestFile(t, tx, "testdata/get_all.sql")

				// When:
				repo := New(tx)
				rs, err := repo.GetByCriteria(ctx, tc.givenInput)

				// Then:
				if tc.err != nil {
					require.EqualError(t, err, tc.err.Error())
				} else {
					test.Compare(t, tc.expRes, rs, model.User{}, "CreatedAt", "UpdatedAt")
				}
			})
		})
	}
}
