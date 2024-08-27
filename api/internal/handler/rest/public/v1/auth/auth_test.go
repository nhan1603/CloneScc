package auth

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/nhan1603/CloneScc/api/internal/controller/auth"
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_AuthenticateOperationUser(t *testing.T) {
	type mockCheckAuthCtrl struct {
		expCall bool
		input   auth.LoginInput
		output  model.User
		token   string
		err     error
	}

	tcs := map[string]struct {
		givenInput        string
		mockCheckAuthCtrl mockCheckAuthCtrl
		expStatusCode     int
		expRes            string
	}{
		"error when parse json": {
			givenInput:    `{`,
			expRes:        `{"error":"parse_body_failed","error_description":"unexpected end of JSON input"}`,
			expStatusCode: http.StatusBadRequest,
		},
		"error when email invalid": {
			givenInput:    `{}`,
			expRes:        `{"error":"validation_failed", "error_description":"invalid email / password"}`,
			expStatusCode: http.StatusBadRequest,
		},
		"error when call CheckAuth": {
			givenInput: `{"email":"test@gmail.com","password":"1234"}`,
			mockCheckAuthCtrl: mockCheckAuthCtrl{
				expCall: true,
				input: auth.LoginInput{
					Email:    "test@gmail.com",
					Password: "1234",
					Role:     model.UserRoleOperationUser,
				},
				err: errors.New("something went wrong"),
			},
			expRes:        `{"error":"internal_error", "error_description":"Something went wrong"}`,
			expStatusCode: http.StatusInternalServerError,
		},
		"success": {
			givenInput: `{"email":"test@gmail.com","password":"1234"}`,
			mockCheckAuthCtrl: mockCheckAuthCtrl{
				expCall: true,
				input: auth.LoginInput{
					Email:    "test@gmail.com",
					Password: "1234",
					Role:     model.UserRoleOperationUser,
				},
				output: model.User{
					ID:          1,
					DisplayName: "test",
					Email:       "test@gmail.com",
					Password:    "1234",
					Role:        model.UserRoleOperationUser,
				},
				token: "test",
			},
			expRes:        `{"message":"authenticated successfully", "token":"test", "userID":1}`,
			expStatusCode: http.StatusOK,
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

			mockAuthCtrl := new(auth.MockController)
			if tc.mockCheckAuthCtrl.expCall {
				mockAuthCtrl.ExpectedCalls = append(mockAuthCtrl.ExpectedCalls, []*mock.Call{
					mockAuthCtrl.On("CheckAuth", ctx, tc.mockCheckAuthCtrl.input).Return(
						tc.mockCheckAuthCtrl.output,
						tc.mockCheckAuthCtrl.token,
						tc.mockCheckAuthCtrl.err,
					),
				}...)
			}
			// When
			instance := New(mockAuthCtrl)
			handler := instance.AuthenticateOperationUser()
			handler.ServeHTTP(res, req)

			// Then
			require.Equal(t, tc.expStatusCode, res.Code)
			require.JSONEq(t, tc.expRes, res.Body.String())
		})
	}
}

func TestHandler_AuthenticateSecurityGuard(t *testing.T) {
	type mockCheckAuthCtrl struct {
		expCall bool
		input   auth.LoginInput
		output  string
		err     error
	}

	tcs := map[string]struct {
		givenInput        string
		mockCheckAuthCtrl mockCheckAuthCtrl
		expStatusCode     int
		expRes            string
	}{
		"error when parse json": {
			givenInput:    `{`,
			expRes:        `{"error":"parse_body_failed","error_description":"unexpected end of JSON input"}`,
			expStatusCode: http.StatusBadRequest,
		},
		"error when email invalid": {
			givenInput:    `{}`,
			expRes:        `{"error":"validation_failed", "error_description":"invalid email / password"}`,
			expStatusCode: http.StatusBadRequest,
		},
		"error when call CheckAuth": {
			givenInput: `{"email":"test@gmail.com","password":"1234"}`,
			mockCheckAuthCtrl: mockCheckAuthCtrl{
				expCall: true,
				input: auth.LoginInput{
					Email:    "test@gmail.com",
					Password: "1234",
					Role:     model.UserRoleSecurityGuard,
				},
				err: errors.New("something went wrong"),
			},
			expRes:        `{"error":"internal_error", "error_description":"Something went wrong"}`,
			expStatusCode: http.StatusInternalServerError,
		},
		"success": {
			givenInput: `{"email":"test@gmail.com","password":"1234"}`,
			mockCheckAuthCtrl: mockCheckAuthCtrl{
				expCall: true,
				input: auth.LoginInput{
					Email:    "test@gmail.com",
					Password: "1234",
					Role:     model.UserRoleSecurityGuard,
				},
				output: "test",
			},
			expRes:        `{"message":"authenticated successfully","token":"test","userID":0}`,
			expStatusCode: http.StatusOK,
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Setup
			req := httptest.NewRequest(http.MethodPost, "/api/public/v1/auth/sg", strings.NewReader(tc.givenInput))
			req.Header.Set("Content-Type", "application/json")

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext())
			req = req.WithContext(ctx)

			res := httptest.NewRecorder()

			userMock := model.User{}

			mockAuthCtrl := new(auth.MockController)
			if tc.mockCheckAuthCtrl.expCall {
				mockAuthCtrl.ExpectedCalls = append(mockAuthCtrl.ExpectedCalls, []*mock.Call{
					mockAuthCtrl.On("CheckAuth", ctx, tc.mockCheckAuthCtrl.input).Return(userMock, tc.mockCheckAuthCtrl.output, tc.mockCheckAuthCtrl.err),
				}...)
			}
			// When
			instance := New(mockAuthCtrl)
			handler := instance.AuthenticateSecurityGuard()
			handler.ServeHTTP(res, req)

			// Then
			require.Equal(t, tc.expStatusCode, res.Code)
			require.JSONEq(t, tc.expRes, res.Body.String())
		})
	}
}
