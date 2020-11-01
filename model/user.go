package model

// 生成mock测试文件，模拟数据库操作
// mockgen -source=./model/user.go -destination=./model/user_mock.go -package=model

type UserService interface {
	Get(username string) (User, error)
	Add(User) error
	Delete(ID int) error
	Update(User) error
}

type User struct {
	ID       int
	Name     string
	Password string
	Email    string
}

// 涉及数据库操作
func (u User) Get(username string) (User, error) {
	panic("db option")
}

func (u User) Add(user User) error {
	panic("db option")
}

func (u User) Delete(ID int) error {
	panic("db option")
}

func (u User) Update(user User) error {
	panic("db option")
}
