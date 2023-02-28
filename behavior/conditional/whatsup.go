package conditional

import (
	"fmt"
	"rhino-bot/behavior"
	"strings"
)

type WhatsupConditional struct {
	upWord string
}

func findUpWord(message string) string {
	replacer := strings.NewReplacer(",", "", ".", "", ";", "", "?", "", "!", "")
	replacedMessage := replacer.Replace(message)
	words := strings.Fields(replacedMessage)

	for _, word := range words {
		if strings.HasPrefix(word, "up") && word != "up" {
			return word
		}
	}

	return ""
}

func (wc *WhatsupConditional) Handle(message string) (string, error) {
	return fmt.Sprintf("What's %v?", wc.upWord), nil
}

func (wc *WhatsupConditional) ShouldHandle(message string) bool {
	upWord := findUpWord(message)
	wc.upWord = upWord
	if upWord != "" {
		return behavior.GenerateBool(50)
	}

	return false
}