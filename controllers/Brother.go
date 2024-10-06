package controllers

import (
	"argentina-tresury/model"
	"argentina-tresury/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

const brotherPath = "/brothers"
const brotherPathId = brotherPath + "/{id}"

func RegisterBrotherRoutesOn(r *mux.Router) {
	r.HandleFunc(brotherPath+"/view", GetBrothersView).Methods("GET")
	r.HandleFunc(brotherPath, CreateBrother).Methods("POST")
	r.HandleFunc(brotherPath, GetBrothers).Methods("GET")
	r.HandleFunc(brotherPathId, GetBrother).Methods("GET")
	r.HandleFunc(brotherPathId, UpdateBrother).Methods("PUT")
	r.HandleFunc(brotherPathId, DeleteBrother).Methods("DELETE")
}

func GetBrothersView(w http.ResponseWriter, r *http.Request) {
	templateBrothers, err := template.ParseFiles("templates/brothers.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = templateBrothers.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreateBrother creates a new brother
func CreateBrother(w http.ResponseWriter, r *http.Request) {
	var brother model.Brother
	if err := json.NewDecoder(r.Body).Decode(&brother); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := services.CreateBrother(&brother)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(brother)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetBrothers returns all brothers
func GetBrothers(w http.ResponseWriter, r *http.Request) {
	brothers, err := services.GetBrothers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(brothers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetBrother returns a brother by id
func GetBrother(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	brother, err := services.GetBrother(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(brother)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateBrother updates a brother by id
func UpdateBrother(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var brother model.Brother
	if err := json.NewDecoder(r.Body).Decode(&brother); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	brother.ID = uint(id)
	err = services.UpdateBrother(&brother)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(brother)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteBrother deletes a brother by id
func DeleteBrother(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = model.DB.Delete(&model.Brother{}, id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
