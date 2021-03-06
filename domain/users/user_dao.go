package users

import (
	"bookstore/bookstore_users-api/utils/date_utils"
	"bookstore/bookstore_users-api/utils/errors"
	"fmt"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreted = result.DateCreted
	return nil

}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprint("email %s already resgistred", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreted = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
