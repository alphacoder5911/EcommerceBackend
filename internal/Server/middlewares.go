package server

import (
	"strings"

	"github.com/alphacoder5911/EcommerceBackend/internal/models"
	utills "github.com/alphacoder5911/EcommerceBackend/internal/utills"
	"github.com/gin-gonic/gin"
)

func (s *Server) authMiddleware() gin.HandlerFunc{

	return func(ctx *gin.Context) {
		authHeader:=ctx.GetHeader("Authorization")
		if authHeader==""{
			utills.UnauthorizedResponse(ctx,"Authorization header required")
			ctx.Abort()
			return
		}

		tokenParts:=strings.Split(authHeader," ")
		if len(tokenParts)!=2 || tokenParts[0]!="Bearer"{
			utills.UnauthorizedResponse(ctx,"Invalid authorization token format")
			ctx.Abort()
			return 
		}

		claims,err:=utills.ValidateToken(tokenParts[1],s.Config.JWT.Secret)
		if err!=nil{
			utills.UnauthorizedResponse(ctx,"Invalid token")
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