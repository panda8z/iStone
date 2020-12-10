package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/panda8z/istone/src/pkg/logger"
	"github.com/spf13/viper"
)

type Conf interface {
	//多db设置，⚠️SetDbs不允许并发,可以根据自己的业务，例如app分库、host分库
	SetDbs(key string, db *DBConfig)
	GetDbs() map[string]*DBConfig
	GetDbByKey(key string) *DBConfig
	GetSaas() bool
	SetSaas(bool)

	//单库业务实现这两个接口
	SetDb(db *DBConfig)
	GetDb() *DBConfig

	//使用的路由
	SetEngine(engine http.Handler)
	GetEngine() http.Handler

	//使用go-admin定义的logger，参考来源go-micro
	SetLogger(logger logger.ILog)
	GetLogger() logger.ILog
}

type Config struct {
	saas   bool
	dbs    map[string]*DBConfig
	db     *DBConfig
	engine http.Handler
}

type DBConfig struct {
	Driver string
	DB     *sql.DB
}

// SetDbs 设置对应key的db
func (c *Config) SetDbs(key string, db *DBConfig) {
	c.dbs[key] = db
}

// GetDbs 获取所有map里的db数据
func (c *Config) GetDbs() map[string]*DBConfig {
	return c.dbs
}

// GetDbByKey 根据key获取db
func (c *Config) GetDbByKey(key string) *DBConfig {
	return c.dbs[key]
}

// SetDb 设置单个db
func (c *Config) SetDb(db *DBConfig) {
	c.db = db
}

// GetDb 获取单个db
func (c *Config) GetDb() *DBConfig {
	return c.db
}

// SetEngine 设置路由引擎
func (c *Config) SetEngine(engine http.Handler) {
	c.engine = engine
}

// GetEngine 获取路由引擎
func (c *Config) GetEngine() http.Handler {
	return c.engine
}

// SetLogger 设置日志组件
func (c *Config) SetLogger(l logger.ILog) {
	logger.DefaultLogger = l
}

// GetLogger 获取日志组件
func (c *Config) GetLogger() logger.ILog {
	return logger.DefaultLogger
}

// SetSaas 设置是否是saas应用
func (c *Config) SetSaas(saas bool) {
	c.saas = saas
}

// GetSaas 获取是否是saas应用
func (c *Config) GetSaas() bool {
	return c.saas
}

func DefaultConfig() *Config {
	return &Config{}
}

// 数据库配置项
var cfgDatabase *viper.Viper

// 应用配置项
var cfgApplication *viper.Viper

// Token配置项
var cfgJwt *viper.Viper

// Log配置项
var cfgLogger *viper.Viper

// Ssl配置项 非必须
var cfgSsl *viper.Viper

// 代码生成配置项 非必须
var cfgGen *viper.Viper

//载入配置文件
func Setup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	cfgDatabase = viper.Sub("settings.database")
	if cfgDatabase == nil {
		panic("No found settings.database in the configuration")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("No found settings.application in the configuration")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	cfgJwt = viper.Sub("settings.jwt")
	if cfgJwt == nil {
		panic("No found settings.jwt in the configuration")
	}
	JwtConfig = InitJwt(cfgJwt)

	cfgLogger = viper.Sub("settings.logger")
	if cfgLogger == nil {
		panic("No found settings.logger in the configuration")
	}
	LoggerConfig = InitLog(cfgLogger)

	cfgSsl = viper.Sub("settings.ssl")
	if cfgSsl == nil {
		// Ssl不是系统强制要求的配置，默认可以不用配置，将设置为关闭状态
		fmt.Println("warning config not found settings.ssl in the configuration")
		SslConfig = new(Ssl)
		SslConfig.Enable = false
	} else {
		SslConfig = InitSsl(cfgSsl)
	}

	cfgGen = viper.Sub("settings.gen")
	if cfgGen == nil {
		panic("No found settings.gen")
	}
	GenConfig = InitGen(cfgGen)
}
