package repo

import "errors"

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (ur *userRepo) Create(user User) (*User, error) {

	if user.ID != 0 {
		return &user, nil
	}

	user.ID = len(ur.users) + 1

	ur.users = append(ur.users, user)
	return &user, nil
}

// if possible to returen nill, then we can use pointer to user struct
func (ur *userRepo) Find(email, pass string) (*User, error) {
	for _, user := range ur.users {
		if user.Email == email && user.Password == pass {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}
