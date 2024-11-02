//go:build !ci

package gl

import (
	"runtime"
	"testing"

	"github.com/pirST/fyne/v2"
	"github.com/pirST/fyne/v2/canvas"
	"github.com/pirST/fyne/v2/test"
	"github.com/pirST/fyne/v2/theme"
)

func init() {
	runtime.LockOSThread()
}

func TestDrawImage_Ratio(t *testing.T) {
	//	d := NewGLDriver()
	//	win := d.CreateWindow("Test")
	//	c := win.Canvas().(*glCanvas)

	test.NewApp() // Need an app started to get safeIconLookup to work.

	img := canvas.NewImageFromResource(theme.ComputerIcon())
	img.Resize(fyne.NewSize(10, 10))
	//	c.newGlImageTexture(img)
	//	assert.Equal(t, float32(1.0), c.aspects[img])
}

func TestDrawImage_Ratio2(t *testing.T) {
	//	d := NewGLDriver()
	//	win := d.CreateWindow("Test")
	//	c := win.Canvas().(*glCanvas)

	test.NewApp() // Need an app started to get safeIconLookup to work.

	// make sure we haven't used the visual ratio
	img := canvas.NewImageFromResource(theme.ComputerIcon())
	img.Resize(fyne.NewSize(20, 10))
	//	c.newGlImageTexture(img)
	//	assert.Equal(t, float32(1.0), c.aspects[img])
}
