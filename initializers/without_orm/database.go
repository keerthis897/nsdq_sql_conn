package initializers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func ConnectDB() {
	server := os.Getenv("DB_SERVER")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s;",
		server, port, user, password, database)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	DB = db
}

//  to provide access to the global DB variable
func GetDB() *sql.DB {
    return DB
}