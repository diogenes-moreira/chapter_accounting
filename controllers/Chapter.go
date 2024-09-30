package controllers

import (
	"argentina-tresury/db"
	"argentina-tresury/model"
	"argentina-tresury/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const chapterPath = "/chapters"
const chapterPathId = chapterPath + "/{id}"

func RegisterChapterRoutesOn(r *mux.Router) {
	r.HandleFunc(chapterPath, CreateChapter).Methods("POST")
	r.HandleFunc(chapterPath, GetChapters).Methods("GET")
	r.HandleFunc(chapterPathId, GetChapter).Methods("GET")
	r.HandleFunc(chapterPathId, UpdateChapter).Methods("PUT")
	r.HandleFunc(chapterPathId, DeleteChapter).Methods("DELETE")
	r.HandleFunc(chapterPath+"/view", GetChaptersView).Methods("GET")
}

func GetChaptersView(writer http.ResponseWriter, request *http.Request) {
	// TODO: Implement
}

// CreateChapter creates a new chapter
func CreateChapter(w http.ResponseWriter, r *http.Request) {
	var chapter model.Chapter
	if err := json.NewDecoder(r.Body).Decode(&chapter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := services.CreateChapter(&chapter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(chapter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetChapters returns all chapters
func GetChapters(w http.ResponseWriter, r *http.Request) {
	chapters, err := services.GetChapters()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(chapters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetChapter returns a chapter by id
func GetChapter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chapter, err := services.GetChapter(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(chapter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateChapter updates a chapter by id
func UpdateChapter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var chapter model.Chapter
	if err := json.NewDecoder(r.Body).Decode(&chapter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chapter.ID = uint(id)
	err = services.UpdateChapter(&chapter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(chapter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteChapter deletes a chapter by id
func DeleteChapter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = db.DB.Delete(&model.Chapter{}, id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
