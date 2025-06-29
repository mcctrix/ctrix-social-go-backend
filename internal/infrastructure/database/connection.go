package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mcctrix/ctrix-social-go-backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbInstance *gorm.DB

func GetDB() *gorm.DB {
	return dbInstance
}

func DBConnect(config *config.DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.User, config.Password, config.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}
	dbInstance = db
	return dbInstance, nil
}

func CreateInitialDBStructure() {

	db := GetDB()

	sqlFile, err := os.ReadFile("./internal/infrastructure/database/migrations/001_initial_schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Exec(string(sqlFile)).Error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Initial Tables are Created Successfully!!!")
	}

}
func ResetDB() {
	db := GetDB()

	// Retrieve all table names
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatalf("failed to get tables: %v", err)
	}

	// Drop all tables
	for _, table := range tables {
		err := db.Migrator().DropTable(table)
		if err != nil {
			log.Printf("failed to drop table %s: %v", table, err)
		} else {
			log.Printf("dropped table %s successfully", table)
		}
	}

}

func PopulateDB() {
	db := GetDB()

	sqlFile, err := os.ReadFile("./internal/infrastructure/database/migrations/002_populate_data.sql")
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

func InitNewUser(userid string) error {
	db := GetDB()

	type base struct {
		Id string `gorm:"primaryKey"`
	}

	data := base{Id: userid}

	if err := db.Table("user_profile").Create(data).Error; err != nil {
		return err
	}
	if err := db.Table("user_additional_info").Create(data).Error; err != nil {
		return err
	}
	userSetting := struct {
		Id          string `gorm:"primaryKey"`
		Show_online bool
	}{
		Id:          userid,
		Show_online: true,
	}

	if err := db.Table("user_settings").Create(userSetting).Error; err != nil {
		return err
	}
	return nil
}
