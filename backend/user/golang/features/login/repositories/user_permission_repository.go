package repositories

import (
	"context"
	"database/sql"
	"project-user/features/login/models"
)

type UserPermissionRepository interface {
	FindByUserId(db *sql.DB, ctx context.Context, userId int32) (userPermissions []models.UserPermission, err error)
}

type UserPermissionRepositoryImplementation struct {
}

func NewUserPermissinoRepository() UserPermissionRepository {
	return &UserPermissionRepositoryImplementation{}
}

func (repository *UserPermissionRepositoryImplementation) FindByUserId(db *sql.DB, ctx context.Context, userId int32) (userPermissions []models.UserPermission, err error) {
	rows, err := db.QueryContext(ctx, `SELECT id, user_id, permission_id FROM user_permission WHERE user_id = ?;`, userId)
	if err != nil {
		return
	}
	defer func() {
		errRowsClose := rows.Close()
		if errRowsClose != nil {
			userPermissions = []models.UserPermission{}
			err = errRowsClose
		}
	}()

	for rows.Next() {
		var userPermission models.UserPermission
		err = rows.Scan(&userPermission.Id, &userPermission.UserId, &userPermission.PermissionId)
		if err != nil {
			userPermissions = []models.UserPermission{}
			return
		}
		userPermissions = append(userPermissions, userPermission)
	}
	return
}
