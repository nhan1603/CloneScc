package asset

import (
	"github.com/nhan1603/CloneScc/api/internal/controller/asset"
)

// Handler is the web handler for this pkg
type Handler struct {
	assetCtrl asset.Controller
}

// New instantiates a new Handler and returns it
func New(assetCtrl asset.Controller) Handler {
	return Handler{
		assetCtrl: assetCtrl,
	}
}
