package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/raflibima25/task-management/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresConnection() (*gorm.DB, error) {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
		sslmode  = os.Getenv("DB_SSL_MODE")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	// logger GORM
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // ambang batas slow query
			LogLevel:                  logger.Info, // log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // abaikan error record not found
			Colorful:                  true,        // aktifkan warna
		},
	)

	// koneksi ke database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	// set connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// set max idle connection
	sqlDB.SetMaxIdleConns(10)
	// set max open connection
	sqlDB.SetMaxOpenConns(100)
	// set max connection lifetime
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

// migrate database
func MigrateDB(db *gorm.DB) error {
	err := db.Exec(`
		DO $$ 
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_status') THEN
				CREATE TYPE task_status AS ENUM ('TODO', 'IN_PROGRESS', 'COMPLETED');
			END IF;
			
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_priority') THEN
				CREATE TYPE task_priority AS ENUM ('LOW', 'MEDIUM', 'HIGH');
			END IF;
		END $$;
	`).Error
	if err != nil {
		return err
	}

	// auto migrate model
	err = db.AutoMigrate(
		&model.User{},
		&model.Task{},
		&model.Category{},
	)
	if err != nil {
		return err
	}

	log.Println("Database migrated successfully")
	return nil
}
