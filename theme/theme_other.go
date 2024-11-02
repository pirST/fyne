//go:build !hints

package theme

import (
	"image/color"

	"github.com/pirST/fyne/v2"
)

var (
	fallbackColor = color.Transparent
	fallbackIcon  = &fyne.StaticResource{}
)
