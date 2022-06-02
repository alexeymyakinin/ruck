package dep

import (
	"github.com/alexeymyakinin/ruck/app/src/core/env"
	"github.com/alexeymyakinin/ruck/app/src/core/repo"
	"github.com/alexeymyakinin/ruck/app/src/core/service"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

var (
	db = sqlx.MustOpen("postgres", env.DBConnectionURL)
)

func logger(prefix string) echo.Logger {
	return log.New(prefix)
}

func userRepository() *repo.UserRepository {
	return repo.NewUserRepository(db, logger("repo:user"))
}

func chatRepository() *repo.ChatRepository {
	return repo.NewChatRepository(db, logger("repo:chat"))
}

func GetUserService() *service.UserService {
	return service.NewUserService(userRepository(), logger("service:user"))
}

func GetChatService() *service.ChatService {
	return service.NewChatService(chatRepository(), logger("service:user"))
}
