package Visitors

import "app.com/Models"

func UserTransformer(user Models.User) map[string]interface{} {
	newUser := make(map[string]interface{})
	newUser["username"] = user.Username
	newUser["email"] = user.Email
	newUser["token"] = user.Token
	return newUser
}

func UsersTransformer(users []Models.User) []map[string]interface{} {
	newUsers := make([]map[string]interface{}, 0)
	for _, user := range users {
		newUsers = append(newUsers, UserTransformer(user))
	}

	return newUsers
}
