package alert

import "github.com/nhan1603/CloneScc/api/internal/controller/alerts"

// Handler is the web handler for this pkg
type Handler struct {
	alertCtrl alerts.Controller
}

// New instantiates a new Handler and returns it
func New(alertCtrl alerts.Controller) Handler {
	return Handler{
		alertCtrl: alertCtrl,
	}
}
