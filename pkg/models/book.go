package models

import (
	"github.com/akgmage/go-sequel/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
type Book struct {
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`

}

