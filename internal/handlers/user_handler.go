package handlers

import (
	"imob/internal/entities"
	"imob/internal/logger"
	"imob/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc  services.UserService
	log logger.Logger
}

func NewUserHandler(uc services.UserService, log logger.Logger) *UserHandler {
	return &UserHandler{uc: uc, log: log}
}

func (h *UserHandler) Register(r *gin.Engine) {
	r.POST("/users", h.CreateUser)
	r.GET("/users", h.GetUser)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req entities.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	out, err := h.uc.CreateUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, out)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	email := c.Query("email")
	ctx := c.Request.Context()
	u, err := h.uc.GetUser(ctx, email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, u)
}
