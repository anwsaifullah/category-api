package main

import (
	"fmt"
	"net/http"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{
	{ID: 1, Name: "Cardio", Description: "Latihan untuk meningkatkan kesehatan jantung dan pembakaran kalori melalui aktivitas aerobik"}, {ID: 2, Name: "Calisthenics", Description: "Latihan menggunakan berat badan sendiri untuk membangun kekuatan dan fleksibilitas"}, {ID: 3, Name: "Gym", Description: "Latihan di pusat kebugaran dengan berbagai alat fitness dan beban"}, {ID: 4, Name: "Yoga", Description: "Latihan yang menggabungkan postur tubuh, pernapasan, dan meditasi untuk keseimbangan fisik dan mental"}, {ID: 5, Name: "Pilates", Description: "Latihan low-impact yang fokus pada penguatan core, fleksibilitas, dan postur tubuh"}, {ID: 6, Name: "CrossFit", Description: "Program latihan intensitas tinggi yang menggabungkan angkat beban, cardio, dan gimnastik"}, {ID: 7, Name: "Swimming", Description: "Olahraga air yang melatih seluruh otot tubuh dengan dampak rendah pada sendi"}, {ID: 8, Name: "Running", Description: "Aktivitas lari untuk meningkatkan stamina dan kesehatan kardiovaskular"}, {ID: 9, Name: "Cycling", Description: "Bersepeda untuk melatih kekuatan kaki dan daya tahan kardio"}, {ID: 10, Name: "Boxing", Description: "Olahraga bela diri yang melatih kekuatan, kecepatan, dan koordinasi"}, {ID: 11, Name: "Martial Arts", Description: "Berbagai seni bela diri tradisional untuk pertahanan diri dan disiplin mental"}, {ID: 12, Name: "Zumba", Description: "Latihan aerobik yang menggabungkan gerakan tari dengan musik latin dan internasional"}, {ID: 13, Name: "HIIT", Description: "High-Intensity Interval Training untuk pembakaran lemak dan peningkatan metabolisme"}, {ID: 14, Name: "Strength Training", Description: "Latihan kekuatan menggunakan beban untuk membangun massa otot"}, {ID: 15, Name: "Stretching", Description: "Latihan peregangan untuk meningkatkan fleksibilitas dan mencegah cedera"}, {ID: 16, Name: "Test category", Description: "lorep ipsum"},
}

func main() {
	http.HandleFunc("/categories", categoryHandler)
	http.HandleFunc("/categories/", categoryByIdHandler)

	fmt.Println("Server is running..")
	http.ListenAndServe(":8080", nil)
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {

}

func categoryByIdHandler(w http.ResponseWriter, r *http.Request) {

}
