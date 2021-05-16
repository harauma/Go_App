package controllers

import (
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		log.Println(user)
		if err != nil {
			log.Println(err)
		}
		todos, err := user.GetTodosByUser()
		if err != nil {
			log.Println(err)
		}
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}
