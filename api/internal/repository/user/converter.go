package user

import (
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/dbmodel"
)

func toModelUsers(orms dbmodel.UserSlice) []model.User {
	users := make([]model.User, len(orms))
	for i, o := range orms {
		users[i] = toModelUser(o)
	}
	return users
}

func toModelUser(user *dbmodel.User) model.User {
	return model.User{
		ID:          user.ID,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		Password:    user.Password,
		Role:        model.UserRole(user.UserRole),
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
