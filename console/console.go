package console

import (
	"cardgame/constants"
	"fmt"
)

func CustomPrintLn(color, message string, a ...interface{}) {
	fmt.Print(color)
	fmt.Printf(message, a...)
	fmt.Print(constants.Reset)
	fmt.Println()
}

// Console an error message with red color.
func Error(message string, a ...interface{}) {
	CustomPrintLn(constants.Red, message, a...)
}

// Console a warning message with yellow color.
func Warn(message string, a ...interface{}) {
	CustomPrintLn(constants.Yellow, message, a...)
}

// Console a success message with green color.
func Success(message string, a ...interface{}) {
	CustomPrintLn(constants.Green, message, a...)
}

// Console a prompt with blue color.
func Prompt(message string, a ...interface{}) {
	CustomPrintLn(constants.Blue, message, a...)
}

// Console a prompt with blue color.
func Info(message string, a ...interface{}) {
	CustomPrintLn(constants.Magenta, message, a...)
}
