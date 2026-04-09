package services

import (
	"errors"
	"fmt"

	"time"

	"github.com/alphacoder5911/EcommerceBackend/internal/config"
	"github.com/alphacoder5911/EcommerceBackend/internal/dto"

	//"github.com/alphacoder5911/EcommerceBackend/internal/logger"
	"github.com/alphacoder5911/EcommerceBackend/internal/models"
	"github.com/alphacoder5911/EcommerceBackend/internal/utills"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
	config *config.Config
}

func NewAuthService(db *gorm.DB,config *config.Config) *AuthService{
	return &AuthService{
		db: db,
		config: config,
	}
}

func (s *AuthService) Register(req *dto.RegisterRequest)(*dto.AuthResponse,error){
	// Checking if user already exists 
	var existingUser int64

	if err:= s.db.Where("email = ?",req.Email).First(&existingUser).Error;err!=nil{
		return nil,errors.New("User not found")
	}

	hashedPwd,err:=utills.HashPassword(req.Password)
	if err!=nil{
		return nil,errors.New("Failed to hash password")
	}

	user:=models.User{
		Email: req.Email,
		Password: hashedPwd,
		FirstName: req.FirstName,
		SecondName: req.SecondName,
		Phone: req.Phone,
		Role: models.UserRoleCustomer,
	}

	if err:=s.db.Create(&user).Error;err!=nil{
		return nil,err

	}

	cart:=models.Cart{UserID: user.ID}
	if err:=s.db.Create(&cart).Error; err!=nil{
		fmt.Println("Unable to create cart")

	}
	return s.generateAUthResponse(&user)
}


func (s *AuthService) Login(req *dto.LoginRequest)(*dto.AuthResponse,error){
	var user models.User
	if err:= s.db.Where("email=? AND is_active=?",req.Email,true).Error;err!=nil{
		return nil,errors.New("INvalid credentials ")
	}

	if !utills.CheckPasswordHash(req.Password,user.Password){
		return nil,errors.New("Invalid credentials")
	}

	return s.generateAUthResponse(&user)
}

func (s *AuthService) RefreshToken(req *dto.RefreshTokenRequest)(*dto.AuthResponse,error){
	claims,err:=utills.ValidateToken(req.RefreshToken,s.config.JWT.Secret)
	if err!=nil{
		return nil,errors.New("Invalid refresh token")
	}

	var refreshToken models.RefreshToken
	if err:= s.db.Where("token = ? AND expires_at > ?",req.RefreshToken,time.Now()).First(&refreshToken).Error;err!=nil{
		return nil,errors.New("Refresh token not found or expired")
	}

	var user models.User
	if err:=s.db.First(&user,claims.UserId).Error;err!=nil{
		return nil,errors.New("User not found")
	}

	s.db.Delete(&refreshToken)
 	return 	s.generateAUthResponse(&user)
}

func (s *AuthService) Logout(refreshToken string) error{
	return s.db.Where("token = ?",refreshToken).Delete(&models.RefreshToken{}).Error
}

func (s *AuthService) generateAUthResponse(user *models.User) (*dto.AuthResponse,error){
	accessToken,refreshToken,err:=utills.GenerateTokenPair(
		&s.config.JWT,
		user.ID,
		user.Email,
		string(user.Role),
	)
	if err!=nil{
		return nil,err
	}

	s.db.Create(&refreshToken)
	
	return &dto.AuthResponse{
		User: dto.UserResponse{
			ID: user.ID,
			Email: user.Email,
			FirstName: user.FirstName,
			SecondName: user.SecondName,
			Phone: user.Phone,
			Role: string(user.Role),
			IsActive: user.IsActive,
		},
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	},nil
}