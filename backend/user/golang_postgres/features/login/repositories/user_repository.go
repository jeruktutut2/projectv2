package repositories

import (
	"context"
	"golang-postgres/features/login/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	FindByEmail(pool *pgxpool.Pool, ctx context.Context, email string) (user models.User, err error)
}

type UserRepositoryImplementation struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImplementation{}
}

func (repository *UserRepositoryImplementation) FindByEmail(pool *pgxpool.Pool, ctx context.Context, email string) (user models.User, err error) {
	err = pool.QueryRow(ctx, `SELECT id, username, email, password, utc, created_at FROM users WHERE email = $1;`, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Utc, &user.CreatedAt)
	return
}
