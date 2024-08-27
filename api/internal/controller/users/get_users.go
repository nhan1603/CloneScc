package users

import (
	"context"

	"github.com/nhan1603/CloneScc/api/internal/model"
)

// GetUsers gets all users
func (i impl) GetUsers(ctx context.Context) ([]model.User, error) {
	return i.repo.User().GetAll(ctx)
}
