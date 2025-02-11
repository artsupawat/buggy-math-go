package main

import (
	"encoding/json"
	"net/http"

	"github.com/artsupawat/buggy-math-go/math"
)

type AddRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type AddResponse struct {
	Result int `json:"result"`
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var req AddRequest
	json.NewDecoder(r.Body).Decode(&req)

	result := math.Add(req.X, req.Y)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AddResponse{Result: result})
}

type SubtractRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type SubtractResponse struct {
	Result int `json:"result"`
}

func substractHandler(w http.ResponseWriter, r *http.Request) {
	var req SubtractRequest
	json.NewDecoder(r.Body).Decode(&req)

	result := math.Subtract(req.X, req.Y)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SubtractResponse{Result: result})
}

type MultiplyRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MultiplyResponse struct {
	Result int `json:"result"`
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	var req MultiplyRequest
	json.NewDecoder(r.Body).Decode(&req)

	result := math.Multiply(req.X, req.Y)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MultiplyResponse{Result: result})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/substract", substractHandler)
	http.HandleFunc("/multiply", multiplyHandler)

	http.ListenAndServe(":8080", nil)
}
