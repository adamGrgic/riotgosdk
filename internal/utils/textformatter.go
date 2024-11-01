package utils

import "fmt"

const (
	GreenOpenTag  = "\033[32"
	GreenCloseTag = "\033[0m"
)

func FmtTextColor(text string, color string) string {
	formattedText := fmt.Sprintf("%s%s%s", GreenOpenTag, text, GreenCloseTag)
	return formattedText
}
