package console

import "fmt"

// ANSI color escape sequences
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Blue   = "\033[34m"
)

// Custom console prompts
func PrintStyledMessage(color, message string, a ...interface{}) {
	fmt.Print(color)
	fmt.Printf(message, a...)
	fmt.Print(Reset)
	fmt.Println()
}

// Error prints an error message with red color.
func Error(message string, a ...interface{}) {
	PrintStyledMessage(Red, message, a...)
}

// Warning prints a warning message with yellow color.
func Warning(message string, a ...interface{}) {
	PrintStyledMessage(Yellow, message, a...)
}

// Success prints a success message with green color.
func Success(message string, a ...interface{}) {
	PrintStyledMessage(Green, message, a...)
}

// Prompt prints a prompt message with blue color.
func Prompt(message string, a ...interface{}) {
	PrintStyledMessage(Blue, message, a...)
}
