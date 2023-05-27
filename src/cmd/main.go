package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/pavelkg/tradem-mon-api/docs"
	"github.com/pavelkg/tradem-mon-api/internal/config"
	"github.com/pavelkg/tradem-mon-api/internal/database"
	"github.com/pavelkg/tradem-mon-api/internal/presenter"
	"github.com/pavelkg/tradem-mon-api/internal/router"

	"github.com/pavelkg/tradem-mon-api/internal/domain/service"
	"github.com/pavelkg/tradem-mon-api/internal/repository"
)

// @title Trademark monitoring
// @version 1.0
// @description Trademark monitoring API
// @termsOfService https://trmon.pepex.kg/terms/
// @contact.name Undefined
// @contact.email admin@pepex.kg
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7654
// @BasePath /api

func main() {

	var logOutput io.Writer
	conf := config.GetConfig()

	if !conf.IsDevelopment {
		var err error
		fileName := fmt.Sprintf("%s.log", time.Now().Format("A20060102150405"))
		logOutput, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

	const jwtSecret = "ilsdahJHILUuygaiosb2345"
	//var log = logging.NewLogger(logging.Config{Output: logOutput})

	dbConn, err := database.New(conf.DbParams)
	if err != nil {
		log.Fatal(fmt.Sprint("{APP} Failed to connect to DB: ", err))
	}
	log.Println("{APP} Database Connection Opened")

	repository, err := repository.NewRepository(repository.Config{DB: dbConn})
	if err != nil {
		log.Fatal(fmt.Sprint("Failed to create App repository: ", err))
	}
	log.Println("{APP} App repository loaded")

	services, err := service.NewServices(repository)
	if err != nil {
		log.Fatal(fmt.Sprint("Failed to load App service: ", err))
	}
	log.Println("{APP} App service loaded")

	presenter, err := presenter.NewPresenters(services)

	app := fiber.New(fiber.Config{BodyLimit: 10 * 1024 * 1024})
	app.Use(cors.New(cors.Config{ExposeHeaders: "Location"}))
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "INFO: ${time} {API} ${status} - ${method} ${path} [${ip}] (${latency})\n",
		Output:     logOutput,
		TimeFormat: "2006/01/02 15:04:05",
	}))
	app.Static("/files/img", "./static/img")

	router.SetupRoutes(app, presenter, conf.HostPrefix, jwtSecret)

	json.MarshalIndent(app.Stack(), "", "  ")

	log.Println(fmt.Sprintf("Server Host:%v Port:%v is starting...", conf.Listen.Host, conf.Listen.Port))
	app.Hooks().OnListen(func() error {
		log.Println("Server started")
		return nil
	})
	err = app.Listen(fmt.Sprintf("%v:%v", conf.Listen.Host, conf.Listen.Port))

	if err != nil {
		return
	}
}
