package resources

import "github.com/zvash/go-jwt-auth/internal/database"

type User struct {
	ID    int32  `responsemaker:"id"`
	Name  string `responsemaker:"name"`
	Email string `responsemaker:"email"`
}

func (u *User) FillWithDbUser(dbUser database.User) {
	u.ID = dbUser.ID
	u.Name = dbUser.Name
	u.Email = dbUser.Email
}
