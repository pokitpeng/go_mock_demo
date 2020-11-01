package api

import (
	"testing"

	"github.com/golang/mock/gomock"
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/assert"

	errs "go_mock_demo/error"
	"go_mock_demo/model"
)

func TestLogin(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	// 用户名密码正确
	tom := model.User{ID: 1, Name: "tom", Password: "123456", Email: "tom@email.com"}
	mockUser := model.NewMockUserService(ctl)
	mockUser.EXPECT().Get("tom").Return(tom, nil)
	assert.Equal(t, nil, Login(mockUser, "tom", "123456"))

	// 用户名密码错误
	tom = model.User{ID: 1, Name: "tom", Password: "123456", Email: "tom@email.com"}
	mockUser = model.NewMockUserService(ctl)
	mockUser.EXPECT().Get("tom").Return(tom, nil)
	assert.Equal(t, errs.Trans(errs.ErrUsernameOrPassword), Login(mockUser, "tom", "1234567").Error())
}
