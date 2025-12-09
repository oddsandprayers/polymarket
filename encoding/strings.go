package encoding

import (
	"bytes"
	"encoding/json"

	"github.com/xh3b4sd/tracer"
)

type Strings []string

func (s *Strings) UnmarshalJSON(byt []byte) error {
	var err error

	if bytes.Equal(byt, []byte("null")) {
		return nil
	}

	var str string
	{
		err = json.Unmarshal(byt, &str)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if str == "" {
		{
			*s = nil
		}

		return nil
	}

	var lis []string
	{
		err = json.Unmarshal([]byte(str), &lis)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		*s = lis
	}

	return nil
}
