// Copyright (C) 2024-2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.mugomes.com.br

package main

import (
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/mugomes/mgcolumnview"
	"github.com/mugomes/mgdialogbox"

	c "mugomes/miantivirus/controls"

	//"github.com/mugomes/mgsettings"
	"github.com/mugomes/mgsmartflow"
)

const VERSION_APP string = "6.1.1"

type myDarkTheme struct{}

func (m myDarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// A lógica para forçar o modo escuro é retornar cores escuras.
	// O Fyne usa estas constantes internamente:
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{28, 28, 28, 255} // Fundo preto
	case theme.ColorNameForeground:
		return color.White // Texto branco
	// Adicione outros casos conforme a necessidade (InputBackground, Primary, etc.)
	default:
		// Retorna o tema escuro padrão para as outras cores (se existirem)
		// Aqui estamos apenas definindo as cores principais para garantir o Dark Mode
		return theme.DefaultTheme().Color(name, theme.VariantDark)
	}
}

// 3. Implemente os outros métodos necessários da interface fyne.Theme (usando o tema padrão)
func (m myDarkTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

func (m myDarkTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m myDarkTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

func main() {
	c.LoadTranslations()

	app := app.NewWithID("br.com.mugomes.miantivirus")
	app.Settings().SetTheme(&myDarkTheme{})

	window := app.NewWindow("MiAntivirus")
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(659, 457))

	mnuTools := fyne.NewMenu(c.T("Tools"),
		fyne.NewMenuItem(c.T("Update Database"), func() {}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(c.T("Options"), func() {
			showOptions(app)
		}),
	)

	mnuAbout := fyne.NewMenu(c.T("About"),
		fyne.NewMenuItem(c.T("Check for Updates"), func() {
			url, _ := url.Parse("https://www.mugomes.com.br/p/miantivirus.html")
			app.OpenURL(url)
		}),
		// fyne.NewMenuItem("Seja um assinante", func() {
		// 	url, _ := url.Parse("https://www.mugomes.com.br/p/assinantes.html")
		// 	app.OpenURL(url)
		// }),
		fyne.NewMenuItem(c.T("About MiAntivirus"), func() {}),
	)
	window.SetMainMenu(fyne.NewMainMenu(mnuTools, mnuAbout))

	flow := mgsmartflow.New()

	ctnSpace1 := widget.NewLabel(" ")

	flow.AddRow(ctnSpace1)
	flow.SetResize(ctnSpace1, fyne.NewSize(window.Canvas().Size().Width, 7))

	lstArquivos := mgcolumnview.NewColumnView(
		[]string{c.T("Files")},
		[]float32{38, 400}, true,
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

	})

	flow.AddColumn(
		layout.NewSpacer(),
		btnEscanear,
		layout.NewSpacer(),
	)

	window.SetContent(flow.Container)
	window.ShowAndRun()
}
