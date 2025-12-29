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

	"github.com/mugomes/mgsettings/v2"
	"github.com/mugomes/mgsmartflow"
)

func showScan(app fyne.App, listAll [][]string) {
	c.LoadTranslations()

	window := app.NewWindow(c.T("Scan"))
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(800, 600))

	flow := mgsmartflow.New()

	lblVerificando := widget.NewLabel(c.T("Check:"))
	lblInfo := widget.NewLabel("")
	lblInfo.SetText(c.T("Scanning..."))
	flow.AddColumn(lblVerificando, lblInfo)
	flow.SetResize(lblVerificando, fyne.NewSize(100, 38))

	lstArquivos := mgcolumnview.NewColumnView(
		[]string{c.T("Files"), ""},
		[]float32{38, 400, 400}, true,
	)

	flow.AddRow(lstArquivos)
	flow.SetResize(lstArquivos, fyne.NewSize(window.Canvas().Size().Width, 272))

	btnGerarRelatorio := widget.NewButton(c.T("Generate Report"), func() {
		mgdialogbox.NewSelectDirectory(app, c.T("Save File"), false, func(s []string) {
			if len(s) > 0 && len(lstArquivos.ListAll()) > 0 {
				var txt strings.Builder

				txt.WriteString(c.T("Date: ", time.Now()))
				
				for _, result := range lstArquivos.ListAll() {
					if len(result) > 0 {
						for _, row := range result {
							txt.WriteString(row)
						}
					}
				}

				if err := os.WriteFile(filepath.Join(s[0], "report.txt"), []byte(txt.String()), os.ModeAppend); err != nil {
					fmt.Println("Error: ", err.Error())
				}
			}
		})
	})

	btnRemoverArquivo := widget.NewButton(c.T("Remove File"), func() {
		
	})

	flow.AddColumn(btnGerarRelatorio, btnRemoverArquivo)

	mgconfig := mgsettings.Load("miantivirus", true)

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
	options := mgconfig.Get("options", []string{}).([]interface{})

	for _, row := range options {
		s := row.(string)
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

	tamanho = int(mgconfig.Get("filesize", 0).(float64))
	if tamanho > 0 {
		command += " --max-filesize=" + strconv.Itoa(tamanho) + "M"
	}

	ignorefolders := mgconfig.Get("ignorefolders", []string{}).([]interface{})

	for _, row := range ignorefolders {
		sub, ok := row.([]interface{})
		if !ok {
			continue
		}

		for _, item := range sub {
			s, ok := item.(string)
			if ok {
				command += " --exclude-dir=\"" + s + "\""
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
				command += " --exclude=\"" + s + "\""
			}
		}
	}

	for _, row := range listAll {
		sub := row
		for _, item := range sub {
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

					lstArquivos.AddRow([]string{filename, tipov})
				}
			}
		})

		if err := s.Run(); err != nil {
			fmt.Println("Error: ", err.Error())
		}

		if s.ExitCode() >= 0 {
			lblInfo.SetText(c.T("Finish"))
		}
	}()

	window.SetContent(flow.Container)
	window.Show()
}
