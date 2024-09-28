package handlers

import (
	"fmt"
	"net/http"

	"github.com/csd-world/csd_webstie_server_go/internal/models"
	"github.com/csd-world/csd_webstie_server_go/internal/models/engine"
	"github.com/csd-world/csd_webstie_server_go/internal/services"
	"github.com/csd-world/csd_webstie_server_go/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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

func NewHandler(engine *engine.MysqlEngine, resultCh *pkg.MssChan) *Handler {
	return &Handler{
		Service: *services.NewApiService(engine, resultCh),
	}
}
func (h *Handler) PostEnroll(c *gin.Context) {
	// 定义一个 ApiReqBody 来绑定请求数据
	var req models.ApiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 请求体绑定失败，返回错误响应
		c.JSON(http.StatusBadRequest, models.ApiResponse{
			Code: http.StatusBadRequest,
			Msg:  "Invalid request body",
			Data: &models.ApiRespBody{
				Info: "",
			},
		})
		return
	}

	// 将请求数据转换为 EnrollTable 模型
	var enroll models.EnrollTable
	if err := copier.Copy(&enroll, &req.Data); err != nil {
		panic(fmt.Sprintf("request copy is fail: %v", err))
	}

	// 调用 service 层插入数据
	if err := h.Service.InsertEnroll(&enroll); err != nil {
		// 插入失败，返回错误响应
		c.JSON(http.StatusInternalServerError, models.ApiResponse{
			Code: http.StatusInternalServerError,
			Msg:  "failed to insert enrollment",
			Data: nil,
		})
		return
	}
	ResponseData := "报名成功"
	// 插入成功，返回成功响应
	c.JSON(http.StatusCreated, models.ApiResponse{
		Code: http.StatusCreated,
		Msg:  "enrollment created successfully",
		Data: &models.ApiRespBody{
			Info: ResponseData,
		},
	})
}
