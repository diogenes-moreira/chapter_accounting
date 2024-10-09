package controllers

import (
	"argentina-tresury/services"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// REMOVE IT
	//if username == "admin" && password == "admin" {
	//	services.CreateUser("admin", "admin")
	//}

	user := services.ValidateUser(username, password)
	if user == nil {

		if strings.Contains(r.URL.Path, "api") {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	token, err := services.GenerateToken(user)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(10 * time.Minute),
	})

	// REMOVE IT get CHAPTER ID and current
	http.Redirect(w, r, "/affiliations/view", http.StatusSeeOther)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-10 * time.Minute)})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func SecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAPI := strings.Contains(r.URL.Path, "api")
		if !isAPI && isPublicPath(r) {
			next.ServeHTTP(w, r)
			return
		}
		c, err, ret := extractCookie(w, r, isAPI)
		if ret {
			return
		}
		tokenStr := c.Value
		claim, err := services.ValidateToken(tokenStr)
		if err != nil {
			if !isAPI {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		r.Header.Set("chapter_id", strconv.Itoa(int(claim.ChapterId)))
		next.ServeHTTP(w, r)
	})
}

func extractCookie(w http.ResponseWriter, r *http.Request, isAPI bool) (*http.Cookie, error, bool) {
	c, err := r.Cookie("token")
	if err != nil {
		if !isAPI {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return nil, nil, true
		}
		if errors.Is(err, http.ErrNoCookie) {
			http.Error(w, "No token", http.StatusUnauthorized)
			return nil, nil, true
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return nil, nil, true
	}
	return c, err, false
}

func isPublicPath(r *http.Request) bool {
	return r.URL.Path == "/login" ||
		r.URL.Path == "/logout" ||
		r.URL.Path == "/" || strings.Contains(r.URL.Path, "/js/") || strings.Contains(r.URL.Path, "/css/")
}

func InitServer() {
	r := mux.NewRouter()
	InitView(r)
	r.Use(ReuseBody)
	RegisterIndex(r)
	r.HandleFunc("/login", HandleLogin).Methods("POST")
	r.HandleFunc("/logout", HandleLogout).Methods("GET")
	r.Use(SecurityMiddleware)
	RegisterIndex(r)
	RegisterAffiliationRoutesOn(r)
	RegisterBrotherRoutesOn(r)
	RegisterChapterRoutesOn(r)
	RegisterMovementTypeRoutesOn(r)
	RegisterPeriodRoutesOn(r)
	RegisterTreasuryRoutesOn(r)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
