package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
)

type userResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

type getUsersResponse struct {
	Items []userResponse `json:"items"`
}

// GetUsers retrieves a list of users
func (h Handler) GetUsers() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[GetUsers] START processing requests")
		users, err := h.userCtrl.GetUsers(r.Context())
		if err != nil {
			return err
		}

		items := make([]userResponse, len(users))
		for idx, user := range users {
			items[idx] = userResponse{
				ID:    strconv.FormatInt(user.ID, 10),
				Email: user.Email,
				Name:  user.DisplayName,
				Role:  user.Role.String(),
			}
		}

		httpserver.RespondJSON(w, getUsersResponse{
			Items: items,
		})
		return nil
	})
}
