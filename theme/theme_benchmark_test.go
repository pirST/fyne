package theme

import (
	"testing"

	"github.com/pirST/fyne/v2"
)

func BenchmarkTheme_current(b *testing.B) {
	fyne.CurrentApp().Settings().SetTheme(LightTheme())

	for n := 0; n < b.N; n++ {
		Current()
	}
}
