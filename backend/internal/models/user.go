package models

import (
	"strconv"
)

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email" `
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Role     Role   `gorm:"type:int ;not null" json:"role" `
}


func (u *User) IsEqual(other User)bool{
	return u.Email==other.Email && u.Password==other.Password && u.Role==other.Role
}
func (u *User) ToString()string{
	return "User{Id:"+strconv.FormatUint(uint64(u.Id),10)+",Email:"+u.Email+",Role:"+strconv.FormatUint(uint64(u.Role),10)+"}"
}