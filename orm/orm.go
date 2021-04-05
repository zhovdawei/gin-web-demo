package orm

import "github.com/jinzhu/gorm"

var DataSource *gorm.DB

func TableRegistry(db *gorm.DB) {
	DataSource = db

	// 设置统一表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}

	// 禁用复数表名
	db.SingularTable(true)

	// 创建表 自动迁移
	db.AutoMigrate(&Product{})
}
