package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {

	s := fmt.Sprintf("%d mins", r)

	quotedS := strconv.Quote(s)

	return []byte(quotedS), nil

}
