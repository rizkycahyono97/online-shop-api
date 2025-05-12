package web

type UserRequest struct {
	Name     string `json:"name" binding:"required,max=255"`
	Email    string `json:"email" binding:"required,max=255"`
	Password string `json:"password" binding:"required,max=255"`
	Role     string `json:"role" binding:"required,default=customer"`
}
