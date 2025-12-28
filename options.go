// Copyright (C) 2024-2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.mugomes.com.br

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/mugomes/mgcolumnview"
	"github.com/mugomes/mgdialogbox"
	"github.com/mugomes/mgnumericentry"
	"github.com/mugomes/mgsettings/v2"
	"github.com/mugomes/mgsmartflow"

	c "mugomes/miantivirus/controls"
)

func showOptions(app fyne.App) {
	c.LoadTranslations()

	window := app.NewWindow(c.T("Options"))
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(400, 600))

	mgconfig := mgsettings.Load("miantivirus", true)

	flowGeral := mgsmartflow.New()

	lblEscanear := widget.NewLabel(c.T("Scan"))
	lblEscanear.TextStyle = fyne.TextStyle{Bold: true}
	chkOptions := widget.NewCheckGroup([]string{
		c.T("Detect PUA (Potentially Unwanted Application)"),
		c.T("Heuristic scanning"),
		c.T("Scan hidden files"),
		c.T("Check symbolic files"),
		c.T("Check symbolic folders"),
		c.T("Scan compressed files"),
		c.T("Scan Email"),
	}, nil)

	flowGeral.AddRow(lblEscanear)
	flowGeral.AddRow(chkOptions)

	lblTamanho := widget.NewLabel(c.T("File Size for Scanning"))
	lblTamanho.TextStyle = fyne.TextStyle{Bold: true}
	ctnTamanho, txtTamanho := mgnumericentry.NewMGNumericEntryWithButtons(0, 10000000, 0)

	flowGeral.AddRow(lblTamanho)
	flowGeral.AddRow(ctnTamanho)

	flowIgnorar := mgsmartflow.New()

	cvIgnorarPastas := mgcolumnview.NewColumnView([]string{c.T("Ignore Folders")}, []float32{38, 400}, true)
	btnIgnorarPastasAdd := widget.NewButton(c.T("Add"), func() {
		mgdialogbox.NewSelectDirectory(app, c.T("Select Directory"), true, func(s []string) {
			if len(s) > 0 {
				for _, pathname := range s {
					cvIgnorarPastas.AddRow([]string{pathname})
				}
			}
		})
	})

	btnIgnorarPastasRemove := widget.NewButton(c.T("Remove"), func() {
		cvIgnorarPastas.RemoveSelected()
	})

	flowIgnorar.AddColumn(btnIgnorarPastasAdd, btnIgnorarPastasRemove)
	flowIgnorar.AddRow(cvIgnorarPastas)
	flowIgnorar.SetResize(cvIgnorarPastas, fyne.NewSize(window.Canvas().Size().Width, 157))
	
	separator1 := widget.NewSeparator();
	flowIgnorar.AddRow(separator1)

	cvIgnorarArquivos := mgcolumnview.NewColumnView([]string{c.T("Ignore Files")}, []float32{38, 400}, true)
	btnIgnorarArquivosAdd := widget.NewButton(c.T("Add"), func() {
		mgdialogbox.NewOpenFile(app, c.T("Open File"), []string{}, true, func(s []string) {
			if len(s) > 0 {
				for _, filename := range s {
					cvIgnorarArquivos.AddRow([]string{filename})
				}
			}
		})
	})

	btnIgnorarArquivosRemove := widget.NewButton(c.T("Remove"), func() {
		cvIgnorarArquivos.RemoveSelected()
	})
	flowIgnorar.AddColumn(btnIgnorarArquivosAdd, btnIgnorarArquivosRemove)
	flowIgnorar.AddRow(cvIgnorarArquivos)
	flowIgnorar.SetResize(cvIgnorarArquivos, fyne.NewSize(window.Canvas().Size().Width, 157))

	tabs := container.NewAppTabs(
		container.NewTabItem(c.T("General"), flowGeral.Container),
		container.NewTabItem(c.T("Ignore"), flowIgnorar.Container),
	)

	// Load
	options := mgconfig.Get("options", []string{}).([]interface{})

	var value []string
	for _, row := range options {
		value = append(value, row.(string))
	}

	chkOptions.SetSelected(value)
	txtTamanho.SetValue(int(mgconfig.Get("filesize", 0).(float64)))

	ignorefolders := mgconfig.Get("ignorefolders", []string{}).([]interface{})

	for _, row := range ignorefolders {
		sub, ok := row.([]interface{})
		if !ok {
			continue
		}

		for _, item := range sub {
			s, ok := item.(string)
			if ok {
				cvIgnorarPastas.AddRow([]string{s})
			}
		}
	}

	ignorefiles := mgconfig.Get("ignorefiles", []string{}).([]interface{})

	for _, row := range ignorefiles {
		sub, ok := row.([]interface{})
		if !ok {
			continue
		}

		for _, item := range sub {
			s, ok := item.(string)
			if ok {
				cvIgnorarArquivos.AddRow([]string{s})
			}
		}
	}
	
	flowIgnorar.AddRow(widget.NewLabel(" "))

	// Save
	btnSave := widget.NewButton(c.T("Save"), func() {
		mgconfig.Set("options", chkOptions.Selected)
		mgconfig.Set("filesize", txtTamanho.GetValue())
		mgconfig.Set("ignorefolders", cvIgnorarPastas.ListAll())
		mgconfig.Set("ignorefiles", cvIgnorarArquivos.ListAll())

		mgconfig.Save()
		window.Close()
	})

	window.SetContent(container.NewVBox(
		tabs,
		container.NewHBox(
			layout.NewSpacer(),
			btnSave,
			layout.NewSpacer(),
		),
	))
	window.Show()
}
