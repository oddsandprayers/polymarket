package condition

import (
	"testing"
)

func Test_Client_Onchain_Condition_Interface(t *testing.T) {
	var _ Interface = &Condition{}
}
