package mock

import "errors"

// go 里推荐使用 go-generate 工具，以注释形式直接把 mock 命令写到对应接口这里，go 会把 //go:generate 后根据内容作为命令执行
// 在根目录下执行go generate ./...会生成所有mock文件
//
//go:generate mockgen -package mock -destination mock_user.go learning-go/test/framework/mock IUser
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
