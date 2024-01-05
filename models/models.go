package models

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type Path struct {
	ID          uint   `gorm:"primarykey"`
	Domain      uint   `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string
	Value       string `gorm:"not null"`
	Type        uint   `gorm:"not null"`
	gorm.Model
}

type Domain struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"not null"`
	Description string
	gorm.Model
}

type BrowserInformation map[string]interface{}
type UserInformation map[string]interface{}
type SiteInformation map[string]interface{}
type Cookie map[string]interface{}

func (bi *BrowserInformation) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), bi)
}

func (bi BrowserInformation) Value() (driver.Value, error) {
	return json.Marshal(bi)
}

func (bi *UserInformation) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), bi)
}

func (bi UserInformation) Value() (driver.Value, error) {
	return json.Marshal(bi)
}

func (bi *SiteInformation) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), bi)
}

func (bi SiteInformation) Value() (driver.Value, error) {
	return json.Marshal(bi)
}

func (bi *Cookie) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), bi)
}

func (bi Cookie) Value() (driver.Value, error) {
	return json.Marshal(bi)
}

type Fire struct {
	ID                  uint               `gorm:"primarykey"`
	URL                 string             `gorm:"not null"`
	BrowserInformations BrowserInformation `json:"BrowserInformations"`
	UserInformations    UserInformation    `json:"UserInformations"`
	Cookies             Cookie             `json:"Cookies"`
	SiteInformations    SiteInformation    `json:"SiteInformations"`
	RandomID            string             `gorm:"not null"`
	gorm.Model
}

type Database struct {
	db *gorm.DB
}

type DatabaseInterface interface {
	Migrate(any) error
	FindAll(any) error
	Find(any, any) error
	Crete(any) error
	Update(any) error
	Delete(any) error
}
