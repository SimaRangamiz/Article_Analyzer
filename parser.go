package main

import (
	"bufio"
	"errors"
	"mime/multipart"
	"strings"
)

func ParseArticle(file multipart.File) (string, string, error) {
	var title string
	var bodyLines []string
	scanner := bufio.NewScanner(file)
	firstLine := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if firstLine {
			if line != "" {
				title = line
				firstLine = false
			}
			continue
		}
		bodyLines = append(bodyLines, line)
	}

	if err := scanner.Err(); err != nil {
		return "", "", err
	}

	body := strings.Join(bodyLines, "\n")

	if title == "" || strings.TrimSpace(body) == "" {
		return "", "", errors.New("invalid file structure: title or body empty")
	}

	return title, body, nil
}