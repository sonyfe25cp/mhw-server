package models

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"com.github.sonyfe25cp.mhw-server/utils"
)

type Article struct {
	ID int64

	Title      string
	Content    string
	Source     string
	SourceType string
	Image      string

	CreateTime   int64
	ModifiedTime int64
}

func (a *Article) toString() string {

	return "CreateTime" + utils.Int64ToString(a.CreateTime)
}

type WeixinUser struct {
	ID int64

	Tokens  string
	HeadImg string
	Name    string
	Gender  int32
	VipDate int64
}

func ListArticles(dbUrl string, sourceType string, offset int, limit int) []Article {
	sqlContent := `select id, title, content, source, source_type, image, gmt_create, gmt_modified
				  from articles
				  where source_type = ?
				  limit ?, ?`

	var articles []Article
	db, err := sql.Open("mysql", dbUrl);
	defer func() {
		if db != nil {
			db.Close()
		}
		if recover() != nil {
			log.Println("list statistics error")
		}
	}()
	if err != nil {
		log.Println("error, can not open db:", err)
	} else {
		if err = db.Ping(); err == nil {
			if rows, err := db.Query(sqlContent, sourceType, offset, limit); err == nil {
				for rows.Next() {
					var id int64
					var title string
					var content string
					var source string
					var sourceType string
					var image string
					var gmtCreate int64
					var gmtModified int64

					err = rows.Scan(&id, &title, &content, &source, &sourceType, &image, &gmtCreate, &gmtModified)
					item := Article{ID: id, Title: title, Content: content, Source: source, SourceType: sourceType,
						Image: image, CreateTime: gmtCreate, ModifiedTime: gmtModified}
					articles = append(articles, item)
				}
			} else {
				log.Println("error, can not insert reqlog:", err)
			}
		} else {
			log.Println("error, can not ping:", err)
		}
	}
	return articles
}

func InsertArticles(dbUrl string, article Article) bool {
	sqlContent := "insert into articles(`gmt_create`, `gmt_modified`, `title`, `content`, `source`, `source_type`, `image`)" +
		"values(now(), now(), ?, ?, ?, ?, ?)"

	done := false
	defer func() {
		if recover() != nil {
			log.Println("something wrong about db")
		}
	}()
	if db, err := sql.Open("mysql", dbUrl); err != nil {
		log.Println("error, can not open db:", err)
	} else {
		if db != nil {
			defer db.Close()
		}
		if err = db.Ping(); err == nil {
			stmt, _ := db.Prepare(sqlContent)
			if stmt != nil {
				defer stmt.Close()
			}
			if _, err = stmt.Exec(article.Title, article.Content, article.Source, article.SourceType, article.Image); err != nil {
				log.Println("error, can not insert reqlog:", err)
			}
		} else {
			log.Println("error, can not ping:", err)
		}
	}
	return done
}
