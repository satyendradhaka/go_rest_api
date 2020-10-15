package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var lock sync.Mutex

//articles data handling

type article struct {
	ID                int       `json:"ID"`
	Title             string    `json:"Title"`
	SubTitle          string    `json:"SubTitle"`
	Content           string    `json:"Content"`
	CreationTimestamp time.Time `json:"CreationTimestamp"`
}

type allArticles []article

type ids []int

var id = ids{1}

var articles = allArticles{
	{
		ID:                1,
		Title:             "Introduction to Golang",
		SubTitle:          "Jai mata di",
		Content:           "Come join us for a chance to learn how golang works and get to eventually try it out",
		CreationTimestamp: time.Now(),
	},
}

//function to support search route

func caseInsensitive(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

//functions to handle routes

func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("welcome to REST api")
}

func viewArticles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var afterID, limit int
		for k, v := range r.URL.Query() {
			if k == "after_id" {
				afterID, _ = strconv.Atoi(v[0])
			} else if k == "limit" {
				limit, _ = strconv.Atoi(v[0])
			}
		}
		fmt.Println(afterID)
		fmt.Println(limit)
		if afterID>=len(articles){
			json.NewEncoder(w).Encode("incorrect query for paginantion, after_id has crossed it's limit")
		}
		w.WriteHeader(http.StatusOK)
		if afterID+limit>len(articles){
			json.NewEncoder(w).Encode(articles[afterID:])
		}else {
			json.NewEncoder(w).Encode(articles[afterID : afterID+limit])
		}
	case "POST":
		lock.Lock()
		var newArticle article
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		id = append(id, id[len(id)-1]+1)
		json.Unmarshal(reqBody, &newArticle)
		newArticle.ID = id[len(id)-1]
		newArticle.CreationTimestamp = time.Now()
		fmt.Println(newArticle)
		articles = append(articles, newArticle)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(articles)
		lock.Unlock()
	}
}

func viewSingleArticle(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	fmt.Println(p)
	if len(p) == 1 {
		json.NewEncoder(w).Encode("enter correct path")
	} else {
		ID := p[2]
		fmt.Println(ID)
		articleID, _ := strconv.Atoi(ID)
		if articleID>id[len(id)-1]{
			json.NewEncoder(w).Encode("no article for this id")
		}else{
			json.NewEncoder(w).Encode(articles[articleID-1])
		}
	}
}

func searchArticle(w http.ResponseWriter, r *http.Request) {
	var articleFound bool = false
	for _, v := range r.URL.Query() {
		for _, singleArticle := range articles {
			if caseInsensitive(singleArticle.Title, v[0]) {
				articleFound = true
				json.NewEncoder(w).Encode(singleArticle)
			} else if caseInsensitive(singleArticle.SubTitle, v[0]) {
				articleFound = true
				json.NewEncoder(w).Encode(singleArticle)
			} else if caseInsensitive(singleArticle.Content, v[0]) {
				articleFound = true
				json.NewEncoder(w).Encode(singleArticle)
			}
		}
		if articleFound == false {
			json.NewEncoder(w).Encode("0 articles found")
		}
	}
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", viewArticles)
	http.HandleFunc("/articles/search", searchArticle)
	http.HandleFunc("/articles/", viewSingleArticle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
