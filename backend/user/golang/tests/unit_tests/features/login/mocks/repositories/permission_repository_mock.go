package mockrepositories

import (
	"context"
	"database/sql"
	"project-user/features/login/models"

	"github.com/stretchr/testify/mock"
)

type PermissionRepositoryMock struct {
	Mock mock.Mock
}

func (repository *PermissionRepositoryMock) FindByInId(db *sql.DB, ctx context.Context, ids []interface{}) (permissions []models.Permission, err error) {
	arguments := repository.Mock.Called(db, ctx, ids)
	return arguments.Get(0).([]models.Permission), arguments.Error(1)
}
