package models

import (
	"gorm.io/gorm"
	"study.co.jp/go-rest-gin-gorm/validator"
)

type Student struct {
	gorm.Model //Autmatically adds ID and timestamps

	Name   string `json:"name" validate:"required"`
	Number string `json:"Number" validate:"len=9,custom"`
	SSID   string `json:"SSID" validate:"len=11"`
}

func Validate(student *Student) error {
	return validator.ValidateStruct(student)
}
