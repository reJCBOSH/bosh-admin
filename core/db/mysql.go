package db

import (
	"fmt"

	"bosh-admin/config"
	"bosh-admin/core/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ConnectMysql 连接Mysql
func ConnectMysql(m config.Mysql) *gorm.DB {
	if m.Database == "" {
		return nil
	}
	err := initMysql(m)
	if err != nil {
		log.Error("初始化数据库失败:", err)
		return nil
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.IP, m.Port, m.Database, m.Config)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
		Logger:                                   log.CustomGormLogger("_mysql"),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	return db
}

// 初始化Mysql数据库
func initMysql(m config.Mysql) error {
	// 不指定数据库连接Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?%s", m.Username, m.Password, m.IP, m.Port, m.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("连接Mysql服务器失败:", err)
		return err
	}
	// 检查数据库是否存在
	var count int64
	db.Raw("SELECT COUNT(*) FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = ?", m.Database).Scan(&count)
	// 不存在则创建数据库
	if count == 0 {
		sql := fmt.Sprintf("CREATE DATABASE `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", m.Database)
		if err = db.Exec(sql).Error; err != nil {
			log.Error("创建数据库失败:", err)
			return err
		}
		log.Info("创建数据库成功:", m.Database)
	}
	return nil
}
