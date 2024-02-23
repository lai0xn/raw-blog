package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type H map[string]string

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	var article Article
	json.NewDecoder(r.Body).Decode(&article)
	fmt.Println(article)
	err := CreateArticle(article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	json.NewEncoder(w).Encode(H{
		"message": "article created",
	})
}

func GetAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := GetAllArticles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(articles)
}

func DeleteArticlesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	err := DeleteArticle(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	json.NewEncoder(w).Encode(H{
		"message": "article deleted successfully",
	})
}
