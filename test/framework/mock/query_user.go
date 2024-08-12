package mock

import "errors"

type IUser interface {
	Get(id string) (User, error)
}
type User struct {
	Username string
	Password string
}

var ErrEmptyID = errors.New("id is empty")

func QueryUser(db IUser, id string) (User, error) {
	if id == "" {
		return User{}, ErrEmptyID
	}
	return db.Get(id)
}
