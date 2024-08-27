package auth

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// CheckAuth handles authentication checking
	CheckAuth(ctx context.Context, inp LoginInput) (model.User, string, error)
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
