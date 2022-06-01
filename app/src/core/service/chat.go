package service

import (
	"context"

	"app/src/core/model"
	"app/src/core/repo"
)

type ChatService struct {
	repo *repo.ChatRepository
}

func NewChatService(repo *repo.ChatRepository) *ChatService {
	return &ChatService{
		repo: repo,
	}
}

func (s *ChatService) GetChat(ctx context.Context, id uint64) (*model.Chat, error) {
	return s.repo.SelectWhereId(ctx, id)
}
