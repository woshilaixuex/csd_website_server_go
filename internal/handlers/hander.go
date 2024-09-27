package handlers

import (
	"net/http"

	"github.com/csd-world/csd_webstie_server_go/internal/models"
	"github.com/csd-world/csd_webstie_server_go/internal/services"
	"github.com/gin-gonic/gin"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description:
 * @Date: 2024-09-28 03:25
 */
type Handler struct {
	Service services.ApiService
}

func (h *Handler) PostEnroll(c *gin.Context) {
	var enroll models.EnrollTable
	if err := c.ShouldBindJSON(&enroll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.InsertEnroll(&enroll); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert enrollment"})
		return
	}

	c.JSON(http.StatusCreated, enroll)
}
