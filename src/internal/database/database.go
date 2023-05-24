package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbParams struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func New(params DbParams) (*gorm.DB, error) {
	var err error

	// TODO Need add app logger to gorm config
	//fileName := fmt.Sprintf("%s.log", time.Now().Format("A200601021504051"))
	//
	//var logOutput io.Writer
	//logOutput, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	////var l = logging.NewLogger(logging.Config{Output: logOutput})
	//
	//newLogger := logger.New(
	//	log.New(logOutput, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold:             time.Second, // Slow SQL threshold
	//		LogLevel:                  logger.Info, // Log level
	//		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	//		Colorful:                  false,       // Disable color
	//	},
	//)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		params.Host, params.User, params.Password, params.DbName, params.Port)

	// it needs to add check DB connecting error
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		//log.Fatal("Failed to connect to database. \n", err)
		return nil, err
	}

	//log.Println("\nDatabase Connection Opened")
	return dbConn, nil
}
