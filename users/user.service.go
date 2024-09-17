package users

func serviceGetUserById(id int64) *User {
	return getUserById(id)
}

func serviceGetAllUsers() []User {
	return getAllUsers()
}

func serviceAddUser(user User) User {
	return addUser(user)
}

func serviceUpdateUser(id int64, newUser User) *User {
	return updateUser(id, newUser)
}

func serviceRemoveUser(id int64) bool {
	return removeUser(id)
}
