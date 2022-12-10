package proj3

import (
	"testing"
)

func TestForChatRoom(t *testing.T)  {
	sig := GetSignKeys();
	if sig == nil {
		t.Error("signature not generated")
	}
}