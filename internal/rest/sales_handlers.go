package rest

import (
	"net/http"

	internal "github.com/ShadrackAdwera/pg-csv/internal/sqlc"
	"github.com/gin-gonic/gin"
)

func (srv *Server) uploadSalesCsv(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "ping POST method from IP: " + ctx.ClientIP()})
}

type ListSalesDataArgs struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (srv *Server) fetchSalesData(ctx *gin.Context) {
	var listSalesDataArgs ListSalesDataArgs

	if err := ctx.ShouldBindQuery(&listSalesDataArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error - invalid query"})
		return
	}

	salesData, err := srv.store.ListSales(ctx, internal.ListSalesParams{
		Limit:  listSalesDataArgs.PageSize,
		Offset: (listSalesDataArgs.PageID - 1) * listSalesDataArgs.PageSize,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error fetching sales data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": salesData})
}
