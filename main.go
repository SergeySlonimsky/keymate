package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/SergeySlonimsky/keymate/lang"
)

func main() {
	translateLang := lang.EnRu
	myApp := app.New()
	myWindow := myApp.NewWindow("Keymate")
	myWindow.Resize(fyne.Size{Width: 640, Height: 250})
	inputArea := widget.NewMultiLineEntry()
	resultArea := widget.NewMultiLineEntry()
	langSelect := widget.NewSelect([]string{"Eng-Ru"}, func(val string) {
		if val == "Eng-Ru" {
			translateLang = lang.EnRu
		}
	})
	langSelect.Selected = "Eng-Ru"
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Entry", Widget: inputArea},
			{Text: "Result", Widget: resultArea},
		},
		OnSubmit: func() {
			res := lang.Convert(inputArea.Text, translateLang)
			fmt.Println(res)
			resultArea.SetText(res)
		},
	}
	form.SubmitText = "Translate"
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), langSelect, form))
	myWindow.ShowAndRun()
}
