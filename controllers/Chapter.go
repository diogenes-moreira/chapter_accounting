package controllers

import (
	"argentina-tresury/model"
	"argentina-tresury/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const chapterPath = "/api/chapters"
const chapterPathId = chapterPath + "/{id:[0-9]+}"

func RegisterChapterRoutesOn(r *mux.Router) {
	r.HandleFunc("/chapters", GetChaptersView).Methods("GET")
	r.HandleFunc("/great-chapter", getGreatChapterView).Methods("GET")
	r.HandleFunc(chapterPath+"/treasury", GetTreasury).Methods("GET")
	r.HandleFunc(chapterPath+"/affiliations", GetChaptersAffiliations).Methods("GET")
	r.HandleFunc(chapterPath+"/movement", createChapterMovement).Methods("POST")
	r.HandleFunc(chapterPath+"/great-chapter", getGreatChapterStatus).Methods("GET")
	r.HandleFunc(chapterPath+"/deposit", createDeposit).Methods("POST")
	r.HandleFunc(chapterPath+"/update-installment", updateInstallments).Methods("POST")
	r.HandleFunc(chapterPath, CreateChapter).Methods("POST")
	r.HandleFunc(chapterPath, GetChapters).Methods("GET")
	r.HandleFunc(chapterPathId, GetChapter).Methods("GET")
	r.HandleFunc(chapterPathId, UpdateChapter).Methods("PUT")
	r.HandleFunc(chapterPathId, DeleteChapter).Methods("DELETE")
}

type UpdateInstallments struct {
	Amount             float64 `json:"amount"`
	GreatChapterAmount float64 `json:"great_chapter_amount"`
	TypeId             uint    `json:"type_id"`
}

type Deposit struct {
	Installments []uint64 `json:"installments"`
}

type ChapterMovement struct {
	Amount      float64 `json:"amount"`
	Date        ISODate `json:"date"`
	Receipt     string  `json:"receipt"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
}

func updateInstallments(writer http.ResponseWriter, request *http.Request) {
	ui := &UpdateInstallments{}
	if err := json.NewDecoder(request.Body).Decode(ui); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(request.Header.Get("chapter_id"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = services.UpdateInstallments(uint(id), ui.Amount, ui.GreatChapterAmount, ui.TypeId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write([]byte(`{ "status":"Installments updated" }`))

}

func createDeposit(writer http.ResponseWriter, request *http.Request) {
	deposit := &Deposit{}
	if err := json.NewDecoder(request.Body).Decode(deposit); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(request.Header.Get("chapter_id"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = services.CreateDeposit(uint(id), deposit.Installments)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write([]byte(`{ "status":"Deposit created" }`))
}

func getGreatChapterStatus(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.Header.Get("chapter_id"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	chapter, err := services.GetGreatChapterStatus(uint(id))
	if chapter == nil {
		http.Error(writer, "Chapter not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(chapter)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getGreatChapterView(writer http.ResponseWriter, request *http.Request) {
	templateGreatChapter, err := parseTemplate("greatChapter.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = executeTemplate(writer, request, templateGreatChapter, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createChapterMovement(writer http.ResponseWriter, request *http.Request) {
	movement := &ChapterMovement{}
	if err := json.NewDecoder(request.Body).Decode(movement); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(request.Header.Get("chapter_id"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.CreateChapterMovement(uint(id), movement.Amount, movement.Receipt, movement.Date.Time, movement.Type,
		movement.Description)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write([]byte(`{ "status":"Movement created" }`))
}

func GetTreasury(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.Header.Get("chapter_id"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	treasury, err := services.GetChapter(uint(id))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(treasury.TreasurerRollingBalance)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetChaptersAffiliations(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.Header.Get("chapter_id"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	affiliations, err := services.GetChapterAffiliations(uint(id))
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
	err = executeTemplate(w, r, templateChapters, nil)
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
