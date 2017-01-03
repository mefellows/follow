package db

import (
	"fmt"
	"log"
	"regexp"

	"github.com/mattes/migrate/migrate"

	// Postgres driver
	"github.com/jinzhu/gorm"

	// Postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// GetDatabase returns a database handler for a given type.
func GetDatabase(connectionString string) (db *gorm.DB) {
	var err error
	adapter, gormString := ConnectionStringToGormString(connectionString)
	if db, err = gorm.Open(adapter, gormString); err != nil {
		log.Fatal("could not create database connection:", err)
	}
	log.Println("[DEBUG] setting default table name to pluralized")
	db.SingularTable(false)

	// log.Println("[DEBUG] setting auto migrate on models")
	// db.AutoMigrate(&models.User{}, &models.Location{})

	return
}

// ConnectionStringToGormString Convert Postgres connection string to Gorm string
func ConnectionStringToGormString(connectionString string) (adapter string, gormString string) {

	r := regexp.MustCompile(`([a-zA-Z]+)\:\/\/([a-zA-Z0-9_]+)\:([a-zA-Z0-9_]+)\@([0-9a-zA-Z\_\-\.]+)\:([0-9]{4,5})/([a-zA-Z0-9_]+)(.*)`)
	matches := r.FindStringSubmatch(connectionString)

	return matches[1], fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", matches[4], matches[2], matches[6], matches[3], matches[5])
}

// Migrate runs DB migrations.
func Migrate(connectionString string) []error {
	allErrors, ok := migrate.UpSync(connectionString, "./db/migrations")
	if !ok {
		log.Println("[ERROR] unable to run database migrations:", allErrors)
	}
	return allErrors
}
