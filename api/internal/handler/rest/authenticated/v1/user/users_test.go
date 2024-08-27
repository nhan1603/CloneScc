package user

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/nhan1603/CloneScc/api/internal/controller/users"
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetUsers(t *testing.T) {
	type mockGetUserCtrl struct {
		expCall bool
		output  []model.User
		err     error
	}

	tcs := map[string]struct {
		mockGetUserCtrl mockGetUserCtrl
		givenInput      string
		expStatusCD     int
		expRes          string
	}{
		"success": {
			mockGetUserCtrl: mockGetUserCtrl{
				expCall: true,
				output: []model.User{
					{
						ID:          1,
						Email:       "test@gmail.com",
						DisplayName: "test",
					},
				},
			},
			expStatusCD: http.StatusOK,
			expRes:      `{"items":[{"id":"1","email":"test@gmail.com","name":"test","role":""}]}`,
		},
		"error when get users": {
			mockGetUserCtrl: mockGetUserCtrl{
				expCall: true,
				err:     errors.New("something went wrong"),
			},
			expRes:      `{"error":"internal_error","error_description":"Something went wrong"}`,
			expStatusCD: http.StatusInternalServerError,
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Setup
			req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(tc.givenInput))
			req.Header.Set("Content-Type", "application/json")

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext())
			req = req.WithContext(ctx)

			res := httptest.NewRecorder()

			mockUserCtrl := new(users.MockController)
			if tc.mockGetUserCtrl.expCall {
				mockUserCtrl.ExpectedCalls = append(mockUserCtrl.ExpectedCalls, []*mock.Call{
					mockUserCtrl.On("GetUsers", ctx).Return(tc.mockGetUserCtrl.output, tc.mockGetUserCtrl.err),
				}...)
			}

			// When
			instance := New(mockUserCtrl)
			handler := instance.GetUsers()
			handler.ServeHTTP(res, req)

			// Then
			require.Equal(t, tc.expStatusCD, res.Code)
			require.JSONEq(t, tc.expRes, res.Body.String())
		})
	}
}
