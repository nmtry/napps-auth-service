package auth

func InsertUser(user *User) {
	DB.Insert(*user)
}

func IsEmailUsed(email string) bool {
	_, ok := DB.FindUserByEmail(email)
	return ok
}
