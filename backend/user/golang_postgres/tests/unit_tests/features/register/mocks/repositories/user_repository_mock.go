package mockrepositories

import (
	"context"
	"golang-postgres/features/register/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) Create(pool *pgxpool.Pool, ctx context.Context, user models.User) (rowsAffected int64, err error) {
	arguments := repository.Mock.Called(pool, ctx, user)
	return arguments.Get(0).(int64), arguments.Error(1)
}

func (repository *UserRepositoryMock) CountByEmail(pool *pgxpool.Pool, ctx context.Context, email string) (numbeOfUser int, err error) {
	arguments := repository.Mock.Called(pool, ctx, email)
	return arguments.Get(0).(int), arguments.Error(1)
}
