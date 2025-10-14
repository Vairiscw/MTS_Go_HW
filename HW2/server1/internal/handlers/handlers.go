package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"server1/internal/config"
	"server1/internal/models"
	"server1/internal/sevices"
	"time"
)

func Version(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/version" {
		http.NotFound(w, r)
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	fmt.Fprint(w, config.Version)
}

func Decode(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/decode" {
		http.NotFound(w, r)
	}
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	var encodedJSON models.Encode
	if err := json.NewDecoder(r.Body).Decode(&encodedJSON); err != nil {
		http.Error(w, "Incorrect JSON", http.StatusBadRequest)
	}

	str := encodedJSON.InputString
	decodedBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		http.Error(w, "Decoding error", http.StatusBadRequest)
		return
	}

	resp := models.Decode{
		OutputString: string(decodedBytes),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func HardOp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	time.Sleep(time.Duration(sevices.GetRandomInt(10)+10) * time.Second)
	if sevices.GetRandomInt(2)%2 == 0 {
		http.Error(w, "Internal Server Error", sevices.GetRandomInt(10)+500)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
