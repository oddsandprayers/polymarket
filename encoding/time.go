package encoding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/xh3b4sd/tracer"
)

var (
	layout = []string{
		"2006-01-02T15:04:05.999999Z",
		"2006-01-02 15:04:05Z07",
	}
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(byt []byte) error {
	var err error

	// Accept empty values.

	if len(byt) == 0 || bytes.Equal(byt, []byte(`""`)) {
		return nil
	}

	// Accept JSON null values.

	if bytes.Equal(byt, []byte("null")) {
		return nil
	}

	// Accept time formats, e.g. "2025-12-08T18:03:03.779632Z".

	if byt[0] == '"' {
		var str string
		{
			err = json.Unmarshal(byt, &str)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		var tim time.Time
		{
			tim, err = tryTim(str)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			t.Time = tim
		}

		return nil
	}

	// Accept unix timestamps, e.g. 1765823498.

	var sec int64
	{
		err = json.Unmarshal(byt, &sec)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		t.Time = time.Unix(sec, 0).UTC()
	}

	return nil
}

func tryTim(str string) (time.Time, error) {
	for _, x := range layout {
		tim, err := time.Parse(x, str)
		if err == nil {
			return tim.UTC(), nil
		}
	}

	return time.Time{}, tracer.Mask(fmt.Errorf("invalid time format %q", str))
}
