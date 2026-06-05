package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type transparentTheme struct{}

func (t transparentTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		return color.NRGBA{R: 30, G: 30, B: 30, A: 200}
	}
	if name == theme.ColorNameOverlayBackground {
		return color.NRGBA{R: 30, G: 30, B: 30, A: 200}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t transparentTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t transparentTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t transparentTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&transparentTheme{})
	myWindow := myApp.NewWindow("Tour of go")
	myWindow.Resize(fyne.NewSize(500, 400))

	steps := []string{
		"tip 1: Introduction\n\nthis is my tour of go by jackson in terminal\n\nthis was made by 404-not-found129 on github learning go from the website tour of go https://go.dev/tour it will be easy to understand and easy to read",
		"tip 2: fmt package\n\nfmt: fmt is the default package in go\n\nfmt is used to print text to the terminal.\nexample:\n\nfmt.Println(\"Hello, World!\")",
		"tip 3: Integers\n\nintergers: a interger (a number without a deciaml) can be defined with the int word\n\nexample:\n\nvar x int = 10",
		"tip 4: Variables\n\nvarible: a varible (a number or letter that can be manipulized) can be defined with int\n\nexample:\n\nvar name string = \"jackson\"",
		"tip 5: Functions\n\nfunction: a function is like a instruction when doing\n\nexample:\n\nfunc add(x, y int) int {\n\treturn x + y\n}\nfmt.Println(add(3, 4))",
		"tip 6: Booleans\n\na bool aka boolian is a answer with true or false\n\nexample:\nvar isReady bool = true",
		"tip 7: Editors\n\neditor: a editor is very very importent in coding it is what makes you able to code\nthere are a few editors i would reccomend:\n\n1. vim\n2. jetbrains\n3. vs-codium",
		"tip 8: Keyboard\n\nuse a split keyboard in order to type faster",
		"tip 9: operating systems\n\ni use arch bcause of the compatibilty with widow managers and its genarly my favorite with code and also macos is great but windows is trash\n\nthis is the top operating systems to code with\n1: linux\n2: macos\n3: windows",
		"tip 10: focusing on one language\n\nthe reason you want to focus on learning oen ocding language is that you learn alot more like lets say of you use go it has a simliar syntax to c and c is very similar to rust",
	}

	currentStep := 0
	stepDisplay := widget.NewLabel(steps[currentStep])

	progress := widget.NewProgressBar()
	progress.SetValue(float64(currentStep+1) / float64(len(steps)))

	nextButton := widget.NewButton("Next", nil)
	prevButton := widget.NewButton("Back", nil)
	prevButton.Disable()
	exitButton := widget.NewButton("Exit", func() {
		myApp.Quit()
	})


// Source - https://stackoverflow.com/q/77975268
// Posted by artem6987, modified by community. See post 'Timeline' for change history
// Retrieved 2026-06-05, License - CC BY-SA 4.0

background := canvas.NewRectangle(color.RGBA{R: 205, G: 92, B: 92, A: 255})
container.NewMax(background)

btn := widget.NewButton("test", func() {})


	nextButton.OnTapped = func() {
		if currentStep < len(steps)-1 {
			currentStep++
			stepDisplay.SetText(steps[currentStep])
			progress.SetValue(float64(currentStep+1) / float64(len(steps)))
			prevButton.Enable()
		}
		if currentStep == len(steps)-1 {
			nextButton.Disable()
		}
	}

	prevButton.OnTapped = func() {
		if currentStep > 0 {
			currentStep--
			stepDisplay.SetText(steps[currentStep])
			progress.SetValue(float64(currentStep+1) / float64(len(steps)))
			nextButton.Enable()
		}
		if currentStep == 0 {
			prevButton.Disable()
		}
	}

	buttonLayout := container.NewHBox(prevButton, nextButton, exitButton)
	mainLayout := container.NewVBox(
		widget.NewLabelWithStyle("Tour of Go by Jackson", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		stepDisplay,
		buttonLayout,
		
		
		
		progress,
	)

	myWindow.SetContent(mainLayout)
	myWindow.ShowAndRun()
}
