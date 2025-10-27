package handlers

import (
	"net/http"
	"strconv"
	"github.com/houssybadr/lawyermanagement/backend/internal/dto"
	"github.com/houssybadr/lawyermanagement/backend/internal/models"
	"github.com/houssybadr/lawyermanagement/backend/internal/repository"
	"github.com/houssybadr/lawyermanagement/backend/internal/utils"
	"github.com/houssybadr/lawyermanagement/backend/internal/webhooks"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
 


type AuthHandler struct{
	authRepository repository.AuthRepository
}

func (a *AuthHandler) SetRepository(repo repository.AuthRepository){
	a.authRepository=repo
}

func (a *AuthHandler) SignUpAdmin(c *gin.Context) {
	var req dto.AdminSignupRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	user:=req.User
	admin:=req.Admin
	hashed,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err!=nil{
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	user.Password=string(hashed)
	user.Role=models.AdminRole
	if err:=a.authRepository.SignUp(&user,&admin);err!=nil{
		c.JSON(404,gin.H{"error":err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated,admin)
}

func (a *AuthHandler) SignUpAvocat(c *gin.Context){
	var req dto.AvocatSignupRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	user:=req.User
	avocat:=req.Avocat
	hashed,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err!=nil{
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	user.Password=string(hashed)
	user.Role=models.AvocatRole
	if err:=a.authRepository.SignUp(&user,&avocat);err!=nil{
		c.JSON(404,gin.H{"error":err.Error()})
		return
	}
	go webhooks.CreatedAvocatN8nWebhook(avocat)
	c.JSON(http.StatusCreated,avocat)
}

func (a *AuthHandler) SignIn(c *gin.Context){
	var req dto.SignInRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	var user models.User
	if err:=a.authRepository.GetUserByEmail(req.Email,&user);err!=nil{
		c.JSON(404,gin.H{"error":err.Error()})
		return
	}
	 if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
	switch user.Role{
	case models.AdminRole:
		var admin models.Admin
		if err:=a.authRepository.GetActorByUserID(&admin,user.Id);err!=nil{
			c.JSON(404,gin.H{"error":err.Error()})
			return
		}
		jwtKey:=[]byte(utils.GetJwtSecret())
		tokenString,err:=utils.GetJwtSignedToken(user.Email,user.Id,user.Role.String(),jwtKey)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		var response = dto.SuccessfulAdminSignInResponse{
			Role: user.Role,
			Admin: admin,
			Token: tokenString,
		}
		c.JSON(http.StatusOK,response)
	case models.AvocatRole:
		var avocat models.Avocat
		if err:=a.authRepository.GetActorByUserID(&avocat,user.Id);err!=nil{
			c.JSON(404,gin.H{"error":err.Error()})
			return
		}
		jwtKey:=[]byte(utils.GetJwtSecret())
		tokenString,err:=utils.GetJwtSignedToken(user.Email,user.Id,user.Role.String(),jwtKey)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		var response = dto.SuccessfulAvocatSignInResponse{
			Role: user.Role,
			Avocat: avocat,
			Token: tokenString,
		}
		c.JSON(http.StatusOK,response)
	default:
		c.JSON(400,gin.H{"error":"invalid role"})
		return
	}
}

func (a *AuthHandler)ChangePassword(c *gin.Context){
	var passwordChangeRequest dto.PasswordChangeRequest
	if err:=c.ShouldBindJSON(&passwordChangeRequest);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	userId,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{ 
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID"})
	}
	var oldHashedPassword string
	if err:=a.authRepository.GetPassword(uint(userId),&oldHashedPassword);err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound,gin.H{"error":"User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	err=bcrypt.CompareHashAndPassword([]byte(oldHashedPassword),[]byte(passwordChangeRequest.OldPassword))
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Old password is incorrect"})
		return
	}
	newHashedPassword,err:=bcrypt.GenerateFromPassword([]byte(passwordChangeRequest.NewPassword),bcrypt.DefaultCost)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	
	if err=a.authRepository.ChangePassword(uint(userId),string(newHashedPassword));err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"Password changed successfully"})
}

func (a *AuthHandler)DeleteUser(c *gin.Context){
	userId,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid User ID"})
		return
	}
	if err=a.authRepository.DeleteUser(uint(userId));err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound,gin.H{"error":"User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"User deleted successfully"})
}