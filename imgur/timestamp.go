package imgur

import (
	"strconv"
	"time"
	"unsafe"
)

type Timestamp time.Time

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	n, err := strconv.ParseInt(*(*string)(unsafe.Pointer(&data)), 10, 64)
	if err != nil {
		return err
	}
	*t = Timestamp(time.Unix(n, 0))
	return nil
}
