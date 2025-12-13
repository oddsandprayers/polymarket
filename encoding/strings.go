package encoding

import (
	"bytes"
	"encoding/json"

	"github.com/xh3b4sd/tracer"
)

type Strings []string

func (s *Strings) UnmarshalJSON(byt []byte) error {
	var err error

	// Accept empty values.

	if len(byt) == 0 || bytes.Equal(byt, []byte(`""`)) {
		return nil
	}

	// Accept JSON null values.

	if bytes.Equal(byt, []byte("null")) {
		return nil
	}

	// Accept typed string slices.

	if byt[0] != '[' {
		var str string
		{
			err = json.Unmarshal(byt, &str)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			byt = []byte(str)
		}
	}

	// Accept JSON encoded strings.

	var lis []string
	{
		err = json.Unmarshal(byt, &lis)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		*s = lis
	}

	return nil
}
