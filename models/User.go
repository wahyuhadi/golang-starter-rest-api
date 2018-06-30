/**
	Rahmat wahyu hadi
**/

package models

import (
	"errors"
	"time"

	"../config"
)

// menyesuaikan databases
type User struct {
	ID int `json:"id" gorm:"primary_key"`

	FirstName string     `json:"firstname, omitempty" gorm:"not null; type:varchar(100)"`
	LastName  string     `json:"lastname, omitempty" gorm:"not null; type:varchar(100)"`
	Email     string     `json:"email, omitempty" gorm:"not null; type:varchar(100)"`
	CreatedAt *time.Time `json:"createdAt, omitempty"`
	UpdatedAt *time.Time `json:"updatedAt, omitempty"`
	DeletedAt *time.Time `json:"deletedAt, omitempty" sql:"index"`
}

func (User) TableName() string {
	return "users" // nama tables
}

func (u *User) FetchAll() []User {
	db := config.GetDatabaseConnection()

	var users []User
	db.Find(&users)

	return users
}

func (u *User) FetchById() error {
	db := config.GetDatabaseConnection()

	if err := db.Where("id = ?", u.ID).Find(&u).Error; err != nil {
		return errors.New("Could not find the user")
	}

	return nil
}

func (u *User) Save() error {
	db := config.GetDatabaseConnection()

	if db.NewRecord(u) {
		if err := db.Create(&u).Error; err != nil {
			return errors.New("Could not create user")
		}
	} else {
		if err := db.Save(&u).Error; err != nil {
			return errors.New("Could not update user")
		}
	}

	return nil
}

func (u *User) Delete() error {
	db := config.GetDatabaseConnection()

	if err := db.Delete(&u).Error; err != nil {
		return errors.New("Could not find the user")
	}

	return nil
}
