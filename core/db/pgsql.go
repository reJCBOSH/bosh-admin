package db

import (
	"fmt"

	"bosh-admin/config"
	"bosh-admin/core/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ConnectPgsql 连接pgsql
func ConnectPgsql(p config.Pgsql) *gorm.DB {
	if p.Dbname == "" {
		return nil
	}
	err := initPgsql(p)
	if err != nil {
		log.Error("初始化数据库失败:", err)
		return nil
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s", p.Host, p.User, p.Password, p.Dbname, p.Port, p.Config)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
		Logger:                                   log.CustomGormLogger("_pgsql"),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	return db
}

func initPgsql(p config.Pgsql) error {
	// 首先连接到默认的postgres数据库
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s %s", p.Host, p.User, p.Password, p.Port, p.Config)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("连接pgsql服务器失败:", err)
		return err
	}

	// 检查数据库是否存在
	var count int64
	db.Raw("SELECT count(*) FROM pg_database WHERE datname = ?", p.Dbname).Scan(&count)
	if count == 0 {
		// 创建数据库
		err = db.Exec(fmt.Sprintf(`CREATE DATABASE "%s"`, p.Dbname)).Error
		if err != nil {
			return err
		}
	}

	return nil
}
