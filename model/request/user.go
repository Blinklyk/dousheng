package request

type RegisterRequest struct {
	Username string `json:"username" gorm:"not null; comment:username for register;" form:"username"`
	Password string `json:"password" gorm:"not null; comment:password for register" form:"password"`
}

type LoginRequest struct {
	Username string `json:"username" gorm:"not null; comment:username for register;" form:"username"`
	Password string `json:"password" gorm:"not null; comment:password for register" form:"password"`
}
