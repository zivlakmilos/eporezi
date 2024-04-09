package gui

import (
	"net/url"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run(eporeziUrl string) {
	a := app.New()
	w := a.NewWindow("ePorezi")

	query, err := parseUrl(eporeziUrl)
	if err != nil {
		return
	}

	lblTitle := widget.NewLabel("ePorezi")
	lblUrl := widget.NewLabel(query.Get("loginKey"))

	w.SetContent(container.NewVBox(
		lblTitle,
		lblUrl,
	))

	w.ShowAndRun()
}

func parseUrl(u string) (url.Values, error) {
	query := strings.ReplaceAll(u, "eporezi://", "")
	if len(query) > 0 && query[len(query)-1] == '/' {
		query = query[:len(query)-1]
	}

	q, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}

	return q, nil
}
