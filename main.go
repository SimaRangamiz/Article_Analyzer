package main

import ("net/http"
        "encoding/json"
		"time"
		"context"
		"strings"
		"strconv"
		"bufio"
)

func Analyzer_Request(w http.ResponseWriter, r *http.Request) {
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

	top_num_Str := r.FormValue("top_num")
	top_num := 5 
	if top_num_Str != "" {
		parsed_num, err := strconv.Atoi(top_num_Str)
		if err == nil && parsed_num > 0 {
			top_num = parsed_num 
		}
	}

	var title string
	var body_text []string
	scanner := bufio.NewScanner(file)
	First_Line := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		if First_Line {
			if line != "" {
				title = line
				First_Line = false
			}
			continue
		}
		
		body_text = append(body_text, line)
	}

	if err := scanner.Err(); err != nil {
		http.Error(w, "Error reading", http.StatusInternalServerError)
		return
	}

	body := strings.Join(body_text, "\n")

	if title == "" || strings.TrimSpace(body) == "" {
		http.Error(w, "Invalid file structure", http.StatusBadRequest)
		return
	}

	top_tags := extract(body, top_num)
	
	New_Article := Article{
		Title:     title,
		Body:      body,
		Tags:      top_tags,
		Created_Time: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = DB.InsertOne(ctx, New_Article)
	if err != nil {
		http.Error(w, "Failed to save to DB: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := Article_Response{
			Title:        title, 
			Tags:         top_tags,
			Created_Time: New_Article.Created_Time, 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func main() {

	Connect_DB()

	http.HandleFunc("/analyze", Analyzer_Request)
	err := http.ListenAndServe(":8085", nil)
	if err != nil{
		return
	}
}