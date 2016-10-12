package errors

import "fmt"

type Fatal interface {
	IsFatal() bool
}

func IsFatal(err error) bool {
	if f, ok := err.(Fatal); ok {
		if f.IsFatal() {
			return true
		}
	}
	return false
}

func Fatalf(format string, a ...interface{}) error {
	err := fatal{isfatal: true}
	if len(a) == 0 {
		err.message = format
	} else {
		err.message = fmt.Sprint(format, a)
	}
	return err
}

type fatal struct {
	message string
	isfatal bool
}

func (f fatal) Error() string { return f.message }
func (f fatal) IsFatal() bool { return f.isfatal }
