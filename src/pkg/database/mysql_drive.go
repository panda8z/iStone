package database

import (
	"database/sql"
	"github.com/panda8z/istone/src/pkg/config"
	"github.com/panda8z/istone/src/pkg/global"
	"github.com/panda8z/istone/src/pkg/log"
	iLogger "github.com/panda8z/istone/src/pkg/logger"
	"github.com/panda8z/istone/src/pkg/tools"
	. "log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Mysql mysql配置结构体
type Mysql struct {
}

// Setup 配置步骤
func (e *Mysql) Setup() {
	global.Source = e.GetConnect()
	log.Info(tools.Green(global.Source))
	db, err := sql.Open("mysql", global.Source)
	if err != nil {
		log.Fatal(tools.Red(e.GetDriver()+" connect error :"), err)
	}
	global.Cfg.SetDb(&config.DBConfig{
		Driver: "mysql",
		DB:     db,
	})
	global.Eloquent, err = e.Open(db, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(tools.Red(e.GetDriver()+" connect error :"), err)
	} else {
		log.Info(tools.Green(e.GetDriver() + " connect success !"))
	}

	if global.Eloquent.Error != nil {
		log.Fatal(tools.Red(" database error :"), global.Eloquent.Error)
	}

	if config.LoggerConfig.EnabledDB {
		global.Eloquent.Logger = gormLogger.New(
			New(iLogger.DefaultLogger.Options().Out, "\r\n", LstdFlags),
			gormLogger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel: gormLogger.LogLevel(
					iLogger.DefaultLogger.Options().Level.LevelForGorm()),
			})
	}
}

// Open 打开数据库连接
func (e *Mysql) Open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), cfg)
}

// GetConnect 获取数据库连接
func (e *Mysql) GetConnect() string {
	return config.DatabaseConfig.Source
}

// GetDriver 获取连接
func (e *Mysql) GetDriver() string {
	return config.DatabaseConfig.Driver
}
