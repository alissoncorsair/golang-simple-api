package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	err = createProductTable(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createProductTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		product_name VARCHAR(100) NOT NULL,
		price NUMERIC(10, 2) NOT NULL
	);
	`

	_, err := db.Exec(query)

	if err != nil {
		return fmt.Errorf("failed to create product table: %w", err)
	}

	fmt.Println("Product table is ready!")
	return nil
}
