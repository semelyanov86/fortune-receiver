package main

import (
	"context"
	"strings"
)

type GetWebRequest interface {
	FetchBytes(ctx context.Context, typeFort int) ([]byte, error)
}

type FortuneResult struct {
	Content string
}

func (result FortuneResult) splitString(maxLength int) string {
	var output string
	lines := strings.Split(result.Content, "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		lineLength := 0
		for _, word := range words {
			wordLength := len(word)
			if lineLength+wordLength+1 > maxLength {
				output += "\n" + word + " "
				lineLength = 0
			} else {
				output += word + " "
				lineLength += wordLength + 1
			}
		}
		output += "\n"
	}
	output = strings.Trim(output, "\n")
	output = strings.Trim(output, " ")
	return strings.Trim(output, " \n")
}
