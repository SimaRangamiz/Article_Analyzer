package main

import ("strings"
        "github.com/bbalet/stopwords"
)

func extract(body string, topNum int) []string {
	lowerBody := strings.ToLower(body)
	cleanedText := stopwords.CleanString(lowerBody, "en", false)
	replacer := strings.NewReplacer(".", "", ",", "", "!", "", "?", "")
	cleanBody := replacer.Replace(cleanedText)

	words := strings.Fields(cleanBody)

	counts := make(map[string]int)
	for i := 0; i < len(words); i++{
		word := words[i]
		
		if len(word) > 1{
			counts[word] = counts[word] + 1
		}
	}

	var tags []string

	for i := 0; i < topNum; i++{
		maxCount := 0
		bestWord := ""

		for word, count := range counts {
			if count > maxCount || (count == maxCount && (bestWord == "" || word < bestWord)) {
				maxCount = count
				bestWord = word
			}
		}
		if bestWord != "" {
			tags = append(tags, bestWord)
			delete(counts, bestWord)
		}
	}

	return tags
} 