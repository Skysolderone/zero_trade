package svc

import (
	"time"

	"trade/user-rpc/internal/config"
	"trade/user-rpc/internal/model"

	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Orm    *gorm.DB
	Rdb    redis.ClusterClient
	Bun    bun.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:gg123456@tcp(127.0.0.1:3306)/zero_trade?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                                  // default size for string fields
		DisableDatetimePrecision:  true,                                                                                 // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                 // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                 // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour * 24)

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"}, // Password: "", // no password set
		// DB:       0,  // use default DB
	})
	return &ServiceContext{
		Config: c,
		Orm:    db,
		Rdb:    *rdb,
	}
}
