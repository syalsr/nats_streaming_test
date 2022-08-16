package database

import (
	"Wildberries_L0/detail"
	"database/sql/driver"
	"fmt"
	"github.com/joho/godotenv"
	sql "github.com/lib/pq"
	"os"
	"strconv"
)

func InsertData(connection *sql.Connector, info detail.OrderInfo) {

}

func Connect() (driver.Conn, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbname := os.Getenv("DATABASE")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open(psqlconn)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected!")

	return db, nil
}
