package main

func RepoFindUser(id int) User {
	user := User{}
	DB.Get(&user, "SELECT * FROM users WHERE id = $1 LIMIT 1", id)
	return user
}

func RepoFindUserByEmail(email string) User {
	user := User{}
	DB.Get(&user, "SELECT * FROM users WHERE email = $1 LIMIT 1", email)
	return user
}

func RepoInsertUser(user User) User {
	password := CreateHash(user.Password)
	err := DB.QueryRow("INSERT INTO users(email, password, created_at) VALUES($1, $2, now()) RETURNING id, created_at", user.Email, password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		panic(err)
	}
	user.Password = ""
	return user
}
