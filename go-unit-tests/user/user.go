package user

type User struct {
	ID   int
	Name string
}

type Repository interface {
	FindByID(id int) (*User, error)
}
