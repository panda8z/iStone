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

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	gormLogger"gorm.io/gorm/logger"


)

type PgSql struct {
}

func (e *PgSql) Setup() {
	var err error

	global.Source = e.GetConnect()
	log.Info(global.Source)
	db, err := sql.Open("postgresql", global.Source)
	if err != nil {
		global.Logger.Fatal(tools.Red(e.GetDriver()+" connect error :"), err)
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
		log.Fatalf("%s connect error %v", e.GetDriver(), err)
	} else {
		log.Infof("%s connect success!", e.GetDriver())
	}

	if global.Eloquent.Error != nil {
		log.Fatalf("database error %v", global.Eloquent.Error)
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

// 打开数据库连接
func (e *PgSql) Open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{Conn: db}), cfg)
}

func (e *PgSql) GetConnect() string {
	return config.DatabaseConfig.Source
}

func (e *PgSql) GetDriver() string {
	return config.DatabaseConfig.Driver
}
