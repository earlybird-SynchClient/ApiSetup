package market

import "github.com/jinzhu/gorm"

type Market struct {
	gorm.Model
	Name	string		`json:"marketname"`
}
