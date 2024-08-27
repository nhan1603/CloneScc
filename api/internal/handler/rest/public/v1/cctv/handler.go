package cctv

import (
	"github.com/nhan1603/CloneScc/api/internal/controller/asset"
)

// Handler is the web handler for this pkg
type Handler struct {
	asset asset.Controller
}

// New instantiates a new Handler and returns it
func New(asset asset.Controller) Handler {
	return Handler{asset: asset}
}
