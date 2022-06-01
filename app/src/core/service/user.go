package service

import (
	"app/src/api/http/schema"
	"app/src/core/model"
	"app/src/core/repo"
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	log  echo.Logger
	repo *repo.UserRepository
}

func NewUserService(repo *repo.UserRepository, log echo.Logger) *UserService {
	return &UserService{
		log:  log,
		repo: repo,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user *schema.UserCreateRequest) (*schema.UserCreateResponse, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		us.log.Error(err)
		return nil, err
	}

	res, err := us.repo.Insert(ctx, &model.User{Email: user.Email, Username: user.Username, Password: string(pwd)})
	if err != nil {
		us.log.Error(err)
		return nil, err
	}
	return schema.NewUserCreateResponse(res.ID, res.Username, nil), nil
}

func (us *UserService) GetUser(ctx context.Context, userID uint64) (*schema.UserSimpleResponse, error) {
	res, err := us.repo.SelectWhereId(ctx, userID)
	if err != nil {
		us.log.Error(err)
		return nil, err
	}

	return schema.NewUserSimpleResponse(res.ID, res.Username, nil, schema.UserOffline), nil

}
