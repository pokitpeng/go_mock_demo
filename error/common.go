package error

var errorMaps = map[int]string{
	BaseErrUnknown:    "unknown error",
	BaseErrHaveNoPerm: "have no permission",
	BaseErrTimeout:    "timeout",
}

// error define
const (
	BaseErrUnknown = iota + 1100
	BaseErrHaveNoPerm
	BaseErrTimeout
)

func registerMap(maps map[int]string) {
	for k, v := range maps {
		errorMaps[k] = v
	}
}

// Trans code to message
func Trans(code int) string {
	if errStr, ok := errorMaps[code]; ok {
		return errStr
	}
	return errorMaps[BaseErrUnknown]
}

// TransError code to message
func TransError(err error) string {
	if e, ok := err.(MyError); ok {
		if errStr, ok := errorMaps[e.Code()]; ok {
			return errStr
		}
	}
	return err.Error()
}
