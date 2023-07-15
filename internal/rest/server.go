package rest

import (
	internal "github.com/ShadrackAdwera/pg-csv/internal/sqlc"
	"github.com/ShadrackAdwera/pg-csv/internal/workers"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	router *gin.Engine
	store  internal.TxStore
	pool   *pgxpool.Pool
	distro workers.TaskDistributor
}

func NewServer(pool *pgxpool.Pool, store internal.TxStore, distro workers.TaskDistributor) *Server {
	router := gin.Default()

	srv := &Server{
		pool:   pool,
		store:  store,
		distro: distro,
	}

	router.POST("/api/sales-upload", srv.uploadSalesCsv)
	router.GET("/api/sales", srv.fetchSalesData)
	srv.router = router
	return srv
}

func (s *Server) StartServer(addr string) error {
	return s.router.Run(addr)
}
