package main

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"
	"unicode"
)

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, str)
}

func RemoveHTMLComments(content []byte) []byte {
	htmlcmt := regexp.MustCompile(`<!--[^>]*-->`)
	return htmlcmt.ReplaceAll(content, []byte(""))
}

func MinifyHTML(html []byte) string {
	minifiedHTML := ""
	scanner := bufio.NewScanner(bytes.NewReader(RemoveHTMLComments(html)))
	for scanner.Scan() {
		lineTrimmed := strings.TrimSpace(scanner.Text())
		minifiedHTML += lineTrimmed
		if len(lineTrimmed) > 0 {
			minifiedHTML += ""
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return minifiedHTML
}
