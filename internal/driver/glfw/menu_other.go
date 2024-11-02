//go:build !darwin || wasm || test_web_driver || no_native_menus

package glfw

import "github.com/pirST/fyne/v2"

func hasNativeMenu() bool {
	return false
}

func setupNativeMenu(_ *window, _ *fyne.MainMenu) {
	// no-op
}
