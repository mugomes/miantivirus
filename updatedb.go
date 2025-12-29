package main

import (
	c "mugomes/miantivirus/controls"
	"github.com/mugomes/mgrun"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/mugomes/mgdialogbox"
	"github.com/mugomes/mgsmartflow"
)

func showUpdateDB(app fyne.App) {
	c.LoadTranslations()
	window := app.NewWindow(c.T("Update Database"))
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(400, 400))

	flow := mgsmartflow.New()

	lblInfo := widget.NewLabel(c.T("Checking for updates..."))
	lblInfo.TextStyle = fyne.TextStyle{Bold: true}

	txtResult := widget.NewEntry()
	txtResult.MultiLine = true
	txtResult.Scroll = fyne.ScrollBoth
	txtResult.Wrapping = fyne.TextWrapBreak

	flow.AddRow(lblInfo)
	flow.AddRow(txtResult)
	flow.SetResize(txtResult, fyne.NewSize(window.Canvas().Size().Width, 279))

	go func() {
		// 
		s := mgrun.New("pkexec sh -c 'killall freshclam;freshclam'")
		pathHome,_ := os.UserHomeDir()
		s.SetDir(pathHome)
		s.AddEnv("teste", "abc")

		s.OnStderr(func(s string) {
			fyne.Do(func() {
				txtResult.Text += s + "\n"
				txtResult.Refresh()
			})
		})
		s.OnStdout(func(s string) {
			fyne.Do(func() {
				txtResult.Text += s + "\n"
				txtResult.Refresh()
			})
		})

		if err := s.Run(); err != nil {
			mgdialogbox.NewAlert(app, c.T("Update Database"), err.Error(), true, "Ok")
		}

		fyne.Do(func() {
			lblInfo.SetText(c.T("Finish"))
		})
	}()
	window.SetContent(flow.Container)
	window.Show()
}
