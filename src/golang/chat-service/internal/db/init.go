package db

import (
	"chat-service/configs"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	Ctx  = context.Background()
	Db   *sql.DB
	conf = configs.GetDbConfig()
)

func Init() {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.Username,
		conf.Password,
		conf.Database,
	)

	var err error
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Failed to connect to the database:", err)
	}

	if err = Db.Ping(); err != nil {
		log.Fatal("❌ Failed to ping the database:", err)
	}

	log.Println("✅ Connected to PostgreSQL successfully")
}
