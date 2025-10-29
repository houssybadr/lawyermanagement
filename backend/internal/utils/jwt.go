package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetJwtSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetJwtSignedToken(email string,user_id uint,user_role string,jwtkey []byte) (string,error){
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":user_id,
		"user_role":user_role,
		"email":email,
		"exp":time.Now().Add(time.Hour*24).Unix(),
	})
	stringToken,err:=token.SignedString(jwtkey)
	if err!=nil{
		return "",err
	}
	return stringToken,nil
}

func ParseJwtToken(signedString string,jwtkey []byte) (*jwt.Token,error){
	token,err:=jwt.Parse(signedString,func(token *jwt.Token)(interface{},error){
		if token.Method!=jwt.SigningMethodHS256{
			return nil,jwt.ErrSignatureInvalid
		}
		return jwtkey,nil
	})
	if err!=nil{
		return nil,err
	}
	return token,nil
}