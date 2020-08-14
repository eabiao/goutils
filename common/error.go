package common

// error
type Error string

func (e Error) Error() string {
	return string(e)
}

// throw error
func Throw(err error) {
	panic(err)
}

// catch error
func Catch(errHandler func(err error)) {
	if err := recover(); err != nil {
		errHandler(err.(error))
	}
}
