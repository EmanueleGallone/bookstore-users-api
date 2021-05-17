package users

//dto stands for data transfer object.

import (
	"bookstore-users-api/util/errors"
	"regexp"
	"unicode"
)

const (
	StatusActive = "active"
	//StatusPending = "pending"
	//StatusInactive = "inactive"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User // alias used in marshaller

/*
validators should always be a responsibility of the package they are validating!
follow this example.
*/
func (user User) Validate() *errors.RestError {
	isValid := isEmailValid(user.Email)
	isPasswordValid := isPasswordValid(user.Password)

	if !isValid {
		return errors.NewBadRequestError("Invalid Email Address")
	}
	if !isPasswordValid {
		return errors.NewBadRequestError("Invalid Password")
	}

	return nil
}

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return pattern.MatchString(e)
}

// Password validates plain password against the rules defined below.
//
// upp: at least one upper case letter.
// low: at least one lower case letter.
// num: at least one digit.
// sym: at least one special character.
// No empty string or whitespace.
func isPasswordValid(pass string) bool {
	const minLength = 2
	var (
		upp, low, num, sym bool
		tot                uint8
	)
	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < minLength {
		return false
	}

	return true
}
