// Package dto contains data transfer objects for authentication and user management.
package dto

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string  `json:"password"  binding:"required,min=8"`
	FirstName string  `json:"first_name" binding:"required"`
	SecondName string `json:"last_name" binding:"required"`
	Phone    string  `json:"phone" binding:"required"`
}

type LoginRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type AuthResponse struct {
	User UserResponse `json:"user"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	ID uint `json:"id"`
	Email string `json:"email"`
    FirstName string `json:"first_name"`
    SecondName string `json:"last_name"`
	Phone string `json:"phone"`
	Role string `json:"role"`
	IsActive bool `json:"is_active"`


}

type UpdatedProfileRequest struct{
	FirstName string `json:"first_name" binding:"required"`
	SecondName string `json:"last_name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}