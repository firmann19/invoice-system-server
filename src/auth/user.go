package auth

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

var Users = []User{
	{
		ID:       1,
		Username: "admin",
		Password: "admin123",
		Role:     "admin",
	},
	{
		ID:       2,
		Username: "kerani",
		Password: "kerani123",
		Role:     "kerani",
	},
}