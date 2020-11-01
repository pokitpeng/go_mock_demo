package api

import (
	errs "go_mock_demo/error"
	"go_mock_demo/model"
)

// Login 这里传入了一个UserService，执行测试的时候传MockUserService，
// 生产环境的时候传入UserService,这样就做到了业务代码和环境依赖的解耦，
// 就算没有数据库，也能进行业务测试
func Login(ui model.UserService, username, password string) error {
	user, err := ui.Get(username)
	if err != nil {
		return err
	}

	if user.Password == password {
		return nil
	} else {
		return errs.NewWithCode(errs.ErrUsernameOrPassword)
	}
}
