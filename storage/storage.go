package storage

import (
	"context"

	"github.com/user/api/models"
)

type StorageI interface {
	Close()
	User() UserRepoI
	PhoneNumber() PhoneNumberRepoI
}

type UserRepoI interface {
	Create(context.Context, *models.CreateUser) (string, error)
	GetByID(context.Context, *models.UserPrimaryKey) (*models.User, error)
	GetList(context.Context, *models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(context.Context, *models.UpdateUser) (int64, error)
	Delete(context.Context, *models.UserPrimaryKey) error
}
type PhoneNumberRepoI interface {
	Create(ctx context.Context, req *models.CreatePhoneNumber) (string, error)
	GetByID(ctx context.Context, req *models.PhoneNumberPrimaryKey) (*models.PhoneNumber, error)
	GetList(ctx context.Context, req *models.PhoneNumberGetListRequest) (*models.PhoneNumberGetListResponse, error)
	Update(ctx context.Context, req *models.UpdatePhoneNumber) (int64, error)
	Delete(ctx context.Context, req *models.PhoneNumberPrimaryKey) error
}
