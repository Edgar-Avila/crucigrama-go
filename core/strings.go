package core

import (
	"errors"
	"strings"
	"unicode"

	textrank "github.com/DavidBelicza/TextRank/v2"
)

func MostImportantWords(text string, wordCount int, maxSize int) ([]string, error) {
	tr := textrank.NewTextRank()
	lang := textrank.NewDefaultLanguage()
	algo := textrank.NewChainAlgorithm()
	rule := textrank.NewDefaultRule()
	tr.Populate(text, lang, rule)
	tr.Ranking(algo)
	words := textrank.FindSingleWords(tr)

	var importantWords []string
	for i := 0; i < len(words); i++ {
		word := words[i].Word
		if !IsAlpha(word) {
			word = ToAlpha(word)
			if len(word) == 0 {
				continue
			}
		}
		if len(word) > maxSize {
			continue
		}
		importantWords = append(importantWords, strings.ToUpper((word)))
		if len(importantWords) == wordCount {
			break
		}
	}

	if len(importantWords) < wordCount {
		return nil, errors.New("no se encontraron suficientes palabras")
	}

	return importantWords, nil
}

func IsAlpha(text string) bool {
	for _, r := range text {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func ToAlpha(text string) string {
	var result []rune
	for _, r := range text {
		if unicode.IsLetter(r) {
			result = append(result, r)
		}
	}
	return string(result)
}
