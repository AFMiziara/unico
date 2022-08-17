package main

import (
	"fmt"
	"os"
	"time"

	router "github.com/fsvxavier/unico/api"
	"github.com/fsvxavier/unico/database"
	_ "github.com/fsvxavier/unico/docs"
	"github.com/fsvxavier/unico/pkg/enviroment"
	loggerLogrus "github.com/fsvxavier/unico/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /
func main() {

	//numCPUs := runtime.NumCPU()
	//runtime.GOMAXPROCS(numCPUs)
	os.Setenv("ENV", "production")
	var env enviroment.ConfigEnviroment
	env.SetFileConfig("./config/env.json")
	env.GetTag("ENV")

	var logLogrus loggerLogrus.GenericLogger
	logLogrus.GetLogger()
	logLogrus.LogIt("INFO", "Starting Api...", nil)

	currentTime := time.Now()

	// Start a new fiber app
	app := fiber.New()

	app.Use(recover.New())

	// Initialize config customization CORS
	var ConfigCors = cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "POST, OPTIONS, GET, PUT, DELETE", // "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS"
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		AllowCredentials: false,
		ExposeHeaders:    "*",
		MaxAge:           0,
	}

	app.Use(cors.New(ConfigCors))

	app.Use(func(c *fiber.Ctx) error {

		c.Accepts("application/json") // "application/json

		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// Go to next middleware:
		return c.Next()
	})

	app.Get("/pagemetrics", monitor.New(monitor.Config{
		Title:   "Fiber Monitor",
		Refresh: 5 * time.Second,
		APIOnly: false,
		Next:    nil,
	}))

	app.Get("/apimetrics", monitor.New(monitor.Config{
		APIOnly: true,
		Next:    nil,
	}))

	app.Use(pprof.New())

	// Initialize config default Logger
	fileLog, err := os.OpenFile("logs/"+currentTime.Format("2006-01-02")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		msgError := fmt.Sprintf("Error opening file: %v", err)
		logLogrus.LogIt("ERROR", msgError, nil)
	}

	loggerConfig := logger.Config{
		Next:         nil,
		Format:       "[${time}] ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       fileLog,
	}

	app.Use(logger.New(loggerConfig))

	// Provide a custom compression level
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	/*
		// Provide a minimal config
		app.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				"admin": "123456",
			},
		}))
	*/

	// Connect to the Database
	var dbConn database.DbConnect
	dbConn.ConnectDB()

	router.SetupRoutes(app)

	// Listen on PORT 5000
	app.Listen(":" + os.Getenv("PORT"))
}
