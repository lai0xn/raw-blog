package main

import "github.com/gorilla/mux"

func InitRoutes(r *mux.Router) {
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/articles", GetAllArticlesHandler).Methods("GET")
	r.HandleFunc("/article/create", CreateArticleHandler).Methods("POST")
	r.HandleFunc("/article/delete/{id}", DeleteArticlesHandler).Methods("DELETE")
}
