package controller

import (
	"JET_DEV_TASK/JetDev/api/service"
	"JET_DEV_TASK/JetDev/models"
	"JET_DEV_TASK/JetDev/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	service service.BlogService
}

func NewBlogController(s service.BlogService) BlogController {
	return BlogController{
		service: s,
	}
}

//ListAll
func (b BlogController) ListAll(ctx *gin.Context) {
	var articles models.Article
	data, total, err := b.service.FindAll(articles)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find ")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)
	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}
	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Post result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})

}

//ListComment
func (b BlogController) ListComment(ctx *gin.Context) {
	var articles models.Article
	data, total, err := b.service.FindAll(articles)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find ")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)
	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}
	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Post result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})

}

//ListContent
func (b BlogController) ListContent(ctx *gin.Context) {
	var articles models.Article
	keyword := ctx.Query("keyword")
	data, total, err := b.service.ListContent(articles, keyword)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}
	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Post result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})

}

// AddArticle
func (b *BlogController) AddArticle(ctx *gin.Context) {
	var articles models.Article
	ctx.ShouldBindJSON(&articles)

	if articles.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if articles.Nickname == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Nickname is required")
		return
	}
	if articles.Content == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Content is required")
		return
	}

	err := b.service.Save(articles)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create article")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Article")
}
