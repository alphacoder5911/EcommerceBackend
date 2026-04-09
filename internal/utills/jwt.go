package utills

import (
	"errors"
	"time"

	"github.com/alphacoder5911/EcommerceBackend/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	UserId uint `json:"user_id"`
	Email string `json:"email"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateTokenPair(cfg *config.JWTConfig,userID uint, email,role string)(AccessToken ,RefreshToken string ,err error){

	AccessClaims:=&Claims{
		UserId:userID,
		Email: email,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.ExpiresIn)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	AcToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,AccessClaims)
	AccessToken,err=AcToken.SignedString([]byte(cfg.Secret))
	if err!=nil{
		return "", "", err
	}


	//Refresh token 
	RefreshClaims:=&Claims{
		UserId:userID,
		Email: email,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.ExpiresIn)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	ReToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,RefreshClaims)
	RefreshToken,err=ReToken.SignedString([]byte(cfg.Secret))
	if err!=nil{
		return "", "", err
	}

	


	return AccessToken,RefreshToken,nil

}	


func ValidateToken(tokenString string , secret string)(*Claims,error){
	token,err:=jwt.ParseWithClaims(tokenString,&Claims{},func(token *jwt.Token)(interface{},error){
	return []byte(secret),nil
	})

	if err!=nil{
		return nil, err
	}

	if claims,ok:=token.Claims.(*Claims);ok && token.Valid{
		return claims,nil
	}
	return nil, errors.New("Invalid token")
}