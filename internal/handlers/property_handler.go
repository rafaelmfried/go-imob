package handlers

import (
	"imob/internal/entities"
	"imob/internal/logger"
	"imob/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PropertyHandler struct {
	uc  services.PropertyService
	log logger.Logger
}

func NewPropertyHandler(uc services.PropertyService, log logger.Logger) *PropertyHandler {
	return &PropertyHandler{uc: uc, log: log}
}

func (h *PropertyHandler) Register(r *gin.Engine) {
	r.POST("/properties", h.CreateProperty)
	r.GET("/properties", h.GetProperty)
}

func (h *PropertyHandler) CreateProperty(c *gin.Context) {
	var req entities.Property
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	out, err := h.uc.Create(ctx, req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, out)
}

func (h *PropertyHandler) GetProperty(c *gin.Context) {
	owner_id := c.Query("owner_id")
	ctx := c.Request.Context()
	u, err := h.uc.Get(ctx, owner_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, u)
}
