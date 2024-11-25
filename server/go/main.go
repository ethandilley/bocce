package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

type Food struct {
	Name    string `json:"name"`
	Protein int    `json:"protein"`
	Carb    int    `json:"carb"`
	Fat     int    `json:"fat"`
}

var (
	Foods = []Food{}
	mutex sync.Mutex // To handle concurrent access
)

func main() {
	http.HandleFunc("/api/v1/food", foodHandler)
	http.HandleFunc("/api/v1/food/", foodWithNameHandler)

	log.Println("Server is running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Handles the /api/v1/food endpoint
func foodHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getFood(w)
	case http.MethodPost:
		createFood(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getFood(w http.ResponseWriter) {
	mutex.Lock()
	defer mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Foods)
}

func createFood(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var food Food
	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for _, existingFood := range Foods {
		if food.Name == existingFood.Name {
			http.Error(w, "The food already exists", http.StatusBadRequest)
			return
		}
	}

	Foods = append(Foods, food)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(food)
}

// Handles the /api/v1/food/{name} endpoint
func foodWithNameHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the {name} parameter
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/food/")
	if path == "" || strings.Contains(path, "/") {
		http.NotFound(w, r)
		return
	}
	foodName := path

	switch r.Method {
	case http.MethodGet:
		getFoodWithName(w, foodName)
	case http.MethodDelete:
		deleteFood(w, foodName)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getFoodWithName(w http.ResponseWriter, name string) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, food := range Foods {
		if food.Name == name {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(food)
			return
		}
	}
	http.Error(w, "Food not found", http.StatusNotFound)
}

func deleteFood(w http.ResponseWriter, name string) {
	mutex.Lock()
	defer mutex.Unlock()

	for i, food := range Foods {
		if food.Name == name {
			Foods = append(Foods[:i], Foods[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(food)
			return
		}
	}
	http.Error(w, "Food not found", http.StatusNotFound)
}
