package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pirST/fyne/v2"
	"github.com/pirST/fyne/v2/internal/cache"
	"github.com/pirST/fyne/v2/theme"
)

func TestMenuItem_Disabled(t *testing.T) {
	i := fyne.NewMenuItem("Disabled", func() {})
	m := fyne.NewMenu("top", []*fyne.MenuItem{i}...)
	i.Disabled = true
	w := newMenuItem(i, NewMenu(m))
	r := cache.Renderer(w)

	assert.Equal(t, theme.Color(theme.ColorNameDisabled), r.(*menuItemRenderer).text.Color)
}
