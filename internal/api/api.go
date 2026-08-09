package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	auth "scryv/GoVault/internal/auth"

	"github.com/joho/godotenv"
)

var tkk []byte
var ygap string

func ApiStart() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}

	tkk := []byte(os.Getenv("TKN_K"))
	ygap := os.Getenv("YGAP_K")

	log.Println(ygap)
	log.Println(tkk)
	router := http.NewServeMux()

	router.HandleFunc("/{$}", handleRoot)
	router.HandleFunc("/auth/login", handleLogin)
	router.HandleFunc("/auth/signup", handleSignup)
	router.HandleFunc("/login", login)
	router.HandleFunc("/signup", signup)
	router.HandleFunc("/error", erreur)

	log.Println("Listning")
	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_token") //looks if there is cookie access_token
	if err != nil {
		index := template.Must(template.ParseFiles("internal/api/web/templates/index.html"))
		index.Execute(w, nil)
		return
	}
	usr, err := auth.VerifyToken(cookie.Value) //validates cookie
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusSeeOther) //invalid goes to error
		return
	}
	log.Println(usr)
	log.Println("Sir There has been a cookie detected") //else will be able to log in
}

func login(w http.ResponseWriter, _ *http.Request) {
	login := template.Must(template.ParseFiles("internal/api/web/templates/login.html"))
	login.Execute(w, nil)
}

func signup(w http.ResponseWriter, _ *http.Request) {
	signup := template.Must(template.ParseFiles("internal/api/web/templates/signup.html"))
	signup.Execute(w, nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	log.Println("Trying to log in as: ", username, password)
	log.Println(password)

	tkn, _ := auth.CreateToken(username) //creates the token with tkn func using username
	cookie := http.Cookie{
		Name:     "access_token",
		Value:    tkn,
		MaxAge:   20,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, &cookie) //cookifies is
	fmt.Fprintln(w, "Cooker set")
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	log.Println("These where used to make account: ", username, password)
	if len(password) > 0 && len(username) > 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func erreur(w http.ResponseWriter, _ *http.Request) { //le page d'erreur
	log.Println("TEST")
}
