package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fsvxavier/unico/internal/models"
	"github.com/fsvxavier/unico/pkg/logger"
	gorm_logrus "github.com/onrik/gorm-logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

// Declare the variable for the database
var (
	DbConn *gorm.DB
)

type DbConnect struct{}

// ConnectDB connect to db
func (d *DbConnect) ConnectDB() *gorm.DB {

	logg := new(logger.GenericLogger)
	logg.Module = "api"
	logg.GetLogger()

	// Connection URL to connect to Postgres Master Database
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed to convert string to int. Erro: %s", err.Error()), nil)
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	// Connection URL to connect to Postgres Replica Database
	pSL := os.Getenv("DB_SL_PORT")
	portSL, err := strconv.ParseUint(pSL, 10, 32)
	if err != nil {
		logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed to convert string to int. Erro: %s", err.Error()), nil)
	}
	dsn_replica := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_SL_HOST"),
		portSL, os.Getenv("DB_SL_USER"),
		os.Getenv("DB_SL_PASSWORD"),
		os.Getenv("DB_SL_NAME"))

	// Connect to the DB and initialize the DB variable
	DbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: gorm_logrus.New(),
	})
	if err != nil {
		logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed to connect database. Erro: %s", err.Error()), nil)
	}

	DbConn.Use(dbresolver.Register(dbresolver.Config{
		// use `dns` a source, `dns_replica` a replica
		Sources:  []gorm.Dialector{postgres.Open(dsn)},
		Replicas: []gorm.Dialector{postgres.Open(dsn_replica)},
		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
	}))

	sqlDB, err := DbConn.DB()

	// Ping Database connection verify
	errPing := sqlDB.Ping()
	if errPing != nil {
		logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed to ping database. Erro: %s", err.Error()), nil)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	DB_MAX_IDLE_CONNS, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	sqlDB.SetMaxIdleConns(DB_MAX_IDLE_CONNS)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	DB_MAX_OPEN_CONNS, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	sqlDB.SetMaxOpenConns(DB_MAX_OPEN_CONNS)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	DB_MAX_LIFE_TIME, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFE_TIME"))
	sqlDB.SetConnMaxLifetime(time.Duration(DB_MAX_LIFE_TIME) * time.Minute)

	// Migrate the database
	execMigration, err := strconv.ParseBool(os.Getenv("EXECUTE_MIGRATION"))
	if err != nil {
		logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed set execute migration boolean. Erro: %s", err.Error()), nil)
	}
	if execMigration {

		err = DbConn.AutoMigrate(&models.FeiraLivre{})
		if err != nil {
			logg.LogIt("FATAL", fmt.Sprintf("[FATAL] - Failed execute migration database. Erro: %s", err.Error()), nil)
		} else {
			fmt.Println("Database Migrated")
		}

	}

	return DbConn
}
