package server

import (
	services "github.com/alphacoder5911/EcommerceBackend/internal/Services"
	"github.com/alphacoder5911/EcommerceBackend/internal/dto"
	"github.com/alphacoder5911/EcommerceBackend/internal/utills"
	"github.com/gin-gonic/gin"
)

func (s *Server) Register(c *gin.Context){
	var req dto.RegisterRequest
	if err:= c.ShouldBindJSON(&req);err!=nil{
		utills.BadRequestResponse(c,"Invalid requres data",err)
		return
	}

	authService:= services.NewAuthService(s.db,s.Config)
	response,err:=authService.Register(&req)
	if err!=nil{
		utills.BadRequestResponse(c,"Registration failed",err)
		return 
	}	

	utills.CreatedResponse(c,"User registered successfully",response)
	
}

func (s *Server) Login(c *gin.Context){
	var req dto.LoginRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		utills.BadRequestResponse(c,"invalid credentials",err)
		return
	}

	authService:=services.NewAuthService(s.db,s.Config)

	response,err:=authService.Login(&req)
	if err!=nil{
		utills.BadRequestResponse(c,"Login failed ",err)
		return
	}

	utills.SuccessResponse(c,"User logged in sussfully",response)

}

func (s *Server) RefreshToken(c *gin.Context){
	var req dto.RefreshTokenRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		utills.BadRequestResponse(c,"Invalid creds",err)
		return
	}
	
	authService:=services.NewAuthService(s.db,s.Config)
	response,err:=authService.RefreshToken(&req)
	if err!=nil{
		utills.UnauthorizedResponse(c,"Couldnt refresh token")
		return
	}

	utills.SuccessResponse(c,"Token refreshed successfully",response)
}

func (s *Server) Logout(c *gin.Context){
	var req  dto.RefreshTokenRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		utills.BadRequestResponse(c,"INvalid request",err)
		return 
	}

	authService:=services.NewAuthService(s.db,s.Config)
	if err:=authService.Logout(req.RefreshToken);err!=nil{
		utills.InternalServerErrorResponse(c,"Logout failed",err)
		return
	}

	utills.SuccessResponse(c,"Logout Successfulll",nil)
}