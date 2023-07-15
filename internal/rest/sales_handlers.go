package rest

import (
	"mime/multipart"
	"net/http"

	internal "github.com/ShadrackAdwera/pg-csv/internal/sqlc"
	"github.com/gin-gonic/gin"
)

type salesCsvData struct {
	MatchCsvFile *multipart.FileHeader `form:"file" binding:"required"`
}

func (srv *Server) uploadSalesCsv(ctx *gin.Context) {
	var csvFile salesCsvData

	if err := ctx.ShouldBind(&csvFile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "provide the csv sales file"})
		return
	}

	if err := srv.distro.DistroDataOnCsv(ctx, csvFile.MatchCsvFile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to distribute task to queue"})
		return
	}

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
