package db

import (
	"database/sql"
	"fmt"
	config "gin-demo/config"
	"gin-demo/defs"
	logger "gin-demo/logger"

	"github.com/go-redis/redis"

	_ "github.com/go-sql-driver/mysql" /* mysql driver init */
	"github.com/jinzhu/gorm"
)

/* TODO: split dbs */
var (
	gormMysqlClient *gorm.DB
	sqlClient       *sql.DB
	redisClient     *redis.Client
	dbCfg           = config.Config().DB
	err             error
	mysqlCfg        = "Username:Password@tcp(Host:Port)/DbName?charset=utf8"
)

func init() {
	InitCfg(&mysqlCfg, dbCfg.Mysql)
	logger.Noticef("Mysql config: %s", mysqlCfg)
}

// ConnGormMysql gormMysql
func ConnGormMysql() *gorm.DB {
	if gormMysqlClient == nil {
		gormMysqlClient, err = gorm.Open("mysql", mysqlCfg)
		if err != nil {
			logger.Error(defs.ConnDBErr, err.Error())
		}
		gormMysqlClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&TBL_USERS{})
		// 为`name`列添加索引`idx_user_name`
		// gormMysqlClient.Model(&TBL_USERS{}).AddIndex("idx_user_name", "name")
	}
	return gormMysqlClient
}

// ConnMysql  connect to mysql
func ConnMysql() *sql.DB {
	if sqlClient == nil || sqlClient.Ping() != nil {
		logger.Info("mysql database connection initialization ...")
		sqlClient, err = sql.Open("mysql", mysqlCfg)
		if err != nil {
			logger.Error(defs.ConnDBErr, err.Error())
			panic(err.Error())
		}
	}
	return sqlClient
}

// ConnRedis connect to redis
func ConnRedis() *redis.Client {
	if redisClient == nil {
		logger.Info("redis database connection initialization ...")
		client := redis.NewClient(&redis.Options{
			Addr:     dbCfg.Redis.Host + ":" + dbCfg.Redis.Port,
			Password: dbCfg.Redis.Password,
			DB:       dbCfg.Redis.DbName,
		})
		fmt.Println(client.Ping().Result())
	}
	return redisClient
}
