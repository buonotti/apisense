package db

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/buonotti/apisense/api/jwt"
	"github.com/buonotti/apisense/errors"
)

type UserData struct {
	Token string
	Uid   string
}

func LoginUser(username string, password string) (UserData, error) {
	rows := db.QueryRow("SELECT * FROM users WHERE username = ? AND password = ? AND enabled = 1", username, password)

	if rows == nil {
		return UserData{}, errors.LoginError.New("invalid credentials")
	}

	token, err := jwt.Service().GenerateToken(username)
	if err != nil {
		return UserData{}, err
	}

	return UserData{Token: token, Uid: username}, nil
}

func RegisterUser(username string, password string) (UserData, error) {
	rows, err := db.Query("SELECT * FROM users WHERE username = ?", username)

	if rows.Next() {
		return UserData{}, errors.UserAlreadyExistsError.New("user already exists")
	}

	passwordHash := hashPassword(password)

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, passwordHash)
	if err != nil {
		return UserData{}, err
	}

	token, err := jwt.Service().GenerateToken(username)
	if err != nil {
		return UserData{}, err
	}

	return UserData{Token: token, Uid: username}, nil
}

func EnableUser(uid string) error {
	_, err := db.Exec("UPDATE users SET enabled = 1 WHERE username = ?", uid)
	return err
}

func DisableUser(uid string) error {
	_, err := db.Exec("UPDATE users SET enabled = 0 WHERE username = ?", uid)
	return err
}

func DeleteUser(uid string) error {
	_, err := db.Exec("DELETE FROM users WHERE username = ?", uid)
	return err
}

type User struct {
	Username string
	Enabled  bool
}

func ListUsers() ([]User, error) {
	rows, err := db.Query("SELECT username, enabled FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var username string
		var enabled bool
		err := rows.Scan(&username, &enabled)
		if err != nil {
			return nil, err
		}
		users = append(users, User{Username: username, Enabled: enabled})
	}

	return users, nil
}

func IsUserEnabled(uid string) bool {
	rows := db.QueryRow("SELECT enabled FROM users WHERE username = ?", uid)

	if rows == nil {
		return false
	}

	var enabled bool

	err := rows.Scan(&enabled)
	if err != nil {
		return false
	}

	return enabled
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
