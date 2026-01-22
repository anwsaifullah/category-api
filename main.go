package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{
	{ID: 1, Name: "Cardio", Description: "Latihan untuk meningkatkan kesehatan jantung dan pembakaran kalori melalui aktivitas aerobik"}, {ID: 2, Name: "Calisthenics", Description: "Latihan menggunakan berat badan sendiri untuk membangun kekuatan dan fleksibilitas"}, {ID: 3, Name: "Gym", Description: "Latihan di pusat kebugaran dengan berbagai alat fitness dan beban"}, {ID: 4, Name: "Yoga", Description: "Latihan yang menggabungkan postur tubuh, pernapasan, dan meditasi untuk keseimbangan fisik dan mental"}, {ID: 5, Name: "Pilates", Description: "Latihan low-impact yang fokus pada penguatan core, fleksibilitas, dan postur tubuh"}, {ID: 6, Name: "CrossFit", Description: "Program latihan intensitas tinggi yang menggabungkan angkat beban, cardio, dan gimnastik"}, {ID: 7, Name: "Swimming", Description: "Olahraga air yang melatih seluruh otot tubuh dengan dampak rendah pada sendi"}, {ID: 8, Name: "Running", Description: "Aktivitas lari untuk meningkatkan stamina dan kesehatan kardiovaskular"}, {ID: 9, Name: "Cycling", Description: "Bersepeda untuk melatih kekuatan kaki dan daya tahan kardio"}, {ID: 10, Name: "Boxing", Description: "Olahraga bela diri yang melatih kekuatan, kecepatan, dan koordinasi"}, {ID: 11, Name: "Martial Arts", Description: "Berbagai seni bela diri tradisional untuk pertahanan diri dan disiplin mental"}, {ID: 12, Name: "Zumba", Description: "Latihan aerobik yang menggabungkan gerakan tari dengan musik latin dan internasional"}, {ID: 13, Name: "HIIT", Description: "High-Intensity Interval Training untuk pembakaran lemak dan peningkatan metabolisme"}, {ID: 14, Name: "Strength Training", Description: "Latihan kekuatan menggunakan beban untuk membangun massa otot"}, {ID: 15, Name: "Stretching", Description: "Latihan peregangan untuk meningkatkan fleksibilitas dan mencegah cedera"},
}

func main() {
	http.HandleFunc("/categories", categoryHandler)
	http.HandleFunc("/categories/", categoryByIdHandler)

	fmt.Println("Server is running..")
	http.ListenAndServe(":8080", nil)
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(categories)

	case "POST":
		var newCategory Category
		err := json.NewDecoder(r.Body).Decode(&newCategory)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		newCategory.ID = len(categories) + 1
		categories = append(categories, newCategory)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newCategory)
	}
}

func categoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/categories/"))

	for i, cat := range categories {
		if id == cat.ID {
			switch r.Method {
			case "PUT":
				var updatedCategory Category
				err := json.NewDecoder(r.Body).Decode(&updatedCategory)
				if err != nil {
					http.Error(w, "Invalid request", http.StatusBadRequest)
					return
				}
				updatedCategory.ID = id
				categories[i] = updatedCategory

				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(updatedCategory)
			case "GET":
				json.NewEncoder(w).Encode(cat)
			case "DELETE":
				categories = append(categories[:i], categories[i+1:]...)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Category deleted successfully",
				})
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}
