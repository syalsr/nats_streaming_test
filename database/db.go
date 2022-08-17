package database

import (
	"Wildberries_L0/config"
	"Wildberries_L0/detail"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func GetUID(connection *pgx.Conn, OrderUID string) *detail.OrderInfo {
	query := `
		select info from orders where order_uid=$1
	`
	info := new(detail.OrderInfo)
	connection.QueryRow(context.Background(), query, OrderUID).Scan(&info)
	return info
}

func InsertData(connection *pgx.Conn, info detail.OrderInfo) {
	query := `
		INSERT INTO orders (info, order_uid)
		VALUES($1, $2)
	`
	connection.QueryRow(context.Background(), query, info, info.OrderUID)
}

func Connect() *pgx.Conn {
	cfg := config.ConfigDatabase()

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)
	db, err := pgx.Connect(context.Background(), psqlconn)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Connected!")
	return db
}
