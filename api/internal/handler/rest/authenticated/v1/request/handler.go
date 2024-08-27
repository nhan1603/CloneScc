package request

import (
	"github.com/nhan1603/CloneScc/api/internal/controller/requests"
)

// Handler is the web handler for this pkg
type Handler struct {
	requestCtrl requests.Controller
}

// New instantiates a new Handler and returns it
func New(requestCtrl requests.Controller) Handler {
	return Handler{
		requestCtrl: requestCtrl,
	}
}
