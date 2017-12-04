package main

import (
	"flag"
	"log"
	"time"
	"net/http"
	_ "net/http/pprof"
)

const (
	ArticleList    = "/articles"
	IndexSlideList = "/index_slides"
	ReferSlideList = "/refer_slides"

	Article = "/article"

	AdminArticleList   = "/admin/articles"
	AdminArticleInsert = "/admin/insert_article"
)

const (

	News_SourceType         = "news"
	IndexSlides_SourceType  = "indexSlides"
	RerferSlides_SourceType = "referSlides"
)

func main() {
	var keyPath, addr string
	var test bool
	flag.StringVar(&addr, "addr", "0.0.0.0:8080", "Config file path")
	flag.StringVar(&keyPath, "keyPath", "/Users/aiao/golang/ownpkg/src/com.github.sonyfe25cp.mhw-server/sample/", "Server Key Path")
	flag.BoolVar(&test, "test", false, "Test config file and exits")
	flag.Parse()

	//flag.BoolVar(private, "online", false, "whether this program is proxy product-env api or daily-env api.")

	server := &Server{Start: time.Now()}

	http.HandleFunc("/", server.index)                   //设置访问的路由
	http.HandleFunc(Article, server.getArticle)    //设置访问的路由
	http.HandleFunc(ArticleList, server.listArticles)    //设置访问的路由
	http.HandleFunc(IndexSlideList, server.listArticles) //设置访问的路由
	http.HandleFunc(ReferSlideList, server.listArticles) //设置访问的路由

	http.HandleFunc(AdminArticleList, server.adminListArticles)     //设置访问的路由
	http.HandleFunc(AdminArticleInsert, server.adminInsertArticles) //设置访问的路由

	//http.HandleFunc("/stats", server.status)    //设置访问的路由

	log.Println("listen on :", addr)
	//crtFile := keyPath + "server.crt"
	//keyFile := keyPath + "server.key"
	//err := http.ListenAndServeTLS(addr, crtFile, keyFile, nil) //设置监听的端口

	err := http.ListenAndServe(addr, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
