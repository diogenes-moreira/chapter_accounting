package controllers

import (
	"argentina-tresury/model"
	"argentina-tresury/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const companionPath = "/api/companions"
const companionPathId = companionPath + "/{id}"

func RegisterCompanionRoutesOn(r *mux.Router) {
	r.HandleFunc("/companions/view", getCompanionsView).Methods("GET")
	r.HandleFunc(companionPath+"/exaltation", createExaltation).Methods("POST")
	r.HandleFunc(companionPath+"/affiliation", createCompanionAffiliation).Methods("POST")
	r.HandleFunc(companionPath, createCompanion).Methods("POST")
	r.HandleFunc(companionPath, getCompanions).Methods("GET")
	r.HandleFunc(companionPathId, getCompanion).Methods("GET")
	r.HandleFunc(companionPathId, updateCompanion).Methods("PUT")
	r.HandleFunc(companionPathId, deleteCompanion).Methods("DELETE")
}

type exaltation struct {
	Companion  *model.Companion `json:"companion"`
	IsHonorary bool             `json:"is_honorary"`
}

func createCompanionAffiliation(w http.ResponseWriter, r *http.Request) {
	var e exaltation

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chapter, err := processChapter(w, r)
	if err != nil {
		return
	}
	err = services.CreateCompanionAffiliation(e.Companion, e.IsHonorary, chapter)
	processCreation(w, err)
}

func processCreation(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(`{"status":"ok"}`))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func processChapter(w http.ResponseWriter, r *http.Request) (*model.Chapter, error) {
	chapterId, err := strconv.Atoi(r.Header.Get("chapter_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	chapter, err := services.GetChapter(uint(chapterId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	return chapter, nil
}

func createExaltation(w http.ResponseWriter, r *http.Request) {
	var e exaltation

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chapter, err := processChapter(w, r)
	if err != nil {
		return
	}
	err = services.CreateExaltation(e.Companion, e.IsHonorary, chapter)
	processCreation(w, err)
}

func getCompanionsView(w http.ResponseWriter, r *http.Request) {
	templateCompanions, err := parseTemplate("companions.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = executeTemplate(w, r, templateCompanions, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// createCompanion creates a new companion
func createCompanion(w http.ResponseWriter, r *http.Request) {
	var companion model.Companion
	if err := json.NewDecoder(r.Body).Decode(&companion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := services.CreateCompanion(&companion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(companion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getCompanions returns all companions
func getCompanions(w http.ResponseWriter, r *http.Request) {
	companions, err := services.GetCompanions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(companions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getCompanion returns a companion by id
func getCompanion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	companion, err := services.GetCompanion(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(companion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// updateCompanion updates a companion by id
func updateCompanion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var companion model.Companion
	if err := json.NewDecoder(r.Body).Decode(&companion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	companion.ID = uint(id)
	err = services.UpdateCompanion(&companion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(companion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// deleteCompanion deletes a companion by id
func deleteCompanion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = model.DB.Delete(&model.Companion{}, id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
