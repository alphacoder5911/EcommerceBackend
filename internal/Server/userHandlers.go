package server

import (
	"github.com/alphacoder5911/EcommerceBackend/internal/dto"
	"github.com/alphacoder5911/EcommerceBackend/internal/utills"
	"github.com/gin-gonic/gin"
)

func (s *Server) getProfile(c *gin.Context){
	userId:=c.GetUint("user_id")
	

	user,err:=s.userService.GetProfile(userId)
	if err!=nil{
		utills.InternalServerErrorResponse(c,"Failed to fetch user profile",err)
		return 
	}

	utills.SuccessResponse(c,"Found",user)
}

func (s *Server) UpdateProfile(c *gin.Context){
	userId:=c.GetUint("user_id")
	

	var user *dto.UpdatedProfileRequest
	if err:= c.ShouldBindJSON(&user);err!=nil{
		utills.BadRequestResponse(c,"Invalid json data",err)
		return 
	}

	profile,err:=s.userService.UpdateProfile(userId,user)
	if err!=nil{
		utills.InternalServerErrorResponse(c,"Failed to update profile",err)
		return

	}

	utills.SuccessResponse(c,"Profile updated successfully",profile)

}