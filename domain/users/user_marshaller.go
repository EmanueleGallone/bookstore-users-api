package users

import "encoding/json"

type PublicUser struct {
	Id int64 `json:"id"`
	//FirstName   string `json:"first_name"`
	//LastName    string `json:"last_name"`
	//Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	//Password    string `json:"password"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	//Password    string `json:"password"`
}

func (users Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPublic)
	}
	return result
}

func (u User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          u.Id,
			DateCreated: u.DateCreated,
			Status:      u.Status,
		}
	}
	userJson, _ := json.Marshal(u) // this approach can be used only if you have the same json structure
	var privateUser PrivateUser
	_ = json.Unmarshal(userJson, &privateUser)

	return privateUser

}
