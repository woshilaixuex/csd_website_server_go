package routes

import (
	"github.com/csd-world/csd_webstie_server_go/internal/handlers"
	"github.com/gin-gonic/gin"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: 路由注册
 * @Date: 2024-09-27 23:54
 */
func RegisterRoutes(r *gin.Engine, h *handlers.Handler) {
	// 注册路由
	apiGroup := r.Group("/api")
	{
		// POST 请求：提交报名表单
		apiGroup.POST("/enrolls", h.PostEnroll)
	}

}
