package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"

	c "mugomes/miantivirus/controls"

	"github.com/mugomes/mgsettings/v2"
)

func showScan(app fyne.App, listAll [][]string) {
	c.LoadTranslations()

	window := app.NewWindow(c.T("Scan"))
	window.CenterOnScreen()
	window.SetFixedSize(true)

	mgconfig := mgsettings.Load("miantivirus", true)

	var (
		command           string = "--verbose --recursive=yes --no-summary"
		pua               string = " --detect-pua=no --alert-broken=no --alert-macros=no"
		heuristica        string = " --heuristic-alerts=no"
		arquivocompactado string = " --scan-archive=no"
		arquivosocultos   string = ""
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
			arquivosocultos = " --exclude=\"\\/\\.\""
			arquivosocultos += " --exclude-dir=\"\\/\\.\""
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

	sParts := strings.Fields(command)
	sExec := exec.Command("clamscan", sParts[0:]...)
	
	window.Show()
}
