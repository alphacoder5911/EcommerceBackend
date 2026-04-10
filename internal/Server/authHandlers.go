package server

import (

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

	
	response,err:=s.authService.Register(&req)
	if err!=nil{
		utills.BadRequestResponse(c,"Registration failed",err)
		return 	
	}	

	c.SetCookie("access_token",response.AccessToken,int(s.Config.JWT.ExpiresIn),"/","",false,true)
	c.SetCookie("refresh_token",response.RefreshToken,int(s.Config.JWT.RefreshTokenExpires),"/","",false,true)

	utills.CreatedResponse(c,"User registered successfully",response)
	
}

func (s *Server) Login(c *gin.Context){
	var req dto.LoginRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		utills.BadRequestResponse(c,"invalid credentials",err)
		return
	}

	// authService:=services.NewAuthService(s.db,s.Config)

	response,err:=s.authService.Login(&req)
	if err!=nil{
		utills.BadRequestResponse(c,"Login failed ",err)
		return
	}
	c.SetCookie("access_token",response.AccessToken,int(s.Config.JWT.ExpiresIn),"/","",false,true)
	c.SetCookie("refresh_token",response.RefreshToken,int(s.Config.JWT.RefreshTokenExpires),"/","",false,true)

	utills.SuccessResponse(c,"User logged in sussfully",response)

}

func (s *Server) RefreshToken(c *gin.Context){
	var req dto.RefreshTokenRequest
	// if err:=c.ShouldBindJSON(&req);err!=nil{
	// 	utills.BadRequestResponse(c,"Invalid creds",err)
	// 	return
	// }                     
	
	refreshToken,err:=c.Cookie("refresh_token")
	if err!=nil{
		utills.InternalServerErrorResponse(c,"Couldnt Fetch refresh token ",err)
		return
	}

	req.RefreshToken=refreshToken
	
	response,err:=s.authService.RefreshToken(&req)
	if err!=nil{
		utills.UnauthorizedResponse(c,"Couldnt refresh token")
		return
	}

	utills.SuccessResponse(c,"Token refreshed successfully",response)
}

func (s *Server) Logout(c *gin.Context){
	var req  dto.RefreshTokenRequest
	// if err:=c.ShouldBindJSON(&req);err!=nil{
	// 	utills.BadRequestResponse(c,"INvalid request",err)
	// 	return 
	// }
	refToken,err:=c.Cookie("refresh_token")
	if err!=nil{
		utills.UnauthorizedResponse(c,"Refresh token not found")
		return 
	}
	req.RefreshToken=refToken

	if err:=s.authService.Logout(req.RefreshToken);err!=nil{
		utills.InternalServerErrorResponse(c,"Logout failed",err)
		return
	}
	c.SetCookie("access_token","",-1,"/","",false,true)
	c.SetCookie("refresh_token","",-1,"/","",false,true)

	utills.SuccessResponse(c,"Logout Successfulll",nil)
}

func (s *Server) profile(c *gin.Context){
	utills.SuccessResponse(c,"Profile  Fetched successfully",nil)
}