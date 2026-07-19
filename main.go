package main

import ("net/http"
        "encoding/json"
		"time"
		"strconv"
)

func AnalyzerRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Please upload a file using the key 'file'", http.StatusBadRequest)
		return
	}
	defer file.Close()

	topNumStr := r.FormValue("topNum")
	topNum := 5 
	if topNumStr != "" {
		parsedNum, err := strconv.Atoi(topNumStr)
		if err == nil && parsedNum > 0 {
			topNum = parsedNum 
		}
	}

	title, body, err := ParseArticle(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	topTags := extract(body, topNum)
	
	NewArticle := Article{
		Title:     title,
		Body:      body,
		Tags:      topTags,
		CreatedTime: time.Now(),
	}

	err = SaveArticle(NewArticle)
	if err != nil {
		http.Error(w, "Failed to save to DB", http.StatusInternalServerError)
		return
	}

	response := ArticleResponse{
			Title:        title, 
			Tags:         topTags,
			CreatedTime: NewArticle.CreatedTime, 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func main() {

	ConnectDB()

	http.HandleFunc("/analyze", AnalyzerRequest)
	err := http.ListenAndServe(":8085", nil)
	if err != nil{
		return
	}
}