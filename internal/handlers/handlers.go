package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHtml(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	convertedData := service.TextOrMorse(data)
	newFile, err := os.Create(time.Now().UTC().Format("20060102_150405"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()
	_, err = newFile.WriteString(convertedData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(convertedData))

}
