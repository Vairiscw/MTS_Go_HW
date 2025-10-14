package main

import (
	"net/http"
	"net/http/httptest"
	"server2/internal/handlers"
)

func main() {
	req, _ := http.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", "/test", nil)

	handlers.HarassingVersion(w, req)
	handlers.HarassingJSON(w, req2)
	handlers.HarassingRandom(w, req)

	http.ListenAndServe(":8081", nil)
}
