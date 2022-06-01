package service

import (
	"app/src/api/http/schema"
	"app/src/core/model"
	"app/src/core/repo"
	"context"
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

func (cs *ChatService) CreateChat(ctx context.Context, chat *schema.ChatCreateRequest) (*schema.ChatCreateResponse, error) {
	res, err := cs.repo.Insert(ctx, &model.Chat{Name: chat.Name})
	if err != nil {
		cs.log.Error(err)
		return nil, err
	}

	return schema.NewChatCreateResponse(res.ID, res.Name), nil
}

func (cs *ChatService) GetChat(ctx context.Context, chatId uint64) (*schema.ChatSimpleResponse, error) {
	res, err := cs.repo.SelectWhereId(ctx, chatId)
	if err != nil {
		cs.log.Error(err)
		return nil, err
	}

	return schema.NewChatSimpleResponse(res.ID, res.Name), nil

}
