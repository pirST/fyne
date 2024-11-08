package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pirST/fyne/v2/canvas"
	"github.com/pirST/fyne/v2/theme"
)

func TestSyncPool(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		pool := &syncPool{}
		assert.Nil(t, pool.Obtain())
	})
	t.Run("Single", func(t *testing.T) {
		pool := &syncPool{}
		rect := canvas.NewRectangle(theme.Color(theme.ColorNamePrimary))
		pool.Release(rect)
		assert.Equal(t, rect, pool.Obtain())
		assert.Nil(t, pool.Obtain())
	})
	t.Run("Multiple", func(t *testing.T) {
		pool := &syncPool{}
		rect := canvas.NewRectangle(theme.Color(theme.ColorNamePrimary))
		circle := canvas.NewCircle(theme.Color(theme.ColorNamePrimary))
		pool.Release(rect)
		pool.Release(circle)
		a := pool.Obtain()
		b := pool.Obtain()
		assert.NotNil(t, a)
		assert.NotNil(t, b)
		if a == rect && b == circle {
			// Pass
		} else if a == circle && b == rect {
			// Pass
		} else {
			t.Error("Obtained incorrect objects")
		}
		assert.Nil(t, pool.Obtain())
	})
}
