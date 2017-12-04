package main

import (
	"time"
	"net/http"
	"html/template"
	"com.github.sonyfe25cp.mhw-server/models"
	"com.github.sonyfe25cp.mhw-server/utils"
	"com.github.sonyfe25cp.mhw-server/log"
)

type Server struct {
	Start time.Time
}

const DB_URL = "mhw_admin:mhw_admin_1803@tcp(127.0.0.1:3306)/mhw"

func Route(path string) string {
	var sourceType string
	switch path {
	case ArticleList:
		sourceType = News_SourceType
		break
	case Article:
		sourceType = News_SourceType
		break
	case IndexSlideList:
		sourceType = IndexSlides_SourceType
		break
	case ReferSlideList:
		sourceType = RerferSlides_SourceType
		break
	default:
		break
	}
	return sourceType
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(w, nil)
	}
}
func (s *Server) getArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := utils.StringtoIntWithDefault(r.FormValue("id"), 0)
		sourceType := Route(r.URL.Path)
		article := models.GetArticle(DB_URL, sourceType, id)
		utils.WriteJson(w, article)
	}
}
func (s *Server) listArticles(w http.ResponseWriter, r *http.Request) {
	utils.DebugFormValues(r.Form)
	logs.Info(r.URL.Path)
	if r.Method == "GET" {

		offset := utils.StringtoIntWithDefault(r.FormValue("offset"), 0)
		limit := utils.StringtoIntWithDefault(r.FormValue("limit"), 10)

		sourceType := Route(r.URL.Path)

		if len(sourceType) != 0 {
			articles := models.ListArticles(DB_URL, sourceType, offset, limit)
			utils.WriteJson(w, articles)
		} else {
			logs.Error("error url:", sourceType)
		}
	}
}

func (s *Server) adminListArticles(w http.ResponseWriter, r *http.Request) {
	utils.DebugFormValues(r.Form)
	logs.Info("st:", r.FormValue("source_type"))
	if r.Method == "GET" {

		offset := utils.StringtoIntWithDefault(r.FormValue("offset"), 0)
		limit := utils.StringtoIntWithDefault(r.FormValue("limit"), 10)
		sourceType := r.FormValue("source_type")

		articles := models.ListArticles(DB_URL, sourceType, offset, limit)

		t, _ := template.ParseFiles("templates/list_admin_articles_page.html")
		t.Execute(w, map[string]interface{}{
			"articles": articles,
		})
	}
}

func (s *Server) adminInsertArticles(w http.ResponseWriter, r *http.Request) {
	//logs.Info("method: ", r.Method)
	utils.DebugFormValues(r.Form)
	if r.Method == "POST" {

		titile := r.FormValue("title")
		content := r.FormValue("content")
		source := r.FormValue("source")
		sourceType := r.FormValue("source_type")
		image := r.FormValue("image")

		article := models.Article{Title: titile, Content: content, Source: source, SourceType: sourceType, Image: image}

		models.InsertArticles(DB_URL, article)

		http.Redirect(w, r, AdminArticleList, http.StatusFound)
	} else if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/insert_article_page.html")
		t.Execute(w, nil)
	}
}
