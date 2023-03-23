package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r *Runtime) MarshalJSON() ([]byte, error) {
	res := fmt.Sprintf(`"%d mins"`, r)
	return []byte(res), nil
}

func (r *Runtime) UnmarshalJSON(js []byte) error {
	json, err := strconv.Unquote(string(js))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(json, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	runtime, err := strconv.Atoi(parts[0])
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(runtime)

	return nil
}
