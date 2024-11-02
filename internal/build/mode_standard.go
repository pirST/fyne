//go:build !debug && !release

package build

import "github.com/pirST/fyne/v2"

// Mode is the application's build mode.
const Mode = fyne.BuildStandard
