package cmd

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"
	"time"

	"github.com/panda8z/istone/src/pkg/database"
	"github.com/panda8z/istone/src/pkg/global"
	"github.com/panda8z/istone/src/pkg/logger"

	"github.com/panda8z/istone/src/pkg/config"
	"github.com/panda8z/istone/src/pkg/models"
	"github.com/panda8z/istone/src/pkg/tools"
	"github.com/spf13/cobra"
)

var (
	generate bool
	iStone   bool
)

var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "Initialize the database",
	Example: "istone migrate -c config/settings.yml",
	Run: func(cmd *cobra.Command, args []string) {
		runMigrateCmd()
	},
}

func initMigrateCmd() {
	migrateCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	migrateCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate migration file")
	migrateCmd.PersistentFlags().BoolVarP(&iStone, "goAdmin", "a", false, "generate istone migration file")
}

func runMigrateCmd() {
	usage := `start init`
	fmt.Println(usage)

	if !generate {
		//1. 读取配置
		config.Setup(configYml)
		//2. 设置日志
		global.Logger.ILog = logger.SetupLogger(config.LoggerConfig.Path, "bus")
		global.JobLogger.ILog = logger.SetupLogger(config.LoggerConfig.Path, "job")
		global.RequestLogger.ILog = logger.SetupLogger(config.LoggerConfig.Path, "request")
		_ = initDB()
	} else {
		_ = genFile()
	}
}

func migrateModel() error {
	if config.DatabaseConfig.Driver == "mysql" {
		global.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	err := global.Eloquent.Debug().AutoMigrate(&models.Migration{})
	if err != nil {
		return err
	}
	Migrate.SetDb(global.Eloquent.Debug())
	Migrate.Migrate()
	return err
}
func initDB() error {
	//3. 初始化数据库链接
	database.Setup(config.DatabaseConfig.Driver)
	//4. 数据库迁移
	fmt.Println("数据库迁移开始")
	_ = migrateModel()
	fmt.Println(`数据库基础数据初始化成功`)
	return nil
}

func genFile() error {
	t1, err := template.ParseFiles("template/migrate.template")
	if err != nil {
		return err
	}
	m := map[string]string{}
	m["GenerateTime"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	m["Package"] = "version_local"
	if iStone {
		m["Package"] = "version"
	}
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	if iStone {
		tools.FileCreate(b1, "./cmd/migrate/migration/version/"+m["GenerateTime"]+"_migrate.go")
	} else {
		tools.FileCreate(b1, "./cmd/migrate/migration/version-local/"+m["GenerateTime"]+"_migrate.go")
	}
	return nil
}
