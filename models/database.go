package models

import (
	"ulfr/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() (*Database, error) {
	Dbs := config.MysqlUser + ":" + config.MysqlPass + "@tcp(" + config.MysqlHost + ":" + config.MysqlPort + ")/" + config.MysqlDb + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(Dbs), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (database *Database) Migrate(obj interface{}) error {
	err := database.db.AutoMigrate(obj)
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) FindAll(obj interface{}) (any, error) {
	var objs = obj
	err := database.db.Order("id DESC").Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (database *Database) Find2(obj, where any) error {
	return database.db.Order("id DESC").Find(obj, where).Error
}

func (database *Database) Find(obj interface{}, where interface{}) (any, error) {
	var objs = obj
	err := database.db.Order("id DESC").Find(&objs, where).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (database *Database) Create(obj interface{}) error {
	err := database.db.Create(obj).Error
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) Update(obj interface{}, column string) error {
	err := database.db.Model(obj).Update(column, obj).Error
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) Updates(obj interface{}, data any) error {
	err := database.db.Model(obj).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) Delete(obj interface{}) error {
	err := database.db.Delete(obj).Error
	if err != nil {
		return err
	}
	return nil
}
