Jackson's Tour of Go

A lightweight, GUI-based introduction to the Go programming language. Inspired by the official Tour of Go (https://go.dev/tour).

## Overview

This app walks through 20 Go concepts one step at a time using a simple Fyne window with Back/Next navigation and a progress bar.

Concepts covered:

- Packages
- fmt
- Variables and Constants
- Integers, Floats, Strings, Booleans
- Functions
- If / Else
- For Loops
- Switch
- Arrays and Slices
- Maps
- Structs
- Pointers
- Error Handling
- Defer
- Goroutines and Channels

## Live Web Version

A browser-based version is available at: https://jackson-tour-of-go.netlify.app

## How to Run

Prerequisites: Go installed (`sudo pacman -S go` on Arch).

```bash
go run tour.go
```

Or build a binary:

```bash
go build -o tour-of-go tour.go
./tour-of-go
```

## Built With

- Go + Fyne (GUI toolkit)
- Vim on Arch Linux

## Sources

- https://go.dev/tour
- https://docs.fyne.io/
