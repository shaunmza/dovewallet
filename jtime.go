package dovewallet

import (
	"encoding/json"
	"time"
	"fmt"
)

const TIME_FORMAT = "2006-01-02T15:04:05"

type jTime time.Time


func (jt *jTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(TIME_FORMAT, s)
	if err != nil {
		return err
	}
	*jt = jTime(t)
	return nil
}

func (jt jTime) MarshalJSON() ([]byte, error) {
	return []byte(jt.String()), nil
}

func (jt *jTime) String() string {
	t := time.Time(*jt)
	return fmt.Sprintf("%q", t.Format(TIME_FORMAT))
}