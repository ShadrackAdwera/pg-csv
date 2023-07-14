package rest

import (
	internal "github.com/ShadrackAdwera/pg-csv/internal/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	router *gin.Engine
	store  internal.TxStore
	pool   *pgxpool.Pool
}

func NewServer(pool *pgxpool.Pool, store internal.TxStore) *Server {
	router := gin.Default()

	srv := &Server{
		pool:  pool,
		store: store,
	}

	router.POST("/api/sales-upload", srv.uploadSalesCsv)
	router.GET("/api/sales", srv.fetchSalesData)
	srv.router = router
	return srv
}

func (s *Server) StartServer(addr string) error {
	return s.router.Run(addr)
}
