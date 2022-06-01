package service

import (
	"context"

	"app/src/api/http/schema"
	"app/src/core/model"
	"app/src/core/repo"
	logger "github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repo.UserRepository
}

func NewUserService(repo *repo.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *schema.UserCreateRequest) (
	*schema.UserCreateResponse, error,
) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error().Err(err).Msg("UserService:CreateUser cannot get password hash")
		return nil, err
	}

	res, err := s.repo.Insert(ctx, &model.User{Email: user.Email, Username: user.Username, Password: string(pwd)})
	if err != nil {
		logger.Error().Err(err).Msg("UserService:CreateUser cannot create user")
		return nil, err
	}

	return &schema.UserCreateResponse{User: schema.NewUserDetailResponse(res.ID, res.Username, nil)}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, userID uint64) (*schema.UserSimpleResponse, error) {
	usr, err := s.repo.SelectWhereId(ctx, userID)
	if err != nil {
		logger.Error().Err(err).Msg("UserService:GetUserById cannot get user")
		return nil, err
	}

	return schema.NewUserSimpleResponse(usr.ID, usr.Username, nil), nil

}
