package controllers

import (
	"argentina-tresury/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

const affiliationPath = "/api/affiliations"
const affiliationPathId = affiliationPath + "/{id:[0-9]+}"

func RegisterAffiliationRoutesOn(r *mux.Router) {
	r.HandleFunc("/affiliations", getAffiliationsView).Methods("GET")
	r.HandleFunc(affiliationPath+"/payment", createPayment).Methods("POST")
	r.HandleFunc(affiliationPath+"/expenses", createAffiliationExpense).Methods("POST")
	r.HandleFunc(affiliationPath, CreateAffiliation).Methods("POST")
	r.HandleFunc(affiliationPath, GetAffiliations).Methods("GET")
	r.HandleFunc(affiliationPathId, GetAffiliation).Methods("GET")
	r.HandleFunc(affiliationPathId, UpdateAffiliation).Methods("PUT")
	r.HandleFunc(affiliationPathId, DeleteAffiliation).Methods("DELETE")

}

type Payment struct {
	Amount        float64 `json:"amount"`
	Receipt       string  `json:"receipt"`
	Date          ISODate `json:"date"`
	AffiliationId uint    `json:"affiliation_id"`
}

type Expense struct {
	Payment
	Type        string `json:"type"`
	Description string `json:"description"`
}

func createAffiliationExpense(w http.ResponseWriter, r *http.Request) {
	expense := &Expense{}
	if err := json.NewDecoder(r.Body).Decode(expense); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := services.CreateAffiliationExpense(expense.AffiliationId, expense.Amount, expense.Receipt, expense.Date.Time,
		expense.Type, expense.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(`{ "status":"Expenses created" }`))
	if err != nil {
		return
	}
}
func createPayment(w http.ResponseWriter, r *http.Request) {
	payment := &Payment{}
	if err := json.NewDecoder(r.Body).Decode(payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := services.CreatePayment(payment.AffiliationId, payment.Amount, payment.Receipt, payment.Date.Time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(`{ "status":"Payment created" }`))
	if err != nil {
		return
	}
}

func DeleteAffiliation(writer http.ResponseWriter, request *http.Request) {
	// TODO: Implement
}

func UpdateAffiliation(writer http.ResponseWriter, request *http.Request) {
	// TODO: Implement
}

func getAffiliationsView(w http.ResponseWriter, r *http.Request) {
	templateAffiliations, err := parseTemplate("affiliations.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = executeTemplate(w, r, templateAffiliations, &View{Title: "Afiliaciones"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateAffiliation(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}

func GetAffiliations(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}

func GetAffiliation(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}
