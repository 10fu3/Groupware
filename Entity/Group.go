package Entity

import "time"

type Group struct {
	Name    string `json:"-" gorm:"column:name"`
	Uuid    string `json:"-" gorm:"primary_key" gorm:"column:uuid"`
	Users   []User `json:"-"`
	Roles   []Role `json:"-"`
	Private bool   `json:"-"`
}

type OutputGroup struct {
	Name  string `json:"name"`
	Uuid  string `json:"uuid"`
	Users []User `json:"users"`
}

type GroupInvite struct {
	Uuid      string
	GroupUuid string
	CreatedAt time.Time
	ClosingAt time.Time
	CreatedBy string
}

type Role struct {
	Uuid        string       `json:"uuid" gorm:"column:uuid" gorm:"primary_key"`
	Name        string       `json:"name" gorm:"column:name"`
	Color       string       `json:"color" gorm:"column:color"`
	Permissions []Permission `json:"has_permissions"`
	Users       []User       `json:"users"`
}

type Permission struct {
	Name       string `json:"name" gorm:"column:name"`
	Permission bool   `json:"has_permission"`
}
