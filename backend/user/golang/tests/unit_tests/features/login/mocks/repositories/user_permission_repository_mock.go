package mockrepositories

import (
	"context"
	"database/sql"

	"project-user/features/login/models"

	"github.com/stretchr/testify/mock"
)

type UserPermissionRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserPermissionRepositoryMock) FindByUserId(db *sql.DB, ctx context.Context, userId int32) (userPermissions []models.UserPermission, err error) {
	arguments := repository.Mock.Called(db, ctx, userId)
	return arguments.Get(0).([]models.UserPermission), arguments.Error(1)
}
