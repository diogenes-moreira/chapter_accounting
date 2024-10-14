package controllers

import (
	"argentina-tresury/services"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	user := services.ValidateUser(username, password)
	if user == nil {

		if strings.Contains(r.URL.Path, "api") {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	token, err := services.GenerateToken(user)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, cookieFrom(token))
	http.Redirect(w, r, "/affiliations", http.StatusSeeOther)
}

func cookieFrom(token string) *http.Cookie {
	return &http.Cookie{
		Path:    "/",
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(10 * time.Minute),
	}
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Path:    "/",
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
		r.Header.Set("user-name", claim.Username)
		if err != nil {
			if !isAPI {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		r.Header.Set("chapter_id", strconv.Itoa(int(claim.ChapterId)))
		tokenStr, err = services.RefreshToken(claim)
		if err != nil {
			http.Error(w, "Failed to refresh token", http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, cookieFrom(tokenStr))
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
	r.Use(SecurityMiddleware)
	RegisterIndex(r)
	r.HandleFunc("/change-password", HandleChangePassword).Methods("POST")
	r.HandleFunc("/change-password", ChangePasswordView).Methods("GET")
	r.HandleFunc("/logout", HandleLogout).Methods("GET")
	RegisterAffiliationRoutesOn(r)
	RegisterBrotherRoutesOn(r)
	RegisterChapterRoutesOn(r)
	RegisterMovementTypeRoutesOn(r)
	RegisterPeriodRoutesOn(r)
	RegisterTreasuryRoutesOn(r)

	err := http.ListenAndServe(os.Getenv("LISTENER"), r)
	if err != nil {
		return
	}
}

func HandleChangePassword(writer http.ResponseWriter, request *http.Request) {
	user := request.Header.Get("user-name")
	err := services.ChangePassword(user, request.FormValue("password"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

func ChangePasswordView(w http.ResponseWriter, r *http.Request) {
	template, err := parseTemplate("changePassword.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
