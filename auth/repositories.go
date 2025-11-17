package auth

func InsertUser(user *User) {
	DB.Insert(*user)
}
