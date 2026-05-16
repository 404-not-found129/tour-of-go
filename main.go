package main

import "fmt"

//example of how to create a function 
func fmtinfo() {
	fmt.Println("fmt: fmt is the default package in go")
	fmt.Println("")
}

func main() {
	// all varibles and constants (importent to read)
	var intro string = "this is my tour of go by jackson in terminal"
		const description string = "this was made by 404-not-found129 on github learning go from the website tour of go https://go.dev/tour it will be easy to understand and easy to read"
		const varibles string = "varible: a varible (a number or letter that can be manipulized) can be defined with int"
		const interger string = "intergers: a interger (a number without a deciaml) can be defined with the int word"
		const editor string = `editor: a editor is very very importent in coding it its what makes you able to code there are a few editors i would reccomend 
	1.vim
	2.jetbrains
	3.vs-codium`
		const intergers string = "interger: a interger is a number without a decimal you can use intergers with int"
		const boolian string = "a bool aka boolian is a answer wit htrue or false"
		const funcitions string = `function: a function is like a instruction when doing

func main() {
    add(x, y int) int
    return x + y
}
fmt.Println(add(number, number) and then shows the answer)
`

	// what prints all the text
	fmt.Println(intro)
	fmt.Println()
	fmt.Println(description)
	fmt.Println()

	// prints the example function i made earlier
	fmtinfo()


	fmt.Println(interger)
	fmt.Println()
	fmt.Println(varibles)
	fmt.Println()
	fmt.Println(funcitions)
	fmt.Println()
	fmt.Println (intergers)
	fmt.Println()
	fmt.Println()
	fmt.Println (boolian) 
	fmt.Println (editor)
	
	
} 
