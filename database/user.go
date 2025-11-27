package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

var users []User

func (u User) Store() User {

	if u.ID != 0 {
		return u
	}

	u.ID = len(users) + 1

	users = append(users, u)
	return u
}

// if possible to returen nill, then we can use pointer to user struct
func Find(email, pass string) *User {
	for _, user := range users {
		if user.Email == email && user.Password == pass {
			return &user
		}
	}

	return nil
}
