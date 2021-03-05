package user_model

func VerifyUser(user User) error {
	_, err := getUserByLoginAndPassword(user.Login, user.Password)
	if err != nil {
		return err
	}
	return nil
}
