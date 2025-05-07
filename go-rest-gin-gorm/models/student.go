package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model //Autmatically adds ID and timestamps

	Name   string `json:"name"`
	Number string `json:"Number"`
	SSID   string `json:"SSID"`
}

var Students []Student
