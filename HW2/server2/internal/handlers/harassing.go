package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"server2/internal/models"
	"time"
)

func HarassingVersion(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080/version")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func HarassingJSON(w http.ResponseWriter, r *http.Request) {
	inputString := models.Encode{
		InputString: "aGVsbG8=",
	}
	jsonString, err := json.Marshal(inputString)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(
		"http://localhost:8080/decode",
		"application/json",
		bytes.NewBuffer(jsonString),
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func HarassingRandom(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/hard-op", nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("request timed out after 15s")
		} else {
			fmt.Println("failed to do http request")
		}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "failed to read response", http.StatusInternalServerError)
		return
	}

	fmt.Printf(string(body))
	fmt.Printf("Status: %d\n", resp.StatusCode)
}
