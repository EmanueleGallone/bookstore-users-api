package services

import (
	"bookstore-users-api/domain/users"
	"bookstore-users-api/util/crypto_util"
	"bookstore-users-api/util/date_util"
	"bookstore-users-api/util/errors"
)

var (
	UserServices UserServiceInterface = &UserService{}
)

type UserService struct{} //struct implementing the interface

type UserServiceInterface interface { //the interface is important also for testing; this way you can mock tests
	GetUser(userID int64) (*users.User, *errors.RestError)
	CreateUser(user *users.User) (*users.User, *errors.RestError)
	UpdateUser(isPartial *bool, user *users.User) (*users.User, *errors.RestError)
	DeleteUser(user *users.User) (*users.User, *errors.RestError)
	SearchUsers(status string) (users.Users, *errors.RestError)
}

func (u *UserService) GetUser(userID int64) (*users.User, *errors.RestError) {
	user := users.User{
		Id: userID,
	}

	if err := user.GetByID(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserService) CreateUser(user *users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_util.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_util.GetSHA256Hash(user.Password)

	if err := user.SaveToDB(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) UpdateUser(isPartial *bool, user *users.User) (*users.User, *errors.RestError) {
	currentUser, err := u.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if *isPartial {
		if user.Email != "" {
			currentUser.Email = user.Email
		}
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	}

	currentUser.UpdateUser()

	return nil, nil
}

func (u *UserService) DeleteUser(user *users.User) (*users.User, *errors.RestError) {
	if err := user.DeleteUser(); err != nil {
		return nil, err
	}

	return user, nil
}

/*
returns a slice of users that matched the given status
*/
func (u *UserService) SearchUsers(status string) (users.Users, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus("Active")
}
