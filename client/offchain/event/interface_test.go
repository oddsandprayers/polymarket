package event

import (
	"testing"
)

func Test_Client_Offchain_Event_Interface(t *testing.T) {
	var _ Interface = &Event{}
}
