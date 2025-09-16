package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func (c *FC) sendLocalData(w http.ResponseWriter, r *http.Request) {
	initDB()
	log.Print("getting post request ", r.Method)
	if r.Method != "POST" {
		log.Print("error bad method", r.Method)
		return
	}
	localObj := getDataLocal(c.db)
	json.NewEncoder(w).Encode(localObj)
	log.Print("file send to client ")
}

func (c *FC) addDataToDB(w http.ResponseWriter, r *http.Request) {
	// Добавляем CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*") //c
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var fileData dFile
	err = json.Unmarshal(body, &fileData)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	var obj locateFile //MODEL INITVALUe
	obj.Name = fileData.Name
	obj.InCloud = "false"
	obj.TypePoint = ".doc"
	obj.LocateInCloud = ""
	obj.Locate = ""
	addDataLocate(c.db, obj)
	log.Printf("Received file metadata: %+v", fileData)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Metadata received"))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Парсим multipart form
	err := r.ParseMultipartForm(32 << 20) // 32 MB max memory
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Получаем файл из формы
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Создаем файл на сервере
	filePath := filepath.Join(uploadDir, header.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error creating file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Копируем содержимое файла
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("File uploaded successfully: %s", header.Filename)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully: " + header.Filename))
}
