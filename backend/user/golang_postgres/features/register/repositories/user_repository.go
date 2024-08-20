package repositories

import (
	"context"
	"golang-postgres/features/register/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(pool *pgxpool.Pool, ctx context.Context, user models.User) (rowsAffected int64, err error)
	CountByEmail(pool *pgxpool.Pool, ctx context.Context, email string) (numbeOfUser int, err error)
}

type UserRepositoryImplementation struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImplementation{}
}

func (repository *UserRepositoryImplementation) Create(pool *pgxpool.Pool, ctx context.Context, user models.User) (rowsAffected int64, err error) {
	result, err := pool.Exec(ctx, `INSERT INTO users (username, email, password, utc, created_at) VALUES($1, $2, $3, $4, $5);`, user.Username, user.Email, user.Password, user.Utc, user.CreatedAt)
	if err != nil {
		return
	}
	rowsAffected = result.RowsAffected()
	return
}

func (repository *UserRepositoryImplementation) CountByEmail(pool *pgxpool.Pool, ctx context.Context, email string) (numbeOfUser int, err error) {
	err = pool.QueryRow(ctx, `SELECT COUNT(*) AS number_of_user FROM users WHERE email = $1;`, email).Scan(&numbeOfUser)
	return
}
