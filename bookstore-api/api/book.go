package api

import (
	"database/sql"
	"net/http"

	db "github.com/alekseiapa/mini-go-projects/book-store/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createBookRequest struct {
	Name        string `json:"Name" binding:"required"`
	Publication string `json:"Publication" binding:"required"`
}

func (server *Server) createBook(ctx *gin.Context) {
	var req createBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateBookParams{
		Name:        req.Name,
		Publication: req.Publication,
	}
	user, err := server.store.CreateBook(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getBookRequest struct {
	Uuid int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getBook(ctx *gin.Context) {
	var req getBookRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	book, err := server.store.GetBook(ctx, req.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

type listBookRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

func (server *Server) listBook(ctx *gin.Context) {
	var req listBookRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListBooksParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	users, err := server.store.ListBooks(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

type updateBookRequestUri struct {
	Uuid int64 `uri:"id" binding:"required,min=1"`
}
type updateBookRequestJson struct {
	Name string `json:"Name" binding:"required"`
}

func (server *Server) updateBookName(ctx *gin.Context) {
	var reqUri updateBookRequestUri
	var reqJson updateBookRequestJson

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&reqJson); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateBookNameParams{
		Uuid: reqUri.Uuid,
		Name: reqJson.Name,
	}
	book, err := server.store.UpdateBookName(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, book)
}

type deleteBookRequest struct {
	Uuid int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteBook(ctx *gin.Context) {
	var req deleteBookRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteBook(ctx, req.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, successDeleteResponse())
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
func successDeleteResponse() gin.H {
	return gin.H{"success": "Deleted successfully"}
}
