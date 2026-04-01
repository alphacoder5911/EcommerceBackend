// This Package dto contains data transfer objects for orders and cart management.
package dto

type AddToCardRequest struct{
	ProductID uint `json:"product_id" binding:"required"`
Quantity int `json:"quantity" binding:"required,gt=0"`
}

type UpdateCartItemRequest struct{
Quantity int `json:"quantity" binding:"required,gt=0"`
}

type CartResponse struct{
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	Items []CartItemResponse `json:"items"`
	Total float64 `json:"total"`
}

type CartItemResponse struct{
	ID uint `json:"id"`
	ProductID uint `json:"product_id"`
	Quantity int `json:"quantity"`
	Subtotal float64 `json:"subtotal"`
}