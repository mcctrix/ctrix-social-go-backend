package v1

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func PopulateDB() {
	db, err := DBConnection()
	if err != nil {
		log.Fatal(err)
	}

	sqlFile, err := os.ReadFile("./sql/populateDB.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Split the SQL file into individual statements
	statements := strings.Split(string(sqlFile), ";")

	// Execute each statement separately
	for _, stmt := range statements {
		// Skip empty statements
		if strings.TrimSpace(stmt) == "" {
			continue
		}

		// Execute the statement
		err = db.Exec(stmt).Error
		if err != nil {
			fmt.Printf("Error executing statement: %v\nStatement: %s\n", err, stmt)
			continue
		}
	}

	fmt.Println("Database populated successfully!")
	fmt.Println("Please Restart the server to use the new database")
}
