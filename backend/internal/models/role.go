package models

import (
	"encoding/json"
	"errors"
	"log"
)



type Role int

const (
	AdminRole Role=iota
	AvocatRole
)

type RoleInterface interface {
	GetRoleName() Role
	SetUserID(id uint)
}

func (r Role) String() string{
	switch r{
		case AdminRole:
			return "Admin"
		case AvocatRole:
			return "Avocat"
		default:
			return "Unknown"
	}
}

func (r Role) MarshalJSON()([]byte,error){
	return json.Marshal(r.String())
}

func (r *Role) UnmarshalJSON(data []byte) error{
	var str string
	err:=json.Unmarshal(data,&str)
	if err!=nil{
		log.Println("Error unmarshalling role:",err)
	}
	switch str{
		case "Admin":
			*r=AdminRole
		case "Avocat":
			*r=AvocatRole
		default:
			return errors.New("unknown role")
	}
	return nil
}