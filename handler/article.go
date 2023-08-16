package handler

import (
	"github.com/gin-gonic/gin"
	"go_dynamodb/database/model"
	"go_dynamodb/pkg/response"
	svctypes "go_dynamodb/types"
	"log"
	"strings"
)

func (app *Application) SaveArticle(c *gin.Context) {
	var req svctypes.Article

	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(c, "bad request", err.Error())
		return
	}

	err = app.ArticleStore.DescribeTable(c)
	if err != nil {
		err = app.ArticleStore.CreateTable(c)
		if err != nil {
			log.Println("create table failed")
			response.InternalServerError(c, "db operation failed", nil)
			return
		}
	}

	article := &model.Article{
		Title:   req.Title,
		Content: req.Content,
		Author:  req.Author,
	}

	err = app.ArticleStore.Save(c, article)
	if err != nil {
		log.Println("error save article", err.Error(), "article", article)
		response.InternalServerError(c, "db operation failed", nil)

		return
	}

	response.Success(c, "article created successfully", nil)
}

func (app *Application) GetArticle(c *gin.Context) {
	title := strings.TrimSpace(c.Param("title"))
	author := strings.TrimSpace(c.Param("author"))

	item, err := app.ArticleStore.Get(c, title, author)
	if err != nil {
		log.Println("error save article", err.Error(), "title", title)
		response.InternalServerError(c, "db operation failed", nil)

		return
	}

	if item == nil {
		response.NotFound(c, "article not found", nil)
		return
	}

	response.Success(c, "article data fetched successfully", item)
}

func (app *Application) GetAllArticles(c *gin.Context) {
	item, err := app.ArticleStore.GetAll(c)
	if err != nil {
		log.Println("error save article", err.Error())
		response.InternalServerError(c, "db operation failed", nil)

		return
	}

	response.Success(c, "article data fetched successfully", item)
}

func (app *Application) UpdateArticle(c *gin.Context) {
	var req svctypes.Article

	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(c, "bad request", err.Error())
		return
	}

	article := &model.Article{
		Title:   req.Title,
		Content: req.Content,
		Author:  req.Author,
	}

	err = app.ArticleStore.Update(c, article)
	if err != nil {
		log.Println("error update article", err.Error(), "article", article)
		response.InternalServerError(c, "db operation failed", nil)

		return
	}

	response.Success(c, "article updated successfully", nil)
}

func (app *Application) DeleteArticle(c *gin.Context) {
	var req svctypes.Article

	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(c, "bad request", err.Error())
		return
	}

	article := &model.Article{
		Title:  req.Title,
		Author: req.Author,
	}

	err = app.ArticleStore.Delete(c, article)
	if err != nil {
		log.Println("error delete article", err.Error(), "title:", article.Title)

		response.InternalServerError(c, "db operation failed", nil)

		return
	}

	response.Success(c, "article deleted successfully", nil)
}
