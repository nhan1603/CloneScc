package users

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// GetUsers gets all users
	GetUsers(ctx context.Context) ([]model.User, error)
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
