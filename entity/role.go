package entity

import "github.com/google/uuid"

type Role struct {
	Id          uuid.UUID    `gorm:"primaryKey;column:id;type:varchar(36)"`
	Name        string       `gorm:"unique;column:name;type:varchar(30)"`
	DisplayName string       `gorm:"column:display_name;type:varchar(50);default:null"`
	Description string       `gorm:"column:description;type:varchar(255);default:null"`
	Users       []User       `gorm:"many2many:role_user;"`
	Permissions []Permission `gorm:"many2many:permission_role;"`
}
