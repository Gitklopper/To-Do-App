package main

import (
	"net/http"
)

//origin = Protocol + domain + port
//Allow-Origin with explicit port number allows frontend to recieve information
//browser allows frontend to recieve info from backend

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") //Frontend and Backend have different port, Frontend :3000
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH") //allows HTTP Methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") //allows JSON packages

		//without checking for options no patch request will be accepted
		if r.Method == "OPTIONS" {
			return
		}
		next(w, r)
	}
}

func main() {
	http.HandleFunc("/todo", corsMiddleware(todoHandler))
	http.HandleFunc("/todo/update", corsMiddleware(updateTodoHandler))
	http.ListenAndServe(":8080", nil)
}
