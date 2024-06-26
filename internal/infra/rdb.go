package infra

import (
	"golang.org/x/exp/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sapphire-server/internal/conf"
)

var (
	DB *gorm.DB
)

func InitDB() error {
	dsn := conf.GetDBConfig()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	slog.Info("DB connected")
	return nil
}

// Insert
// 通过范型实现通用 Insert 函数
func Insert[T any](data T) error {
	res := DB.Create(&data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// UpdateSingleColumn
// 通过范型实现通用 UpdateSingleColumn 函数
func UpdateSingleColumn[T any](data T, column string, value interface{}) error {
	res := DB.Model(&data).Update(column, value)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// FindOne 查询一条数据
func FindOne[T any](conditions ...interface{}) (*T, error) {
	var obj T
	// 这里不使用 `Take()` 方法，因为 `Take()` 方法在没有找到数据时会返回 ErrRecordNotFound 错误
	res := DB.Limit(1).Find(&obj, conditions...)
	if res.Error != nil {
		return nil, res.Error
	}
	return &obj, nil
}

// First 查询第一条数据
func First[T any](conditions ...interface{}) (*T, error) {
	// 和 FindOne 的区别是，当没有找到数据时，First 会返回 ErrRecordNotFound 错误
	var obj T
	res := DB.First(&obj, conditions...)
	// if errors.Is(res.Error, gorm.ErrRecordNotFound) {
	// 	return nil, nil
	// }
	if res.Error != nil {
		return nil, res.Error
	}
	return &obj, nil
}

// FindAll 查询所有数据
func FindAll[T any](conditions ...interface{}) ([]T, error) {
	var objs []T
	res := DB.Find(&objs, conditions...)
	if res.Error != nil {
		return nil, res.Error
	}
	return objs, nil
}
