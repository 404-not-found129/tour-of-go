Jackson's Terminal Tour of Go
A lightweight, terminal-based introduction to the Go programming language. This project serves as a quick-reference guide and interactive learning tool, inspired by the official Tour of Go.

📝 Overview
This program prints out foundational Go concepts directly to the terminal, making it easy to read and understand the basics of the language at a glance. It is designed for beginners who want a quick refresher on Go syntax without leaving the command line.

Concepts Covered:

Variables: Understanding how to define and manipulate data.

Constants: Working with unchanging values.

Integers (int): Basic number types without decimals.

Functions: How to structure instructions (like func main() and custom functions).

Packages: A brief introduction to the default fmt package used for formatted I/O.

🚀 How to Run
Since this is a standard Go application, you can run it easily from your terminal.

Prerequisites:
You will need the Go compiler installed on your system. If you haven't installed it yet, you can grab it from your package manager (e.g., sudo pacman -S go).

Execution:

Clone or download this repository.

Open your terminal and navigate to the directory containing main.go.

Run the following command to compile and execute the program on the fly:

Bash
go run main.go
Alternatively, if you want to build an executable binary:

Bash
go build -o tour-of-go main.go
./tour-of-go
🛠️ Built With
Go - The programming language used.

Vim / Konsole - Developed in a Linux terminal environment.

💡 Motivation
This was built as a personal project to solidify core Go concepts and create an easy-to-read reference guide right in the terminal.
