package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type transparentTheme struct{}

//all the functions
//
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
	myWindow.Resize(fyne.NewSize(600, 520))


			//all the steps 
	steps := []string{
		"1: Introduction\n\nWelcome to the Tour of Go by Jackson.\n\nMade by 404-not-found129 on GitHub, inspired by https://go.dev/tour\n\nThis guide covers the core concepts of Go in a simple, readable format.",
		"2: Packages\n\nEvery Go program starts with a package declaration. The main package is the entry point.\n\nexample:\n\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}",
		"3: fmt\n\nfmt is Go's standard package for formatted I/O. It is used to print text and format strings.\n\nexample:\n\nfmt.Println(\"Hello!\")       // prints with newline\nfmt.Printf(\"Name: %s\\n\", \"Jackson\")  // formatted print\nfmt.Sprintf(\"Score: %d\", 10) // returns a string",
		"4: Variables\n\nA variable stores a value that can change. Use var or the short := syntax.\n\nexample:\n\nvar name string = \"jackson\"\nvar age int = 20\n\n// shorthand (inside a function)\ncity := \"Edinburgh\"",
		"5: Constants\n\nA constant is a value that cannot be changed after it is set. Use the const keyword.\n\nexample:\n\nconst Pi = 3.14159\nconst Language = \"Go\"",
		"6: Integers and Floats\n\nint holds whole numbers. float64 holds numbers with a decimal point.\n\nexample:\n\nvar x int = 10\nvar y float64 = 3.14\n\nsum := x + 5\narea := y * 2.0",
		"7: Strings\n\nStrings are sequences of characters enclosed in double quotes.\n\nexample:\n\nvar greeting string = \"Hello, Go!\"\n\n// concatenation\nfull := \"Hello, \" + \"World!\"\n\n// length\nfmt.Println(len(full)) // 13",
		"8: Booleans\n\nA bool holds either true or false. Used for conditions and logic.\n\nexample:\n\nvar isReady bool = true\nvar done = false\n\nfmt.Println(isReady && !done) // true",
		"9: Functions\n\nFunctions group reusable logic. Go functions can return multiple values.\n\nexample:\n\nfunc add(x, y int) int {\n\treturn x + y\n}\n\nfunc divide(a, b float64) (float64, error) {\n\tif b == 0 {\n\t\treturn 0, errors.New(\"cannot divide by zero\")\n\t}\n\treturn a / b, nil\n}",
		"10: If / Else\n\nif and else let you run code based on a condition.\n\nexample:\n\nif x > 10 {\n\tfmt.Println(\"big\")\n} else if x == 10 {\n\tfmt.Println(\"exactly ten\")\n} else {\n\tfmt.Println(\"small\")\n}\n\n// if with a short statement\nif n := 42; n%2 == 0 {\n\tfmt.Println(\"even\")\n}",
		"11: For Loops\n\nGo has only one loop keyword: for. It covers while-style and range loops too.\n\nexample:\n\n// standard loop\nfor i := 0; i < 5; i++ {\n\tfmt.Println(i)\n}\n\n// while-style\nfor x > 0 {\n\tx--\n}\n\n// range over a slice\nfor i, v := range []int{10, 20, 30} {\n\tfmt.Println(i, v)\n}",
		"12: Switch\n\nswitch is a cleaner way to write multiple if/else branches.\n\nexample:\n\nswitch day {\ncase \"Monday\":\n\tfmt.Println(\"start of the week\")\ncase \"Friday\":\n\tfmt.Println(\"end of the week\")\ndefault:\n\tfmt.Println(\"midweek\")\n}",
		"13: Arrays and Slices\n\nAn array has a fixed size. A slice is a flexible view into an array.\n\nexample:\n\n// array\nvar a [3]int = [3]int{1, 2, 3}\n\n// slice (most common)\nnums := []int{10, 20, 30}\nnums = append(nums, 40)\nfmt.Println(nums[0]) // 10",
		"14: Maps\n\nA map stores key-value pairs. Like a dictionary.\n\nexample:\n\nscores := map[string]int{\n\t\"jackson\": 95,\n\t\"alice\":   87,\n}\n\nscores[\"bob\"] = 78\n\nval, ok := scores[\"alice\"]\nif ok {\n\tfmt.Println(val) // 87\n}",
		"15: Structs\n\nA struct groups related fields together into a custom type.\n\nexample:\n\ntype Person struct {\n\tName string\n\tAge  int\n}\n\np := Person{Name: \"Jackson\", Age: 20}\nfmt.Println(p.Name) // Jackson",
		"16: Pointers\n\nA pointer holds the memory address of a value. Use & to get the address and * to read the value.\n\nexample:\n\nx := 10\np := &x   // p points to x\n*p = 20   // changes x through the pointer\nfmt.Println(x) // 20",
		"17: Error Handling\n\nGo handles errors by returning them as values. Always check the error.\n\nexample:\n\nimport \"errors\"\n\nfunc divide(a, b float64) (float64, error) {\n\tif b == 0 {\n\t\treturn 0, errors.New(\"division by zero\")\n\t}\n\treturn a / b, nil\n}\n\nresult, err := divide(10, 0)\nif err != nil {\n\tfmt.Println(\"Error:\", err)\n}",
		"18: Defer\n\ndefer delays a function call until the surrounding function returns. Great for cleanup.\n\nexample:\n\nfunc readFile() {\n\tf, _ := os.Open(\"file.txt\")\n\tdefer f.Close() // runs when readFile returns\n\n\t// read from f...\n}",
		"19: Goroutines\n\nA goroutine is a lightweight thread managed by Go. Start one with the go keyword.\n\nexample:\n\nfunc sayHello() {\n\tfmt.Println(\"Hello from goroutine\")\n}\n\ngo sayHello() // runs concurrently\nfmt.Println(\"main continues\")",
		"20: Channels\n\nChannels let goroutines communicate safely. Send with <- and receive with <-.\n\nexample:\n\nch := make(chan int)\n\ngo func() {\n\tch <- 42 // send\n}()\n\nval := <-ch // receive\nfmt.Println(val) // 42",
		"end: This is the end\n\nthis is all of what ive learned so far\nif you have any questions just ask me on github as a issue",
	}

	currentStep := 0
	stepDisplay := widget.NewLabel(steps[currentStep])
	stepDisplay.Wrapping = fyne.TextWrapWord

	gopherImage := canvas.NewImageFromFile("gopher.png")
	gopherImage.FillMode = canvas.ImageFillContain
	gopherImage.SetMinSize(fyne.NewSize(160, 200))

	progress := widget.NewProgressBar()
	progress.SetValue(float64(currentStep+1) / float64(len(steps)))

	nextButton := widget.NewButton("Next", nil)
	prevButton := widget.NewButton("Back", nil)
	prevButton.Disable()
	exitButton := widget.NewButton("Exit", func() {
		myApp.Quit()
	})





	nextButton.OnTapped = func() {
		if currentStep < len(steps)-1 {
			currentStep++
			stepDisplay.SetText(steps[currentStep])
			progress.SetValue(float64(currentStep+1) / float64(len(steps)))
			prevButton.Enable()
			if currentStep == 0 {
				gopherImage.Show()
			} else {
				gopherImage.Hide()
			}
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
			if currentStep == 0 {
				gopherImage.Show()
			} else {
				gopherImage.Hide()
			}
		}
		if currentStep == 0 {
			prevButton.Disable()
		}
	}

	buttonLayout := container.NewHBox(prevButton, nextButton, exitButton)
	mainLayout := container.NewVBox(
		widget.NewLabelWithStyle("Tour of Go by Jackson", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewCenter(gopherImage),
		stepDisplay,
		buttonLayout,
		progress,
	)

	myWindow.SetContent(mainLayout)
	myWindow.ShowAndRun()
}
