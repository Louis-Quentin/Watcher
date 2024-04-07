package Models

import (
	_ "golang.org/x/oauth2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email                string `gorm:"unique"`
	Password             string
	Google_refresh_token string
	Google_access_token  string
	Services             []ServiceState `gorm:"foreignKey:ServiceId"`
}

type ServiceState struct {
	gorm.Model
	ServiceId uint
	Active    bool
	Name      string
	Areas     []AreaState `gorm:"foreignKey:AreaId"`
}

type AreaState struct {
	gorm.Model
	AreaId        uint
	Name          string
	Active        bool
	Action        bool
	Field         int
	Description_1 string
	Description_2 string
	Description_3 string
	FieldValue_1  string
	FieldValue_2  string
	FieldValue_3  string
	PrintName     string
}

type Watch struct {
	gorm.Model
	Name		string `gorm:"unique"`
	Price   	float64
	Range		int
	Url     	string
	Brand		string
	Available	bool
}
