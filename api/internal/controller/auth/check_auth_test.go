package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository"
	"github.com/nhan1603/CloneScc/api/internal/repository/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_CheckAuth(t *testing.T) {
	type mockCheckAuthFn struct {
		expCall bool
		input   LoginInput
		output  model.User
		err     error
	}

	tcs := map[string]struct {
		givenInput      LoginInput
		mockCheckAuthFn mockCheckAuthFn
		expResToken     string
		expErr          error
	}{
		"error when call check auth fnc": {
			mockCheckAuthFn: mockCheckAuthFn{
				expCall: true,
				err:     errors.New("something went wrong"),
			},
			expErr: errors.New("something went wrong"),
		},
		"success": {
			mockCheckAuthFn: mockCheckAuthFn{
				expCall: true,
				output: model.User{
					ID:    1,
					Email: "test@gmail.com",
					Role:  "test",
				},
			},
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			ctx := context.Background()
			defer func() {
				checkAuthFn = checkAuth
			}()

			// When
			if tc.mockCheckAuthFn.expCall {
				checkAuthFn = func(ctx context.Context, i impl, inp LoginInput) (model.User, error) {
					require.Equal(t, inp, tc.mockCheckAuthFn.input)
					return tc.mockCheckAuthFn.output, tc.mockCheckAuthFn.err
				}
			}

			instance := New(nil)
			_, rsToken, err := instance.CheckAuth(ctx, tc.givenInput)

			// Then
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, len(rsToken) > 0, true)
			}
		})
	}
}

func TestController_checkAuth(t *testing.T) {
	type mockGetByCriteriaRepo struct {
		expCall bool
		input   user.GetUserInput
		output  model.User
		err     error
	}

	tcs := map[string]struct {
		givenInput            LoginInput
		mockGetByCriteriaRepo mockGetByCriteriaRepo
		expRes                model.User
		expErr                error
	}{
		"error when call GetByCriteria func": {
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				err:     errors.New("something went wrong"),
			},
			expErr: errors.New("something went wrong"),
		},
		"error when call GetByCriteria func with user not found": {
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				err:     user.ErrNotFound,
			},
			expErr: ErrUserNotFound,
		},
		"error when CompareHashAndPassword": {
			givenInput: LoginInput{
				Password: "input pass",
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				output: model.User{
					Password: "test",
				},
			},
			expErr: ErrUserNotFound,
		},
		"success": {
			givenInput: LoginInput{
				Password: "password",
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				output: model.User{
					Password: "$2a$10$3my3mwxVp3.P82E.EXt8r.4n4ShsRt..k3Xe7XzlWb83VX4mJhAwi",
				},
			},
			expRes: model.User{
				Password: "$2a$10$3my3mwxVp3.P82E.EXt8r.4n4ShsRt..k3Xe7XzlWb83VX4mJhAwi",
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
			if tc.mockGetByCriteriaRepo.expCall {
				mockUserRepo.ExpectedCalls = []*mock.Call{
					mockUserRepo.On("GetByCriteria", ctx, tc.mockGetByCriteriaRepo.input).Return(tc.mockGetByCriteriaRepo.output, tc.mockGetByCriteriaRepo.err),
				}
			}

			// When
			rs, err := checkAuth(ctx, impl{repo}, tc.givenInput)

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
