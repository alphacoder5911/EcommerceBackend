package server

import (


	"github.com/alphacoder5911/EcommerceBackend/internal/models"
	utills "github.com/alphacoder5911/EcommerceBackend/internal/utills"
	"github.com/gin-gonic/gin"
)

func (s *Server) authMiddleware() gin.HandlerFunc{

	return func(ctx *gin.Context) {

		cookieToken,err:=ctx.Cookie("access_token")
		if err!=nil{
			utills.UnauthorizedResponse(ctx,"User unauthorized")
			ctx.Abort()
			return
		}

		claims,err:=utills.ValidateToken(cookieToken,s.Config.JWT.Secret)
		if err!=nil{
			utills.UnauthorizedResponse(ctx,"token invalid")
			ctx.Abort()
			return
		}


		ctx.Set("user_id",claims.UserId)
		ctx.Set("email", claims.Email)
		ctx.Set("user_role",claims.Role)

		ctx.Next()
	}
}

func (s *Server) adminMiddleware() gin.HandlerFunc{
	return func (c *gin.Context){
		role,exists:=c.Get("user_role")
		if !exists || role!=string(models.UserRoleAdmin){
			utills.ForbiddenResponse(c,"Forbidden")
			c.Abort()
			return
		}

		c.Next()
	}
}