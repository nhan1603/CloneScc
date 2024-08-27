package users

import (
	"context"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository"
	"github.com/nhan1603/CloneScc/api/internal/repository/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_GetUsers(t *testing.T) {
	type mockGetUsersRepo struct {
		expCall bool
		output  []model.User
		err     error
	}

	tcs := map[string]struct {
		mockGetUsersRepo mockGetUsersRepo
		expRes           []model.User
		expErr           error
	}{
		"success": {
			mockGetUsersRepo: mockGetUsersRepo{
				expCall: true,
				output: []model.User{
					{
						ID:          1,
						Email:       "test@gmail.com",
						DisplayName: "test",
					},
				},
			},
			expRes: []model.User{
				{
					ID:          1,
					Email:       "test@gmail.com",
					DisplayName: "test",
				},
			},
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			ctx := context.Background()
			repo := new(repository.MockRegistry)
			mockUserRepo := new(user.MockRepository)

			repo.ExpectedCalls = []*mock.Call{
				repo.On("User").Return(mockUserRepo),
			}

			// When
			if tc.mockGetUsersRepo.expCall {
				mockUserRepo.ExpectedCalls = []*mock.Call{
					mockUserRepo.On("GetAll", ctx).Return(tc.mockGetUsersRepo.output, tc.mockGetUsersRepo.err),
				}
			}

			instance := New(repo)
			rs, err := instance.GetUsers(ctx)

			// Then
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRes, rs)
			}
		})
	}
}
