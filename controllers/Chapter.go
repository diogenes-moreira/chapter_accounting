package controllers

import (
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
	r.HandleFunc(chapterPath+"/view", GetChaptersView).Methods("GET")
	r.HandleFunc(chapterPathId+"/affiliations/{period}", GetChaptersAffiliations).Methods("GET")
	r.HandleFunc(chapterPath, CreateChapter).Methods("POST")
	r.HandleFunc(chapterPath, GetChapters).Methods("GET")
	r.HandleFunc(chapterPathId, GetChapter).Methods("GET")
	r.HandleFunc(chapterPathId, UpdateChapter).Methods("PUT")
	r.HandleFunc(chapterPathId, DeleteChapter).Methods("DELETE")
}

func GetChaptersAffiliations(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	period, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	affiliations, err := services.GetChapterAffiliations(uint(id), uint(period))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(affiliations)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

}

func GetChaptersView(w http.ResponseWriter, r *http.Request) {
	templateChapters, err := parseTemplate("chapters.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = templateChapters.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	err = model.DB.Delete(&model.Chapter{}, id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
