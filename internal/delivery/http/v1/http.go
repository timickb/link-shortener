package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/timickb/link-shortener/internal/config"
	"github.com/timickb/link-shortener/internal/interfaces"
	"net/http"
)

type Server struct {
	cfg       *config.AppConfig
	log       interfaces.Logger
	shortener interfaces.Shortener
	router    *gin.Engine
}

func New(log interfaces.Logger, cfg *config.AppConfig, short interfaces.Shortener) *Server {
	srv := &Server{
		cfg:       cfg,
		log:       log,
		shortener: short,
		router:    gin.New(),
	}

	api := srv.router.Group("api/v1")
	api.GET("/restore", srv.restore)
	api.POST("/create", srv.create)

	return srv
}

func (s *Server) Run() error {
	if err := s.router.Run(fmt.Sprintf(":%d", s.cfg.HTTPPort)); err != nil {
		return err
	}
	return nil
}

func (s *Server) restore(ctx *gin.Context) {
	shortened := ctx.Query("shortened")
	if shortened == "" {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "empty shortened value",
		})
		return
	}

	original, err := s.shortener.RestoreLink(shortened)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    RestoreResponse{Original: original},
	})
}

func (s *Server) create(ctx *gin.Context) {
	req := &CreateShorteningRequest{}

	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid body",
		})
		return
	}

	short, err := s.shortener.CreateLink(req.URL)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, Response{
		Code:    http.StatusCreated,
		Message: "created",
		Data:    CreateShorteningResponse{Short: short},
	})
}
