package token

import (
	"testing"
)

func Test_Client_Offchain_Token_Interface(t *testing.T) {
	var _ Interface = &Token{}
}
