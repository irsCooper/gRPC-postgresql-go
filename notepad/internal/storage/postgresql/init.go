package postgresql

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToPostgresql(log *slog.Logger) {
	const op = "postgresql.init.ConnectToPostgresql"
	var err error
	dsn := "host=localhost user=crowley password=sheldon dbname=notepadproject port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Info("connect to postgres", slog.String("error", err.Error()))
	}
}
