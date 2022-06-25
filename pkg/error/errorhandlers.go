package errorhandlers

import (
	"fmt"
)

type ErrNotFound struct {
	Url     string
	Code    int
	Message string
}

func (enf ErrNotFound) Error() string {
	return fmt.Sprintf(
		"Not found - url: '%s', response code: '%d', response: '%s'",
		enf.Url,
		enf.Code,
		enf.Message,
	)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
