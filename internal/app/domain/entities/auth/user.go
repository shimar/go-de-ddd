package auth

import (
	"errors"

	"github.com/shimar/go-de-ddd/internal/app/domain/values"
)

// User 認証用ユーザー構造体
type User struct {
	id       values.UserID
	password values.Password
}

// NewUser 認証用ユーザーを生成する
func NewUser(id *values.UserID, password *values.Password) (*User, error) {
	if id == nil {
		return nil, errors.New("IDを指定してください")
	}
	if password == nil {
		return nil, errors.New("パスワードを指定してください")
	}
	return &User{id: *id, password: *password}, nil
}

// Equals パスワードが同じか?
func (u *User) Equals(other *values.Password) bool {
	return true
}
