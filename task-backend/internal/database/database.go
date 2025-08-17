package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Mahathirrr/task-management-backend/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase(cfg *config.DatabaseConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	
	log.Print("Database connected successfully")
	return nil
}

// GetDB mengembalikan instance database
func GetDB() *sql.DB {
	return DB
}