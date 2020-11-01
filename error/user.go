package error

var taskErrorMaps = map[int]string{
	ErrUsernameOrPassword: "username or password error",
}

// code define
const (
	ErrUsernameOrPassword = 12000 + iota
)

func init() {
	registerMap(taskErrorMaps)
}
