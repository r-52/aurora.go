package aurora_http

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
)

var Session *session.Store

func CreateSession() {
	storage := sqlite3.New(sqlite3.Config{
		Database: "./db/session.db",
		Table:    "session",
	})
	Session = session.New(session.Config{
		Storage: storage,
	})
}
