package mobile

import (
	"testing"

	_ "github.com/pirST/fyne/v2/test"
)

func TestDevice_HideVirtualKeyboard_BeforeRun(t *testing.T) {
	hideVirtualKeyboard() // should not crash!
}
