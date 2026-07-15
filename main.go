package main

import "net/http"
import "encoding/json"
import "time"


type Article_Request struct {
	Title  string `json:"title"`   
	Text   string `json:"text"`    
	Top_Num int    `json:"top_num"` 
}


type Article_Response struct {
	Title     string    `json:"title"`      
	Tags      []string  `json:"tags"`       
	Created_Time time.Time `json:"created_at"` 
}

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
	
	response := Article_Response{
		Title:        request.Title, 
		Tags:         top_tags,
		Created_Time: time.Now(), 
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/analyze", Analyzer_Request)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		return
	}
}