package dep

import (
	"app/src/core/cfg"
	"app/src/core/repo"
	"app/src/core/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Dep = &container{
	db: initDB(),
}

type container struct {
	db *sqlx.DB
}

func (c *container) userRepository() *repo.UserRepository {
	return repo.NewUserRepository(c.db)
}

func (c *container) chatRepository() *repo.ChatRepository {
	return repo.NewChatRepository(c.db)
}

func (c *container) UserService() *service.UserService {
	return service.NewUserService(c.userRepository())
}

func (c *container) ChatService() *service.ChatService {
	return service.NewChatService(c.chatRepository())
}

func initDB() *sqlx.DB {
	return sqlx.MustOpen("postgres", cfg.Config.DBConnectionURL)
}
