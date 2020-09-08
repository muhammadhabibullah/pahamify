package repository

import (
	"fmt"
	"log"
	"sync"
	"time"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"pahamify/project/config"
	"pahamify/project/model"
)

var (
	mysqlDB *gorm.DB
	once    sync.Once
)

// GetMySQL returns MySQL database connection instance
func GetMySQL() *gorm.DB {
	once.Do(func() {
		cfg := config.GetConfig()
		dsn := getMySQLConnString(cfg.Mysql)

		var err error
		mysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to mysql database: %s", err)
		}

		if cfg.Migrate {
			if err = mysqlDB.Set("gorm:table_options", "ENGINE=InnoDB").
				AutoMigrate(
					&model.Pokemon{},
					&model.Type{},
				); err != nil {
				log.Fatalf("failed to migrate new model to mysql database: %s", err)
			}

			migrations := &migrate.FileMigrationSource{
				Dir: "migration",
			}

			mysqlDBConn, _ := mysqlDB.DB()
			result, err := migrate.Exec(mysqlDBConn, "mysql", migrations, migrate.Up)
			if err != nil {
				log.Fatalf("migration data failed: %s", err)
			}
			log.Printf("applied %d migrations successfully", result)
		}

		configMySQLConn(cfg.Mysql)
	})

	return mysqlDB
}

// getMySQLConnString return connection string from config
func getMySQLConnString(cfg config.MySQLConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name)
}

// configMySQLConn configure MySQLConnection settings
func configMySQLConn(cfg config.MySQLConfig) {
	db, _ := mysqlDB.DB()
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)
}
