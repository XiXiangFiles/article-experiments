package connectionhandler

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQLConnection struct {
	Host            string
	DBName          string
	EnsureDBCreated bool
	DB              *gorm.DB
}

func NewMySQLConnection(host string, dbName string, EnsureDBCreated bool) *MySQLConnection {
	c := &MySQLConnection{
		Host:            host,
		DBName:          dbName,
		EnsureDBCreated: EnsureDBCreated,
	}
	return c
}

func (cli *MySQLConnection) Connect() error {
	dsnParams := "charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s/%s?%s", cli.Host, cli.DBName, dsnParams)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	config := &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           IndexIXNamingNamingStrategy{},
	}

	db, err := gorm.Open(mysql.Open(dsn), config)

	if err != nil {
		if strings.Contains(err.Error(), "Error 1049") && cli.EnsureDBCreated {
			dsnWithoutDBName := fmt.Sprintf("%s/?%s", cli.Host, dsnParams)

			db, err = gorm.Open(mysql.Open(dsnWithoutDBName), config)
			if err != nil {
				return err
			}
			db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8;", cli.DBName))
			db.Exec(fmt.Sprintf("use %s;", cli.DBName))
			db, err = gorm.Open(mysql.Open(dsn), config)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	cli.DB = db
	return nil
}
