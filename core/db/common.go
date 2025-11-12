package db

import (
	"math/rand"

	"bosh-admin/global"

	"gorm.io/gorm"
)

var NotFound = gorm.ErrRecordNotFound

func GormDB() *gorm.DB {
	return global.GormDB
}

func Begin() *gorm.DB {
	return global.GormDB.Begin()
}

func Create(value interface{}, table ...string) error {
	DB := global.GormDB
	if len(table) > 0 {
		DB = DB.Table(table[0])
	}
	return DB.Create(value).Error
}

func Updates(value interface{}, table ...string) error {
	DB := GormDB()
	if len(table) > 0 {
		DB = DB.Table(table[0])
	}
	return DB.Select("*").Updates(value).Error
}

func QueryById[T any](id any) (*T, error) {
	data := new(T)
	err := global.GormDB.First(data, id).Error
	return data, err
}

func DelById[T any](id any) error {
	model := new(T)
	return global.GormDB.Delete(model, id).Error
}

func DelByIds[T any](ids ...any) error {
	model := new(T)
	return GormDB().Delete(model, ids...).Error
}

// PageScope 分页作用域
func PageScope(pageNo, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNo > 0 {
			db = db.Offset((pageNo - 1) * pageSize).Limit(pageSize)
		}
		return db
	}
}

// OrderByScope 排序作用域
func OrderByScope(orderStr string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if orderStr != "" {
			db = db.Order(orderStr)
		} else {
			db = db.Order("id DESC")
		}
		return db
	}
}

// RandomOrderScope 随机作用域
func RandomOrderScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch db.Dialector.Name() {
		case "postgres", "sqlite":
			db = db.Order("RANDOM()")
		default: // mysql等
			db = db.Order("RAND()")
		}
		return db
	}
}

// SafeRandomOrderScope 事务安全随机作用域
func SafeRandomOrderScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Session(&gorm.Session{}).Scopes(RandomOrderScope())
	}
}

// OptimizedRandomOrderScope 大表随机作用域
func OptimizedRandomOrderScope(table interface{}, pkField ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 确定主键字段名
		pk := "id"
		if len(pkField) > 0 && pkField[0] != "" {
			pk = pkField[0]
		}
		// 创建新会话避免污染原查询
		tx := db.Session(&gorm.Session{})
		// 动态设置表名或模型
		if name, ok := table.(string); ok {
			tx = tx.Table(name)
		} else {
			tx = tx.Model(table)
		}
		// 获取最大ID值
		var maxID int
		tx.Select("MAX(" + pk + ")").Scan(&maxID)
		if maxID <= 0 {
			// 如果表为空或主键非数字，回退到简单随机排序
			switch tx.Dialector.Name() {
			case "postgres", "sqlite":
				return tx.Order("RANDOM()")
			default: // mysql等
				return tx.Order("RAND()")
			}
		}
		// 生成随机ID范围查询
		randomID := rand.Intn(maxID)
		return tx.Where(pk+" >= ?", randomID).Order(pk).Limit(1)
	}
}
