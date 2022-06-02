package service

import (
	"context"
	"github.com/alexeymyakinin/ruck/app/src/api/http/schema"
	"github.com/alexeymyakinin/ruck/app/src/core/model"
	"github.com/alexeymyakinin/ruck/app/src/core/repo"
	"github.com/labstack/echo/v4"
)

type ChatService struct {
	log  echo.Logger
	repo *repo.ChatRepository
}

func NewChatService(repo *repo.ChatRepository, log echo.Logger) *ChatService {
	return &ChatService{
		log:  log,
		repo: repo,
	}
}

func (cs *ChatService) CreateChat(ctx context.Context, chat *schema.ChatCreateRequest) (*model.Chat, error) {
	res, err := cs.repo.Insert(ctx, &model.Chat{Name: chat.Name})
	if err != nil {
		cs.log.Error(err)
		return nil, err
	}

	return res, nil
}

func (cs *ChatService) GetChat(ctx context.Context, chatId uint64) (*model.Chat, error) {
	res, err := cs.repo.SelectWhereId(ctx, chatId)
	if err != nil {
		cs.log.Error(err)
		return nil, err
	}

	return res, nil

}
