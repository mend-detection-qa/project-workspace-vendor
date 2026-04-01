package lib

import "github.com/fatih/color"

// Print prints a message in green
func Print(msg string) {
	color.Green(msg)
}