package controllers

import (
	"go_todo/app/models"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			log.Println(err)
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		newUser := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		user, err := models.GetUserByEmail(newUser.Email)
		if err != nil {
			log.Println(err)
		}
		if &user != nil {
			log.Println("user is already exists")
			message := "このメールアドレスはすでに登録されています"
			generateHTML(w, message, "layout", "public_navbar", "signup")
		} else {
			if err := newUser.CreateUser(); err != nil {
				log.Println(err)
			}

			http.Redirect(w, r, "/", 302)
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			log.Println(err)
			generateHTML(w, nil, "layout", "public_navbar", "login")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		authenticate(w, r)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		message := "ログインに失敗しました"
		generateHTML(w, message, "layout", "public_navbar", "login")
		return
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		message := "ログインに失敗しました"
		generateHTML(w, message, "layout", "public_navbar", "login")
		return
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", 302)
}
