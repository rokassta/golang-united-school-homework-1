package solution

import (
	"github.com/kyokomi/emoji/v2"
	"testing"
)

func TestGetMessage(t *testing.T) {
	result := GetMessage()
	expected := emoji.Sprint("Hello :world_map:!")

	if expected != result {
		t.Error("Expected ", expected, ", got ", result)
	}
}
