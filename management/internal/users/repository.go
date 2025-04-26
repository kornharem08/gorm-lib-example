package users

import (
	"context"

	"github.com/kornharem08/app/internal/model"
	sqlwrap "github.com/kornharem08/gorm"
)

type IRepository interface {
	Find(ctx context.Context) ([]model.User, error)
	FindByEmployeeId(ctx context.Context, employeeId string) (*model.User, error)
}

type repository struct {
	table sqlwrap.ISQL
}

func NewRepository(dbconn sqlwrap.ISQLConnect) IRepository {
	return &repository{table: dbconn.Database().Table(&model.User{})}
}

func (r *repository) Find(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.table.Find(ctx, &users)
	return users, err
}

func (r *repository) FindByEmployeeId(ctx context.Context, employeeId string) (*model.User, error) {
	var user model.User
	var filter = model.User{EmployeeId: employeeId}
	err := r.table.First(ctx, &user, filter)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
