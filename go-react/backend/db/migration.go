package db

import (
	"backend/config"
	"database/sql"
	"embed"

	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations
var migrations embed.FS

func ConnectDB(config *config.Config) *sql.DB {
	// First connect without database name to create it if it doesn't exist
	rootDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort)

	rootDB, err := sql.Open("mysql", rootDSN)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer rootDB.Close()

	// Create database if it doesn't exist
	_, err = rootDB.Exec("CREATE DATABASE IF NOT EXISTS " + config.DBName)
	if err != nil {
		log.Fatal("Error creating database:", err)
	}

	// Now connect to the specific database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	if err := RunMigrations(dsn, db); err != nil {
		log.Fatal("Error running migrations:", err)
	}

	return db
}

func RunMigrations(dsn string, db *sql.DB) error {
	source, err := iofs.New(migrations, "migrations")
	if err != nil {
		return fmt.Errorf("could not create the migration source: %v", err)
	}

	m, err := migrate.NewWithSourceInstance(
		"iofs",
		source,
		"mysql://"+dsn,
	)
	if err != nil {
		return fmt.Errorf("error creating migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error running migrations: %v", err)
	}

	log.Println("Migrations completed successfully")
	return nil
}
