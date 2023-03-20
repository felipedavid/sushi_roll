package data

import "fmt"

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	res := fmt.Sprintf(`"%d mins"`, r)
	return []byte(res), nil
}
