package repository

import (
	"gorm.io/gorm"
	"github.com/houssybadr/lawyermanagement/backend/internal/models"
)

type AuthRepository struct{
	db *gorm.DB
}

func (a* AuthRepository) SetDB(db *gorm.DB){
	a.db=db
}


func (a AuthRepository) SignUp(user *models.User,role models.RoleInterface) error{
	err:=a.db.Transaction(func(tx *gorm.DB)error{
		if err:=tx.Create(user).Error;err!=nil{
			return err
		}
		role.SetUserID(user.Id)
		if err:=tx.Create(role).Error;err!=nil{
			return err
		}
		return nil
	})
	if err!=nil{
		return err
	}
	return nil
}

func (a AuthRepository) GetUserByEmail(email string,user *models.User) error{
	err:=a.db.Where("email=?",email).First(user).Error
	if err!=nil{
		return err
	}
	return nil
}


func (a AuthRepository)GetActorByUserID(item models.RoleInterface,id uint) error{
	if err:=a.db.Where("user_id=?",id).First(&item).Error;err!=nil{
		return err
	}
	return nil
}

func (a AuthRepository)GetPassword(id uint,password *string)(error){
	var user models.User
	if err:=a.db.First(&user,id).Error;err!=nil{
		return err
	}
	*password=user.Password
	return nil
}

func (a AuthRepository)ChangePassword(id uint,hashedPassword string)error{
	var item models.User
	if err:=a.db.Model(&item).Where("id=?",id).Update("password",hashedPassword).Error;err!=nil{
		return err
	}
	return nil
}

func (a AuthRepository)DeleteUser(id uint) error{
	var user models.User
	if err:=a.db.Delete(&user,id).Error;err!=nil{
		return err
	}
	return nil
}