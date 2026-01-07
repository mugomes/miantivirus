// Copyright (C) 2024-2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/mugomes/mgcolumnview"
	"github.com/mugomes/mgdialogbox"

	c "mugomes/miantivirus/controls"

	"github.com/mugomes/mgsmartflow"
)

const VERSION_APP string = "2.0.0"

func main() {
	c.LoadTranslations()

	app := app.NewWithID("br.com.mugomes.miantivirus")
	app.Settings().SetTheme(&myDarkTheme{})
	app.SetIcon(resourcePng)

	window := app.NewWindow("MiAntivirus")
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(659, 457))

	mnuTools := fyne.NewMenu(c.T("Tools"),
		fyne.NewMenuItem(c.T("Update Database"), func() {
			showUpdateDB(app)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(c.T("Options"), func() {
			showOptions(app)
		}),
	)

	mnuAbout := fyne.NewMenu(c.T("About"),
		fyne.NewMenuItem(c.T("Check for Updates"), func() {
			url, _ := url.Parse("https://github.com/mugomes/miantivirus/releases")
			app.OpenURL(url)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(c.T("Support MiAntivirus"), func() {
			url, _ := url.Parse("https://mugomes.github.io/apoie.html")
			app.OpenURL(url)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(c.T("About MiAntivirus"), func() {
			showAbout(app)
		}),
	)
	window.SetMainMenu(fyne.NewMainMenu(mnuTools, mnuAbout))

	flow := mgsmartflow.New()

	ctnSpace1 := widget.NewLabel(" ")

	flow.AddRow(ctnSpace1)
	flow.SetResize(ctnSpace1, fyne.NewSize(window.Canvas().Size().Width, 7))

	lstArquivos := mgcolumnview.NewColumnView(
		[]string{c.T("Files")},
		[]float32{38, 400, 100}, true,
	)

	btnAddFile := widget.NewButton(c.T("Add File"), func() {
		mgdialogbox.NewOpenFile(app, c.T("Open Files"), []string{}, true, func(filenames []string) {
			for _, filename := range filenames {
				lstArquivos.AddRow([]string{filename})
			}
		})
	})

	btnAddFolder := widget.NewButton(c.T("Add Folder"), func() {
		mgdialogbox.NewSelectDirectory(app, c.T("Select Directory"), true, func(filenames []string) {
			for _, filename := range filenames {
				lstArquivos.AddRow([]string{filename})
			}
		})
	})

	btnRemoverLista := widget.NewButton(c.T("Remove from List"), func() {
		lstArquivos.RemoveSelected()
	})

	flow.AddColumn(btnAddFile, btnAddFolder, btnRemoverLista)
	flow.SetGap(btnAddFile, fyne.NewPos(7, 17))

	flow.AddRow(lstArquivos)
	flow.SetResize(lstArquivos, fyne.NewSize(window.Canvas().Size().Width, 272))

	btnEscanear := widget.NewButton(c.T("Scan"), func() {
		showScan(app, lstArquivos.ListAll())
	})

	flow.AddColumn(
		layout.NewSpacer(),
		btnEscanear,
		layout.NewSpacer(),
	)

	window.SetContent(flow.Container)
	window.ShowAndRun()
}
