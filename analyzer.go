package main

import ("strings"
        "github.com/bbalet/stopwords"
)

func extract(body string, top_num int) []string {
	lowerbody := strings.ToLower(body)
	cleanedtext := stopwords.CleanString(lowerbody, "en", false)
	replacer := strings.NewReplacer(".", "", ",", "", "!", "", "?", "")
	cleanbody := replacer.Replace(cleanedtext)

	words := strings.Fields(cleanbody)

	counts := make(map[string]int)
	for i := 0; i < len(words); i++{
		word := words[i]
		
		if len(word) > 1{
			counts[word] = counts[word] + 1
		}
	}

	var tags []string

	for i := 0; i < top_num; i++{
		max_count := 0
		best_word := ""

		for word, count := range counts {
			if count > max_count || (count == max_count && (best_word == "" || word < best_word)) {
				max_count = count
				best_word = word
			}
		}
		if best_word != "" {
			tags = append(tags, best_word)
			delete(counts, best_word)
		}
	}

	return tags
} 