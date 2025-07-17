package model

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" validate:"required,min=3,max=50"`
	Email string `json:"email" vaildate:"required,email"`
}
