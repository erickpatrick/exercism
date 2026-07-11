package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	bars := strings.Repeat("*", numStarsPerLine)
	return bars + "\n" + welcomeMsg + "\n" + bars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	replacedString := strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(replacedString)
}
