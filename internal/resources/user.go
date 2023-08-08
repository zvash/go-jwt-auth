package resources

import "github.com/zvash/go-jwt-auth/internal/database"

type User struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) FillWithDbUser(dbUser database.User) {
	u.ID = dbUser.ID
	u.Name = dbUser.Name
	u.Email = dbUser.Email
}
