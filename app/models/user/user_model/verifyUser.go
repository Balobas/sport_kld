package user_model

func VerifyUser(user User) error {
	_, err := getUserByLoginAndPassword(user.Login, user.Password)
	return err
}
