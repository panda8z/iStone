package tools

import (
	"errors"
	"github.com/panda8z/istone/src/pkg/security"

	"gorm.io/gorm"

	orm "github.com/panda8z/istone/src/pkg/global"
	config2 "github.com/panda8z/istone/src/pkg/config"
)

type DBTables struct {
	TableName      string `gorm:"column:TABLE_NAME" json:"tableName"`
	Engine         string `gorm:"column:ENGINE" json:"engine"`
	TableRows      string `gorm:"column:TABLE_ROWS" json:"tableRows"`
	TableCollation string `gorm:"column:TABLE_COLLATION" json:"tableCollation"`
	CreateTime     string `gorm:"column:CREATE_TIME" json:"createTime"`
	UpdateTime     string `gorm:"column:UPDATE_TIME" json:"updateTime"`
	TableComment   string `gorm:"column:TABLE_COMMENT" json:"tableComment"`
}

func (e *DBTables) GetPage(pageSize int, pageIndex int) ([]DBTables, int, error) {
	var doc []DBTables
	table := new(gorm.DB)
	var count int64

	if config2.DatabaseConfig.Driver == "mysql" {
		table = orm.Eloquent.Table("information_schema.tables")
		table = table.Where("TABLE_NAME not in (select table_name from `" + config2.GenConfig.DBName + "`.sys_tables) ")
		table = table.Where("table_schema= ? ", config2.GenConfig.DBName)

		if e.TableName != "" {
			table = table.Where("TABLE_NAME = ?", e.TableName)
		}
		if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
			return nil, 0, err
		}
	} else {
		security.Assert(true, "目前只支持mysql数据库", 500)
	}

	//table.Count(&count)
	return doc, int(count), nil
}

func (e *DBTables) Get() (DBTables, error) {
	var doc DBTables
	table := new(gorm.DB)
	if config2.DatabaseConfig.Driver == "mysql" {
		table = orm.Eloquent.Table("information_schema.tables")
		table = table.Where("table_schema= ? ", config2.GenConfig.DBName)
		if e.TableName == "" {
			return doc, errors.New("table name cannot be empty！")
		}
		table = table.Where("TABLE_NAME = ?", e.TableName)
	} else {
		security.Assert(true, "目前只支持mysql数据库", 500)
	}
	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}
