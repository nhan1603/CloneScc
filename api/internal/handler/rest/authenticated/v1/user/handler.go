package user

import (
	"github.com/nhan1603/CloneScc/api/internal/controller/users"
)

// Handler is the web handler for this pkg
type Handler struct {
	userCtrl users.Controller
}

// New instantiates a new Handler and returns it
func New(userCtrl users.Controller) Handler {
	return Handler{userCtrl: userCtrl}
}
