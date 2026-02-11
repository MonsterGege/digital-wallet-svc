package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/digital-wallet-svc/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	godotenv.Load(".env")
	// Initialize database connection
	db, err := database.NewDatabase(ctx, os.Getenv("DB_HOST"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"))
	if err != nil {
		return
	}
	defer db.CloseConnection(ctx)
	// Path to the SQL file
	sqlFilePath := `scripts\user.sql`
	// Read the SQL file content
	sqlFile, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		log.Fatal(err)
	}
	// Execute the SQL commands Create table from the file
	_, err = db.GetConnection(ctx).Exec(ctx, string(sqlFile))
	if err != nil {
		log.Fatal(err)
	}
	return
}
