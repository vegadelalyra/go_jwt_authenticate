package models

type LoginRequest struct {
	UserName   string `json:"username" form:"username" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	RememberMe bool   `json:"rememberme" form:"rememberme"`
}
