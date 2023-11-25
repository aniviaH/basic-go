package dao

import "gorm.io/gorm"

// InitTable 后续有新的数据要存储， 都要来这里初始化一下表。
func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
