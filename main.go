package main

import ("net/http"
        "encoding/json"
		"time"
		"context"	
)

func Analyzer_Request(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	var request Article_Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "JSON format error", http.StatusBadRequest)
		return
	}

	top_tags := extract(request.Text, request.Top_Num)
	
	New_Article := Article{
		Title:     request.Title,
		Body:      request.Text,
		Tags:      top_tags,
		Created_Time: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = DB.InsertOne(ctx, New_Article)
	if err != nil {
		http.Error(w, "Failed to save to database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := Article_Response{
			Title:        request.Title, 
			Tags:         top_tags,
			Created_Time: New_Article.Created_Time, 
	}

	w.Header().Set("Content-Type", "application/json")
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