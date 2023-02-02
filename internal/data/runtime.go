package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(fmt.Sprintf("%d mins", r))), nil
}

func (r *Runtime) UnmarshalJSON(js []byte) error {
	jsNoQuote, err := strconv.Unquote(string(js))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(jsNoQuote, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	mins, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(mins)
	return nil
}
