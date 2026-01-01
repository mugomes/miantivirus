// Copyright (C) 2024-2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	c "mugomes/miantivirus/controls"

	"github.com/mugomes/mgcolumnview"
	"github.com/mugomes/mgdialogbox"
	"github.com/mugomes/mgrun"

	"github.com/mugomes/mgsettings/v3"
	"github.com/mugomes/mgsmartflow"
)

func showScan(app fyne.App, listAll []mgcolumnview.SelectRow) {
	c.LoadTranslations()

	window := app.NewWindow(c.T("Scan"))
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(800, 379))

	flow := mgsmartflow.New()

	lblVerificando := widget.NewLabel(c.T("Check:"))
	lblInfo := widget.NewLabel("")
	lblInfo.SetText(c.T("Scanning..."))
	flow.AddColumn(lblVerificando, lblInfo)
	flow.SetResize(lblVerificando, fyne.NewSize(79, 38))

	lstArquivos := mgcolumnview.NewColumnView(
		[]string{c.T("Files"), "", ""},
		[]float32{38, 400, 179, 79}, true,
	)

	flow.AddRow(lstArquivos)
	flow.SetResize(lstArquivos, fyne.NewSize(window.Canvas().Size().Width-7, 272))

	btnGerarRelatorio := widget.NewButton(c.T("Generate Report"), func() {
		mgdialogbox.NewSelectDirectory(app, c.T("Save File"), false, func(s []string) {
			if len(s) > 0 && len(lstArquivos.ListAll()) > 0 {
				var txt strings.Builder
				var sData string

				currentTime := time.Now()

				txt.WriteString(c.T("Date: %v\n\n", currentTime.Format("2006-01-02-15-04-05")))

				for _, result := range lstArquivos.ListAll() {
					if len(result.Data) > 0 {
						sData = ""
						for _, row := range result.Data {
							if row != "" {
								sData += row + " "
							}
						}
						txt.WriteString(sData + "\n\n")
					}
				}

				if err := os.WriteFile(filepath.Join(s[0], fmt.Sprintf("report-%v.txt", currentTime.Format("2006-01-02-15-04-05"))), []byte(txt.String()), os.ModePerm); err != nil {
					fmt.Println("Error: ", err.Error())
				}
			}
		})
	})
	btnGerarRelatorio.Disable()

	btnRemoverArquivo := widget.NewButton(c.T("Remove File"), func() {
		var data string
		lblInfo.SetText(c.T("Removing selected files..."))
		for _, row := range lstArquivos.ListSelected() {
			if len(row.Data) > 0 {
				data = ""
				for _, items := range row.Data {
					if items != "" {
						data += items + "|"
					}
				}

				filename := strings.Split(data, "|")
				if len(filename) > 0 {
					if err := os.Remove(filename[0]); err != nil {
						mgdialogbox.NewAlert(app, c.T("Remove File"), err.Error(), true, "Ok")
					} else {
						lstArquivos.UpdateColumnItem(row.ID, 2, c.T("Deleted"))
					}
				}
			}
		}
		lblInfo.SetText(c.T("Finish"))
	})
	btnRemoverArquivo.Disable()

	btnCancelar := widget.NewButton(c.T("Cancel"), func() {
		run := mgrun.New("killall clamscan")
		if err := run.Run(); err != nil {
			fmt.Println("Error: ", err.Error())
		}
		window.Close()
	})

	flow.AddColumn(btnGerarRelatorio, btnRemoverArquivo, btnCancelar)

	mgconfig, _ := mgsettings.Load("miantivirus", true)

	var (
		command           string = "--verbose --recursive=yes --no-summary"
		pua               string = " --detect-pua=no --alert-broken=no --alert-macros=no"
		heuristica        string = " --heuristic-alerts=no"
		arquivocompactado string = " --scan-archive=no"
		arquivosocultos   string = " --exclude=\"\\/\\.\" --exclude-dir=\"\\/\\.\""
		arquivosimbolico  string = ""
		pastasimbolica    string = ""
		email             string = " --scan-mail=no"
		tamanho                  = 0
	)
	options := mgconfig.GetStringSlice("options", []string{})

	for _, row := range options {
		s := row

		if strings.Contains(s, "1)") {
			pua = " --detect-pua=yes --alert-broken=yes --alert-macros=yes"
		} else if strings.Contains(s, "2)") {
			heuristica = " --heuristic-alerts=yes"
		} else if strings.Contains(s, "3)") {
			arquivosocultos = ""
		} else if strings.Contains(s, "4)") {
			arquivosimbolico = " --follow-file-symlinks=1"
		} else if strings.Contains(s, "5)") {
			pastasimbolica = " --follow-dir-symlinks=1"
		} else if strings.Contains(s, "6)") {
			arquivocompactado = " --scan-archive=yes"
		} else if strings.Contains(s, "7)") {
			email = " --scan-mail=yes"
		}
	}

	command += pua + heuristica + arquivosocultos + arquivosimbolico + pastasimbolica + arquivocompactado + email

	tamanho = mgconfig.GetInt("filesize", 0)
	if tamanho > 0 {
		command += " --max-filesize=" + strconv.Itoa(tamanho) + "M"
	}

	ignorefolders := mgconfig.GetStringSlice("ignorefolders", []string{})

	for _, row := range ignorefolders {
		if row != "" {
			command += " --exclude-dir=\"" + row + "\""
		}
	}

	ignorefiles := mgconfig.GetStringSlice("ignorefiles", []string{})

	for _, row := range ignorefiles {
		if row != "" {
			command += " --exclude=\"" + row + "\""
		}
	}

	for _, row := range listAll {
		for _, item := range row.Data {
			command += " \"" + item + "\""
		}
	}

	//sParts := strings.Fields(command)

	go func() {
		s := mgrun.New("clamscan " + command)
		pathHome, _ := os.UserHomeDir()
		s.SetDir(pathHome)
		s.OnStdout(func(sLine string) {
			if sLine != "" {
				reScanning := regexp.MustCompile(`\s/.*\.(\S+)`)
				matchScan := reScanning.FindString(sLine)
				if matchScan != "" {
					filename := strings.TrimSpace(strings.ReplaceAll(matchScan, "Scanning", ""))
					fyne.Do(func() {
						lblInfo.SetText(filename)
					})
				}

				reFound := regexp.MustCompile(`(/.*):\s*(.*)\sFOUND`)
				matches := reFound.FindStringSubmatch(sLine)

				if len(matches) >= 3 {
					filename := strings.TrimSpace(matches[1])
					tipov := strings.TrimSpace(matches[2])

					fyne.Do(func() {
						lstArquivos.AddRow([]string{filename, tipov, ""})
					})
				}
			}
		})

		if err := s.Run(); err != nil {
			fmt.Println("Error: ", err.Error())
		}

		if s.ExitCode() >= 0 {
			fyne.Do(func() {
				lblInfo.SetText(c.T("Finish"))
				if len(lstArquivos.ListAll()) > 0 {
					btnGerarRelatorio.Enable()
					btnRemoverArquivo.Enable()
					btnCancelar.Disable()
				} else {
					mgdialogbox.NewAlert(app, "MiAntivirus", c.T("No viruses found!"), false, "Ok")
					window.Close()
				}
			})
		}
	}()

	window.SetContent(flow.Container)
	window.Show()
}
