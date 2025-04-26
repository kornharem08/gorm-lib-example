package users

import (
	"context"

	"github.com/kornharem08/app/internal/model"
	sqlwrap "github.com/kornharem08/gorm"
)

type IRepository interface {
	Find(ctx context.Context) ([]model.User, error)
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
