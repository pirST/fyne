package main

import (
	"github.com/pirST/fyne/v2"
	"github.com/pirST/fyne/v2/app"
	"github.com/pirST/fyne/v2/cmd/fyne_settings/settings"
	"github.com/pirST/fyne/v2/container"
)

func main() {
	s := settings.NewSettings()

	a := app.New()
	w := a.NewWindow("Fyne Settings")

	appearance := s.LoadAppearanceScreen(w)
	tabs := container.NewAppTabs(
		&container.TabItem{Text: "Appearance", Icon: s.AppearanceIcon(), Content: appearance})
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)

	w.Resize(fyne.NewSize(520, 520))
	w.ShowAndRun()
}
