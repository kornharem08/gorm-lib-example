package users

import (
	"context"

	"github.com/kornharem08/app/internal/model"
	sqlwrap "github.com/kornharem08/gorm"
)

type IService interface {
	Find(ctx context.Context) ([]model.User, error)
}

type Service struct {
	Repository IRepository
}

func NewService(dbconn sqlwrap.ISQLConnect) IService {
	return &Service{Repository: NewRepository(dbconn)}
}

func (s *Service) Find(ctx context.Context) ([]model.User, error) {
	return s.Repository.Find(ctx)
}
