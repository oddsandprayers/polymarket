package condition

import (
	"github.com/xh3b4sd/tracer"
)

var searchError = &tracer.Error{
	Description: "The search request expects the response data to be valid. Some parts of the response data could not be validated. Therefore the request failed.",
}
